package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/gcoka/monkey/lexer"
	"github.com/gcoka/monkey/token"
)

// PROMPT is a prompt string.
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%v\n", tok)
		}
	}
}
