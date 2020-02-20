package main

import "testing"

func TestStackIsLIFO(t *testing.T) {
	s := jumpStack{}
	b1 := block{StartIndex: 5}
	b2 := block{StartIndex: 10}
	s.Push(b1)
	s.Push(b2)
	v, ok := s.Pop()
	if !ok {
		t.Errorf("Expected ok=%v, received %v", true, ok)
	}
	if v != b2 {
		t.Errorf("Expected value=%v, received %v", b2, v)
	}
}

func TestStackInitialState(t *testing.T) {
	s := jumpStack{}
	v, ok := s.Pop()
	emptyBlock := block{}
	if ok {
		t.Errorf("Expected ok=%v, received %v", false, ok)
	}
	if v != emptyBlock {
		t.Errorf("Expected value=%v, received %v", emptyBlock, v)
	}
}

func TestStackEmpty(t *testing.T) {
	s := jumpStack{}
	emptyBlock := block{}
	s.Push(block{StartIndex: 5})
	s.Push(block{StartIndex: 10})
	_, _ = s.Pop()
	_, _ = s.Pop()
	v, ok := s.Pop()
	if ok {
		t.Errorf("Expected ok=%v, received %v", false, ok)
	}
	if v != emptyBlock {
		t.Errorf("Expected value=%v, received %v", emptyBlock, v)
	}
}
