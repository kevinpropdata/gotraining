// name your module similar to namespacing in java and other languages
package main

import (
	"fmt"
	"io/ioutil"

	"./interpreter"
)

// main func is run when you fire go run main.go
func main() {
	data, err := ioutil.ReadFile("program.basic")
	if err != nil {
		fmt.Println(err)
	}

	lexer := interpreter.NewLexer()
	lexer.BuildDictionary(data)
	interpreter.Parser(lexer.GetTokens())

}
