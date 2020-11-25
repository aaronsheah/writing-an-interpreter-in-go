package main

import (
	"os"

	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
