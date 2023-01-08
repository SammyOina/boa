package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/sammyoina/boa/evaluator"
	"github.com/sammyoina/boa/lexer"
	"github.com/sammyoina/boa/object"
	"github.com/sammyoina/boa/parser"
	"github.com/sammyoina/boa/repl"
	"github.com/spf13/cobra"
)

var version = " not set get, another version"

var rootCmd = &cobra.Command{
	Use:     "boa",
	Short:   "boa is an awesome programming language",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		env := object.NewEnv()
		if len(args) == 0 {
			user, err := user.Current()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Hello %s, Welcome to boa programming language!\n", user.Username)
			repl.Start(os.Stdin, os.Stdout)
		} else if len(args) == 1 {
			_, err := os.Stat(args[0])
			//path.Clean()
			if err != nil {
				log.Println(err)
				return
			}
			f, err := os.ReadFile(args[0])
			if err != nil {
				log.Println(err)
				return
			}
			l := lexer.New(string(f))
			p := parser.New(l)
			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				printParserErrors(p.Errors())
				return
			}
			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				fmt.Println(evaluated.Inspect())
				//fmt.Printf("\n")
			}
		} else {
			log.Printf("need one argument got %d\n", len(args))
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printParserErrors(errors []error) {
	log.Printf("Parser errors:\n")
	for _, msg := range errors {
		log.Printf("\t" + msg.Error() + "\n")
	}
}
