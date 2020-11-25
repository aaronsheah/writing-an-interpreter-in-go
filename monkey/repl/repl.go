package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/lexer"
	"github.com/aaronsheah/writing-an-interpreter-in-go/monkey/token"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for currentToken := l.NextToken(); currentToken.Type != token.EOF; currentToken = l.NextToken() {
			fmt.Printf("%+v\n", currentToken)
		}
	}
}
