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

import "sync"

type Parallel struct {
	wg        *sync.WaitGroup
	que       []*Queue
	childWg   *sync.WaitGroup
	children  []*Parallel
	exception Exception
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

func (p *Parallel) Exception(exception Exception) *Parallel {
	p.exception = exception
	return p
}

func (p *Parallel) Add(f any, args ...any) *Executor {
	return p.Queue().Push(f, args...)
}

func (p *Parallel) Queue() *Queue {
	q := NewQueue()
	p.AddQueue(q)
	return q
}

func (p *Parallel) AddQueue(q ...*Queue) *Parallel {
	p.wg.Add(len(q))
	p.que = append(p.que, q...)
	return p
}

func (p *Parallel) GiveBirth() *Parallel {
	child := NewParallel()
	child.exception = p.exception
	p.ChildJoin(child)
	return child
}

func (p *Parallel) ChildJoin(child ...*Parallel) *Parallel {
	p.childWg.Add(len(p.children))
	p.children = append(p.children, child...)
	return p
}

func (p *Parallel) Wait() {
	for _, child := range p.children {
		go func(ch *Parallel) {
			ch.Wait()
			p.childWg.Done()
		}(child)
	}
	p.childWg.Wait()
	p.do()
	p.wg.Wait()
}

func (p *Parallel) do() {
	if len(p.que) == 1 {
		p.SafeWrapper(p.que[0])
		return
	}
	for _, q := range p.que {
		go p.SafeWrapper(q)
	}
}

func (p *Parallel) SafeWrapper(q *Queue) {
	defer func() {
		err := recover()
		switch err {
		case ErrNotAFunction, ErrArgInputLengthNotMatch, ErrResTypeNotAPtr, ErrResOutOfRange, ErrResNil:
			// basic err, throw panic again
			panic(err)
		case nil:
			p.wg.Done()
		default:
			p.exception.Deal()(err)
		}
	}()
	q.Purge()
}
