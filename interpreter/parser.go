package interpreter

import "fmt"

// Parser - takes a list of tokens and interpretes the relevant code
func Parser(tokenlist []string) {
	for i := range tokenlist {
		if tokenlist[i] == `print` {
			updatedList := tokenlist[i+1]
			str := updatedList[5:][:len(updatedList)-6]
			fmt.Println(str)
			i += 2
		}
	}
}
