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
 @Time    : 2024/9/12 -- 11:57
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: parallel.go
*/

package parallel

import (
	"fmt"
	"sync"
)

type Parallel struct {
	wg        *sync.WaitGroup
	que       []*Queue
	childWg   *sync.WaitGroup
	children  []*Parallel
	exception ExceptionProxy
}

func NewParallel() *Parallel {
	return &Parallel{
		wg:        &sync.WaitGroup{},
		que:       make([]*Queue, 0),
		childWg:   &sync.WaitGroup{},
		children:  make([]*Parallel, 0),
		exception: DefaultException(),
	}
}

func (p *Parallel) Exception(exception ExceptionProxy) *Parallel {
	p.exception = exception
	return p
}

func (p *Parallel) Add(f any, args ...any) *Executor {
	return p.Queue().Push(f, args...)
}

func (p *Parallel) Queue() *Queue {
	q := NewQueue()
	p.addQueue(q)
	return q
}

func (p *Parallel) addQueue(q ...*Queue) *Parallel {
	p.wg.Add(len(q))
	p.que = append(p.que, q...)
	return p
}

func (p *Parallel) GiveBirth() *Parallel {
	child := NewParallel()
	child.exception = p.exception
	p.childJoin(child)
	return child
}

func (p *Parallel) childJoin(child ...*Parallel) *Parallel {
	p.childWg.Add(len(child))
	p.children = append(p.children, child...)
	return p
}

func (p *Parallel) Wait(args ...any) {
	for _, child := range p.children {
		go func(ch *Parallel) {
			ch.Wait(args...)
			p.childWg.Done()
		}(child)
	}
	p.childWg.Wait()
	fmt.Println("child wait finished")
	p.do(args...)
	p.wg.Wait()
}

func (p *Parallel) do(args ...any) {
	if len(p.que) == 1 {
		p.safeWrapper(p.que[0], args...)
		return
	}
	for _, q := range p.que {
		go p.safeWrapper(q, args...)
	}
}

func (p *Parallel) safeWrapper(q *Queue, args ...any) {
	defer func() {
		err := recover()
		switch err {
		case ErrNotAFunction, ErrArgInputLengthNotMatch, ErrResTypeNotAPtr, ErrResOutOfRange, ErrResNil:
			// basic err, throw panic again
			panic(err)
		default:
			p.wg.Done()
			if err != nil {
				p.exception.Deal(args...)(err)
			}
		}
	}()
	q.Purge()
}
