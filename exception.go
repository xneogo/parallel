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
 @Time    : 2024/9/12 -- 16:30
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: exception.go
*/

package parallel

import "fmt"

// ExceptionProxy
// closure interface, you can build your own exception dealer
type ExceptionProxy interface {
	Deal(args ...any) Dealer
}

type Dealer func(any)

type Exception struct{}

func DefaultException() Exception {
	return Exception{}
}

func (e Exception) Deal() Dealer {
	return func(err any) {
		fmt.Println(err)
	}
}
