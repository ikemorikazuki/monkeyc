package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkeyc/cmd/evaluator"
	"monkeyc/cmd/lexer"
	"monkeyc/cmd/object"
	"monkeyc/cmd/parser"
)

const PROMT = "\x1b[31;1mmonkey\x1b[0m:>> "
const MONKEY_FACE = `
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for count := 1; ; count++ {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		if line == ":quit" || line == ":q" {
			break
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, fmt.Sprintf("\x1b[34;1mres%d\x1b[0m: \x1b[32;1m%s\x1b[0m = ", count, evaluated.Type()))
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "woops! we ran into some monkey business here!\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
