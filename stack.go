package main

type block struct {
	StartIndex     int
	SkipOuterBlock bool
}

type stackElement struct {
	value block
	next  *stackElement
}

type jumpStack struct {
	top *stackElement
}

func (s *jumpStack) Push(b block) {
	newTop := stackElement{
		value: b,
	}
	if s.top != nil {
		oldTop := s.top
		newTop.next = oldTop
	}
	s.top = &newTop
}

func (s *jumpStack) Pop() (block, bool) {
	if s.top == nil {
		return block{}, false
	}
	oldTop := *s.top
	s.top = oldTop.next
	return oldTop.value, true
}
