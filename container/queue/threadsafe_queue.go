package queue

import (
	"container/list"
	"sync"
)

type ConcurrentQueue struct {
	lock sync.Mutex
	list *list.List
}

func NewConcurrentQueue() *ConcurrentQueue {
	return &ConcurrentQueue{
		list: list.New(),
	}
}

func (queue *ConcurrentQueue) Len() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	return queue.list.Len()
}

func (queue *ConcurrentQueue) Empty() bool {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	return queue.list.Len() == 0
}

func (queue *ConcurrentQueue) Put(value interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.list.PushFront(value)
}

func (queue *ConcurrentQueue) Get() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.list.Len() == 0 {
		panic("can not get object from an empty queue")
	}

	e := queue.list.Back()
	if e != nil {
		queue.list.Remove(e)
		return e.Value
	}
	return nil
}

func (queue *ConcurrentQueue) Peek() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.list.Len() == 0 {
		panic("can not peek object from an empty queue")
	}

	e := queue.list.Back()
	if e != nil {
		queue.list.Remove(e)
		return e.Value
	}
	return nil
}
