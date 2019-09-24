package queue

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (queue *Queue) Len() int {
	return queue.list.Len()
}

func (queue *Queue) Empty() bool {
	return queue.list.Len() == 0
}

func (queue *Queue) Put(value interface{}) {
	queue.list.PushFront(value)
}

func (queue *Queue) Get() interface{} {
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

func (queue *Queue) Peek() interface{} {
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
