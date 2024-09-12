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
	"github.com/qiguanzhu/parallel"
	"strconv"
)

/*
CallString | CallDts | CallOther
           |         |
End
*/

func CallString(s string) (string, int) {
	d, _ := strconv.Atoi(s)
	return s, d
}

func CallDts(x int) string {
	return fmt.Sprintf("%d", x)
}

type Obj struct {
	name string
}

func CallOther(x, y int) *Obj {
	return &Obj{name: fmt.Sprintf("called: %d~%d", x, y)}
}

func (o *Obj) String() string {
	return o.name
}

func main() {
	var cs string
	var csd int
	var dts string
	var obj *Obj
	var obj2 = &Obj{name: "hello again"}
	var obj2r string

	p := parallel.NewParallel()
	// normally you can use like this
	// add a func
	p.Add(CallString, "Hello World").SetRes(&cs, &csd)
	p.Add(CallDts, 123).SetRes(&dts)
	p.Add(CallOther, 1, 2).SetRes(&obj)
	p.Add(obj2.String).SetRes(&obj2r)
	// block and wait until all func done
	p.Wait()
	fmt.Println(cs, csd)
	fmt.Println(dts)
	fmt.Println(obj)
	fmt.Println(obj2r)
}
