package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sammyoina/boa/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, Welcome to boa programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
