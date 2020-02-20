package main

import (
	"errors"
	"fmt"
)

type instruction func(s *state) error

var instructionMap = map[byte]instruction{
	'.': omitIfSkipBlock(print),
	'<': omitIfSkipBlock(moveLeft),
	'>': omitIfSkipBlock(moveRight),
	'+': omitIfSkipBlock(increment),
	'-': omitIfSkipBlock(decrement),
	',': omitIfSkipBlock(read),
	'[': blockStart,
	']': blockEnd,
}

func omitIfSkipBlock(inner instruction) instruction {
	return func(s *state) error {
		if s.skipBlock {
			return nil
		}
		return inner(s)
	}

}

func print(s *state) error {
	v := s.memory[s.dataPointer]
	fmt.Fprintf(s.output, "%c", v)
	return nil
}

func moveLeft(s *state) error {
	s.dataPointer--
	return nil
}

func moveRight(s *state) error {
	s.dataPointer++
	return nil
}

func increment(s *state) error {
	s.memory[s.dataPointer]++
	return nil
}

func decrement(s *state) error {
	s.memory[s.dataPointer]--
	return nil
}

func read(s *state) error {
	v, err := s.input.ReadByte()
	if err != nil {
		return err
	}
	s.memory[s.dataPointer] = v
	return nil
}

func blockStart(s *state) error {
	skipBlok := s.memory[s.dataPointer] == byte(0)
	b := block{
		StartIndex:     s.instructionIndex,
		SkipOuterBlock: s.skipBlock,
	}
	s.skipBlock = skipBlok
	s.stack.Push(b)
	return nil
}

func blockEnd(s *state) error {
	currentBlock, ok := s.stack.Pop()
	if !ok {
		return errors.New("missing opening tag")
	}
	if s.skipBlock {
		s.skipBlock = currentBlock.SkipOuterBlock
		return nil
	}
	if s.memory[s.dataPointer] != byte(0) {
		s.instructionIndex = currentBlock.StartIndex - 1
	}
	return nil
}
