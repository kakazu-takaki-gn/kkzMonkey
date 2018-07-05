package repl

import (
	"io"
	"bufio"
	"fmt"
	"github.com/kakazu-takaki-gn/kkzMonkey/lexer"
	"github.com/kakazu-takaki-gn/kkzMonkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.WriteCloser) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
