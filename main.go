// name your module similar to namespacing in java and other languages
package main

import (
	"fmt"
	"gotraining/interpreter"
	"io/ioutil"
	"sync"
	"time"
)

func printRandomly(wg *sync.WaitGroup) {
	defer wg.Done()
	var i = 1
	for i <= 20000 {
		fmt.Println(`Printing line: `, i)
		i++
	}
}

// main func is run when you fire go run main.go
func main() {
	// lets measure how long this program took to execute
	var start time.Time
	start = time.Now()

	// defering this function first ensures it'll always execute last
	defer fmt.Println(time.Since(start))

	data, err := ioutil.ReadFile("program.basic")
	if err != nil {
		fmt.Println(err)
	}

	lexer := interpreter.NewLexer()
	lexer.BuildDictionary(data)

	// we want to parser our tokens second to last
	defer interpreter.Parser(lexer.GetTokens())

	// wait groups keep the go program alive until all concurrent operations are done
	var wg sync.WaitGroup

	// execute printRandomly 50 times i.e. 50 * 20000 loop iterations
	// i.e.  1 000 000 for loop iterations
	// i.e.  1 000 000 print statements
	// we howver are going to do this concurrently
	i := 1
	for i <= 50 {
		wg.Add(1)
		go printRandomly(&wg)
		i++
	}

	// hold program until every single goroutine is finished
	wg.Wait()

}
