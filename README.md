## About

KatarzynaSk/brainfuck is a brainfuck interpreter library. It reads input from io.Reader without knowing all input at once.

## Install 

```
go get github.com/KatarzynaSk/brainfuck
```

## Usage 

### In console
To start an interactive session build the project and run it with go tools. The program takes both user input and source code from standard input and writes to standard output.

### In code
```
var output bytes.Buffer
source := strings.NewReader(",>++++++[<-------->-],[<+>-]<.")
input := strings.NewReader("27")

err := Run(input, &output, source)

if err != nil {
    log.Fatal(err)
}
result := output.String()
fmt.Println(result)
```

## Issues
Some common exeptions are not handled gracefully.

e.g.
* memory access out of bounds
* stack overflow 
