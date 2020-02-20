package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	var output bytes.Buffer
	source := strings.NewReader(",>++++++[<-------->-],[<+>-]<.")
	parameters := strings.NewReader("27")
	expected := "9"

	err := Run(parameters, &output, source)

	if err != nil {
		t.Errorf("error not expected - %v", err)
	}
	result := output.String()
	if result != expected {
		t.Errorf("Expected value=%v, received %v", expected, result)
	}
}

func TestPrintHelloWorld(t *testing.T) {
	var output bytes.Buffer
	source := strings.NewReader("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.")
	parameters := strings.NewReader("")
	expected := "Hello World!"

	err := Run(parameters, &output, source)

	if err != nil {
		t.Errorf("error not expected - %v", err)
	}
	result := output.String()
	if result != expected {
		t.Errorf("Expected value=%v, received %v", expected, result)
	}
}

func TestAddNumbersWithSimpleLoop(t *testing.T) {
	var output bytes.Buffer
	source := strings.NewReader(",>,<[->+<]>.")
	parameters := strings.NewReader("01")
	expected := "a"

	err := Run(parameters, &output, source)

	if err != nil {
		t.Errorf("error not expected - %v", err)
	}
	result := output.String()
	if result != expected {
		t.Errorf("Expected value=%v, received %v", expected, result)
	}
}

func TestNestedLoopsWithDivision(t *testing.T) {
	code := ",>,>++++++[-<--------<-------->>]<<[>[->+>+<<]>[-<<-[>]>>>[<[>>>-<<<[-]]>>]<<]>>>+<<[-<<+>>]<<<]>[-]>>>>[-<<<<<+>>>>>]<<<<++++++[-<++++++++>]<."

	testCases := []struct {
		parameters string
		expected   string
	}{
		{"84", "2"},
		{"93", "3"},
		{"82", "4"},
	}
	for _, tc := range testCases {
		t.Run(tc.parameters, func(t *testing.T) {
			output := new(bytes.Buffer)
			parameters := strings.NewReader(tc.parameters)
			source := strings.NewReader(code)

			err := Run(parameters, output, source)
			if err != nil {
				t.Errorf("error not expected - %v", err)
			}
			result := output.String()
			if result != tc.expected {
				t.Errorf("Expected value=%v, received %v", tc.expected, result)
			}
		})
	}
}

func TestErrorOnMissingOpenningTag(t *testing.T) {
	var output bytes.Buffer
	source := strings.NewReader("]")
	parameters := strings.NewReader("")

	err := Run(parameters, &output, source)

	if err == nil {
		t.Errorf("error was expected")
	}
}
