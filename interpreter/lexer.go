package interpreter

// Lexer - struct that analyzes source code and generates a language dictionary
type Lexer struct {
	tokens []string
}

// NewLexer - Initializes a new Lexer
func NewLexer() *Lexer {
	return &Lexer{tokens: []string{}}
}

// BuildDictionary -  does the heavy lifting of converting source code to tokens
func (lexer *Lexer) BuildDictionary(data []byte) {
	lineToken := ``
	tokenState := 0
	tree := ``
	for i := range data {
		lineToken += string(data[i])
		if lineToken == ` ` {
			if tokenState == 0 {
				lineToken = ``

			} else {
				lineToken = ` `
			}
		} else if lineToken == "\n" {
			lineToken = ``
		} else if lineToken == "print" {
			lexer.tokens = append(lexer.tokens, "print")
			lineToken = ``
		} else if lineToken == `"` {
			if tokenState == 0 {
				tokenState = 1
			} else if tokenState == 1 {
				lexer.tokens = append(lexer.tokens, "STR:"+tree+`"`)
				tree = ``
				tokenState = 0
				lineToken = ``
			}
		} else if tokenState == 1 {
			tree += lineToken
			lineToken = ``
		}
	}

}

//GetTokens - returns the tokens generated from our program
func (lexer *Lexer) GetTokens() []string {
	return lexer.tokens

}
