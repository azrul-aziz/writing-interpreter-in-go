package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in) // creates a new scanner that will read input from the in parameter line by line

	for { // start an infinite loop
		fmt.Fprintf(out, PROMPT)  // writes the prompt string to the terminal to prompt user for input. out is the destination where the string will be written
		scanned := scanner.Scan() // reads the next line after the prompt, returns boolean value
		if !scanned {
			return // if there is no input, exit the loop
		}

		line := scanner.Text() // retrieves the most recently scanned line of input as a string and assign it to line variable
		l := lexer.New(line)   // initialize a new lexer struct using the given input

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
