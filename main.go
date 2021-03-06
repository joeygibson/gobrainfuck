package main

import (
	"fmt"
	"github.com/joeygibson/gobrainfuck/machine"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]

	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	compiler := machine.NewCompiler(string(code))
	instructions := compiler.Compile()

	m := machine.NewMachine(instructions, os.Stdin, os.Stdout)
	m.Execute()
}
