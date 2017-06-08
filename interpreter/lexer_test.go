package interpreter

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

// TestLexer - will test our lexer logic
func TestLexer(t *testing.T) {
	data, err := ioutil.ReadFile("../program.basic")
	if err != nil {
		fmt.Println(err)
	}

	lexer := NewLexer()
	lexer.BuildDictionary(data)

	// store parsed tokens to use as our testcase
	testCase := lexer.GetTokens()

	// this is exactly what we expect from the lexer
	// this 2 line declare is not needed but an example of how to manually declare data types of variables
	var expectedResult []string
	expectedResult = []string{`print`, `STR:"Hello World"`, `print`, `STR:"Simple PRINT statement"`}

	// now compare the two splices
	if reflect.DeepEqual(testCase, expectedResult) != true {
		t.Error(`Error: Lexer was not able to accurately decode source tree`)
	}

}
