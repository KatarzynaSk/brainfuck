package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const MEMORY_SIZE = 1000

type state struct {
	dataPointer      int
	memory           [MEMORY_SIZE]byte
	stack            jumpStack
	instructionIndex int
	skipBlock        bool
	output           io.Writer
	input            io.ByteReader
}

type program struct {
	instructions []instruction
	source       io.ByteReader
	state        state
}

func (p *program) getNextInstructionToEval() (instruction, error) {
	if p.state.instructionIndex < len(p.instructions) {
		next := p.instructions[p.state.instructionIndex]
		return next, nil
	}
	for {
		r, err := p.source.ReadByte()
		if err != nil {
			return nil, err
		}
		next, ok := instructionMap[r]
		if ok {
			p.instructions = append(p.instructions, next)
			return next, nil
		}
	}
}

// Run starts interpretation process. It takes tree arguments:
// input -> io.Reader that takes user input
// output -> io.Writer to which all program output is passed
// source -> io.Reader which is used by interpreter to read program instractions
func Run(input io.Reader, output io.Writer, source io.Reader) error {
	initialState := state{
		dataPointer:      0,
		stack:            jumpStack{},
		instructionIndex: 0,
		skipBlock:        false,
		output:           output,
		input:            bufio.NewReader(input),
	}

	p := program{
		state:  initialState,
		source: bufio.NewReader(source),
	}

	for {
		instruction, err := p.getNextInstructionToEval()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("can't read instraction #%v: %v", p.state.instructionIndex+1, err)
		}

		err = instruction(&p.state)
		if err != nil {
			return fmt.Errorf("can't evaluate instraction #%v: %v", p.state.instructionIndex+1, err)
		}
		p.state.instructionIndex++
	}
}

type consoleWriter struct {
	w io.Writer
}

func (c consoleWriter) Write(b []byte) (int, error) {
	value := fmt.Sprintf("%c (ASCII code %v)", b[0], b[0])
	line := fmt.Sprint("bf-output: ", value, "\n")
	return c.w.Write([]byte(line))
}

func main() {
	err := Run(os.Stdin, consoleWriter{w: os.Stdout}, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}
