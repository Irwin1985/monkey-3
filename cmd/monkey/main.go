package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/gcoka/monkey/evaluator"
	"github.com/gcoka/monkey/lexer"
	"github.com/gcoka/monkey/object"
	"github.com/gcoka/monkey/parser"
	"github.com/gcoka/monkey/repl"
)

func main() {

	if len(os.Args) == 1 {
		startRepl()
	} else {
		os.Exit(runScript(os.Args[1]))
	}

}

func startRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey Programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func runScript(filepath string) (exitcode int) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	l := lexer.New(string(f))
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	evaluator.Eval(program, env)

	return 0
}
