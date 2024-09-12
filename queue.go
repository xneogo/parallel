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
 @Time    : 2024/9/12 -- 15:28
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: queue.go
*/

package parallel

type Queue struct {
	handlers []*Executor
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(f any, args ...any) *Executor {
	h := NewExecutor(f, args...)
	q.push(h)
	return h
}

func (q *Queue) push(h *Executor) *Queue {
	q.handlers = append(q.handlers, h)
	return q
}

func (q *Queue) Purge() {
	for _, h := range q.handlers {
		h.Do()
	}
}
