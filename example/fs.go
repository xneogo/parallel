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
 @Time    : 2024/9/13 -- 09:54
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: fs.go
*/

package example

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func CallString(s string) (string, int) {
	fmt.Println("Running CallString")
	time.Sleep(time.Second)
	d, _ := strconv.Atoi(s)
	return s, d
}

func CallDts(x int) string {
	fmt.Println("Running CallDts")
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("%d", x)
}

type Obj struct {
	Name string
}

func CallOther(x, y int) *Obj {
	fmt.Println("Running CallOther")
	time.Sleep(time.Second * 3)
	return &Obj{Name: fmt.Sprintf("called: %d~%d", x, y)}
}

func (o *Obj) String() string {
	fmt.Println("Running Obj.String")
	time.Sleep(time.Second * 4)
	return o.Name
}

type BadCall struct{}

func (b *BadCall) SliceOutOfRange() int {
	s := []int{1, 2, 3, 4, 5}
	return s[100]
}

func (b *BadCall) Assignment2NilMap() {
	var m map[int]*Obj
	m[1] = &Obj{Name: "1"}
}

type A struct {
	Id int64    `json:"id"`
	A1 []string `json:"a1"`
	A2 int32    `json:"a2"`
}

type DataA struct {
	Id int64  `gorm:"column:id" json:"id"`
	A1 string `gorm:"column:a1" json:"a1"`
	A2 int32  `gorm:"column:a2" json:"a2"`
}

func (da DataA) Dto() *A {
	var a1 []string
	_ = json.Unmarshal([]byte(da.A1), &a1)
	return &A{
		Id: da.Id,
		A1: a1,
		A2: da.A2,
	}
}

func GetInfoA(ctx context.Context, id int64) *A {
	fmt.Println("Running GetInfoA")
	time.Sleep(time.Second)
	// get from mysql table a by id
	da := &DataA{
		Id: id,
		A1: `["i am a11", "i am a12"]`,
		A2: int32(id * 100),
	}
	// data 2 obj
	a := da.Dto()
	return a
}

type B struct {
	Id int64
	B1 string
	B2 int
}

type DataB struct {
	Id int64  `gorm:"column:id" json:"id"`
	B1 string `gorm:"column:b1" json:"b1"`
	B2 int32  `gorm:"column:b2" json:"b2"`
}

func (da DataB) Dto() *B {
	return &B{
		Id: da.Id,
		B1: da.B1,
		B2: int(da.B2),
	}
}

func GetInfoB(ctx context.Context, id int64) *B {
	fmt.Println("Running GetInfoB")
	time.Sleep(time.Second * 2)
	// get from mysql table a by id
	da := &DataB{
		Id: id,
		B1: fmt.Sprintf("B: %d", id),
		B2: int32(id * 100),
	}
	// data 2 obj
	a := da.Dto()
	return a
}

type AB struct {
	A *A `json:"a"`
	B *B `json:"b"`
}

func AAndB(a *A, b *B) map[string]interface{} {
	fmt.Println("Running AAndB", a, b)
	time.Sleep(time.Second)
	resObj := &AB{
		A: a,
		B: b,
	}
	bs, _ := json.Marshal(resObj)
	fmt.Println(string(bs))
	res := make(map[string]interface{})
	_ = json.Unmarshal(bs, &res)
	return res
}
