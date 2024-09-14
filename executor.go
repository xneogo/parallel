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
 @Time    : 2024/9/12 -- 15:09
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: executor.go
*/

package parallel

import (
	"reflect"
)

type Executor struct {
	f    any
	args []any
	res  []any
}

func NewExecutor(f any, args ...any) *Executor {
	return &Executor{f: f, args: args}
}

func (h *Executor) SetRes(res ...any) *Executor {
	h.res = res
	return h
}

func (h *Executor) Do() {
	fValue := reflect.ValueOf(h.f)
	fType := fValue.Type()
	if fType.Kind() != reflect.Func {
		panic(ErrNotAFunction)
	}
	if fType.NumIn() != len(h.args) {
		panic(ErrArgInputLengthNotMatch)
	}
	if fType.NumOut() != len(h.res) {
		panic(ErrResOutOfRange)
	}

	for _, r := range h.res {
		if rValue := reflect.ValueOf(r); rValue.Type().Kind() != reflect.Ptr {
			panic(ErrResTypeNotAPtr)
		} else if rValue.IsNil() {
			panic(ErrResNil)
		}
	}

	inputs := make([]reflect.Value, len(h.args))
	for i, arg := range h.args {
		if arg == nil {
			inputs[i] = reflect.Zero(fType.In(i))
		} else {
			inputs[i] = reflect.ValueOf(arg)
		}
	}
	outputs := fValue.Call(inputs)
	for i, r := range h.res {
		value := reflect.ValueOf(r)
		// There are 2 ways of output: value or ptr
		// so if r and output[i] are both ptr, we need specify their means to set value
		if outputs[i].Type().Kind() != reflect.Ptr {
			value.Elem().Set(outputs[i])
		} else {
			value.Elem().Set(outputs[i].Elem())
		}
	}
}
