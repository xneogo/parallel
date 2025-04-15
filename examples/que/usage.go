/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2024/9/13 -- 09:37
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: usage.go
*/

package main

import (
	"fmt"
	"github.com/xneogo/parallel"
	"github.com/xneogo/parallel/examples"
	"time"
)

/*
start
CallString |
CallDts    |
CallOther  |
Obj.String |
End
*/

func main() {
	var cs string
	var csd int
	var dts string
	var obj *examples.Obj
	var obj2 = &examples.Obj{Name: "hello again"}
	var obj2r string

	p := parallel.NewParallel()
	startTime := time.Now().Unix()
	// process func inline
	// add a queue to parallel
	que := p.Queue()
	que.Push(examples.CallString, "Hello World").SetRes(&cs, &csd)
	que.Push(examples.CallDts, 123).SetRes(&dts)
	que.Push(examples.CallOther, 1, 2).SetRes(&obj)
	que.Push(obj2.String).SetRes(&obj2r)
	// block and wait until all func done
	// total cost of time should be 1+2+3+4=10 seconds
	fmt.Println(que)
	p.Wait()
	endTime := time.Now().Unix()
	fmt.Println(cs, csd) // Hello World 0
	fmt.Println(dts)     // 123
	fmt.Println(obj)     // called: 1~2
	fmt.Println(obj2r)   // hello again
	fmt.Printf("time cost: %d s\n", endTime-startTime)
}
