package stack

import (
	"container/list"
	"sync"
)

type ConcurrentStack struct {
	lock sync.Mutex
	list *list.List
}

func NewConcurrentStack() *ConcurrentStack {
	return &ConcurrentStack{
		list: list.New(),
	}
}

func (stack *ConcurrentStack) Len() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	return stack.list.Len()
}

func (stack *ConcurrentStack) Empty() bool {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	return stack.list.Len() == 0
}

func (stack *ConcurrentStack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.list.PushBack(value)
}

func (stack *ConcurrentStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.list.Len() == 0 {
		panic("can not pop object from an empty stack")
	}

	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *ConcurrentStack) Peak() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.list.Len() == 0 {
		panic("can not peek object from an empty stack")
	}

	e := stack.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}
