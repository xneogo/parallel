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
 @Time    : 2024/9/12 -- 16:42
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
CallString | CallDts | CallOther
           |         |
End
*/

func main() {
	var cs string
	var csd int
	var dts string
	var obj examples.Obj
	var obj2 = &examples.Obj{Name: "hello again"}
	var obj2r string

	p := parallel.NewParallel()
	startTime := time.Now().Unix()
	// normally you can use like this
	// add a func
	p.Add(examples.CallString, "Hello World").SetRes(&cs, &csd)
	p.Add(examples.CallDts, 123).SetRes(&dts)
	p.Add(examples.CallOther, 1, 2).SetRes(&obj) // should be value
	p.Add(obj2.String).SetRes(&obj2r)
	// block and wait until all func done
	// total cost of time should be 4 seconds
	p.Wait()
	endTime := time.Now().Unix()
	fmt.Println(cs, csd) // Hello World 0
	fmt.Println(dts)     // 123
	fmt.Println(obj)     // called: 1~2
	fmt.Println(obj2r)   // hello again
	fmt.Printf("time cost: %d s\n", endTime-startTime)
}
