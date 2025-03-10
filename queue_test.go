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
 @Time    : 2024/9/13 -- 09:53
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: queue_test.go
*/

package parallel

import (
	"fmt"
	"github.com/eva-nigouki/parallel/examples"
	"testing"
	"time"
)

func Test(*testing.T) {
	var cs string
	var csd int
	var dts string
	var obj *examples.Obj
	var obj2 = &examples.Obj{Name: "hello again"}
	var obj2r string

	startTime := time.Now().Unix()
	que := NewQueue()
	que.Push(examples.CallString, "Hello World").SetRes(&cs, &csd)
	que.Push(examples.CallDts, 123).SetRes(&dts)
	que.Push(examples.CallOther, 1, 2).SetRes(&obj)
	que.Push(obj2.String).SetRes(&obj2r)
	que.Purge()
	endTime := time.Now().Unix()

	fmt.Println(cs, csd)
	fmt.Println(dts)
	fmt.Println(obj)
	fmt.Println(obj2r)
	fmt.Printf("time cost:%d s\n", endTime-startTime)
}
