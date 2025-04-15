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
 @Time    : 2024/9/13 -- 11:40
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: usage.go
*/

package main

import (
	"context"
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
	// var a *example.A
	// var b *example.B
	// a = &example.A{}
	// b = &example.B{}
	// a := new(example.A)
	// b := new(example.B)
	var a examples.A
	var b examples.B
	var res map[string]interface{}
	ctx := context.Background()

	p := parallel.NewParallel()
	startTime := time.Now().Unix()
	// normally in real life, coding will be like this
	// we get info A from sql or other grpc
	// we get info B from sql or other grpc
	// we use A and B to build a result suit for http route endpoints
	// when query process of A and B has nothing in common, then they can be speed up by parallel
	childA := p.GiveBirth()
	childA.Add(examples.GetInfoA, ctx, int64(1)).SetRes(&a) // cost 1 second
	childB := p.GiveBirth()
	childB.Add(examples.GetInfoB, ctx, int64(2)).SetRes(&b) // cost 2 seconds
	// childA and childB will be processed parallel first
	p.Add(examples.AAndB, &a, &b).SetRes(&res) // cost 1 second
	p.Wait()
	endTime := time.Now().Unix() // cost max(1,2) + 1 = 3s
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(res)
	fmt.Printf("time cost: %d s\n", endTime-startTime)
}
