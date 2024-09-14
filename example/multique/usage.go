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
 @Time    : 2024/9/13 -- 12:00
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: usage.go
*/

package main

import (
	"fmt"
	"github.com/qiguanzhu/parallel"
	"github.com/qiguanzhu/parallel/example"
	"time"
)

/*
start
CallString | CallDts   | CallOther
Obj.String |
End
*/

func main() {
	var cs string
	var csd int
	var dts string
	var obj *example.Obj
	var obj2 = &example.Obj{Name: "hello again"}
	var obj2r string

	p := parallel.NewParallel()
	startTime := time.Now().Unix()
	// process func inline
	// add a queue to parallel
	que1 := p.Queue()
	que1.Push(example.CallString, "Hello World").SetRes(&cs, &csd)
	que1.Push(example.CallOther, 1, 2).SetRes(&obj)
	que2 := p.Queue()
	que2.Push(example.CallDts, 123).SetRes(&dts)
	p.Add(obj2.String).SetRes(&obj2r)
	// block and wait until all func done
	// total cost of time should be max(1+3,2,4)=4 seconds
	p.Wait()
	endTime := time.Now().Unix()
	fmt.Println(cs, csd) // Hello World 0
	fmt.Println(dts)     // 123
	fmt.Println(obj)     // called: 1~2
	fmt.Println(obj2r)   // hello again
	fmt.Printf("time cost: %d s\n", endTime-startTime)
}
