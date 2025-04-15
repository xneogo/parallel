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
 @Time    : 2024/9/13 -- 10:05
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: usage.go
*/

package main

import (
	"fmt"
	"github.com/xneogo/parallel"
	"github.com/xneogo/parallel/examples"
)

type MyException struct{}

func (e MyException) Deal(args ...any) parallel.Dealer {
	return func(err any) {
		fmt.Println(err, args)
	}
}

func main() {
	var a int
	bc := new(examples.BadCall)

	p := parallel.NewParallel()
	p.Exception(MyException{})
	p.Add(bc.Assignment2NilMap)
	p.Add(bc.SliceOutOfRange).SetRes(&a)
	p.Wait(1, 2, 3)
}
