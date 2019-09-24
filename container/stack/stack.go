package stack

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
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

func (stack *Stack) Peak() interface{} {
	if stack.list.Len() == 0 {
		panic("can not peek object from an empty stack")
	}
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}
