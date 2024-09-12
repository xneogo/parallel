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
 @Time    : 2024/9/12 -- 15:12
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: errors.go
*/

package parallel

import "errors"

var (
	ErrNotAFunction           = errors.New("[parallel]: f is not a function")
	ErrArgInputLengthNotMatch = errors.New("[parallel]: arg input length not match")
	ErrResOutOfRange          = errors.New("[parallel]: res out of range")
	ErrResTypeNotAPtr         = errors.New("[parallel]: res type is not a pointer")
	ErrResNil                 = errors.New("[parallel]: res is nil")
)
