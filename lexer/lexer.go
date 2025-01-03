package lexer

import (
	"fmt"
	"unicode"
)

const SPACE = 32

type TokenType int

const(
    NUMBER TokenType = iota
    PLUS 
    MINUS
    INVALID
)

type Token struct{
    Value string 
    Type  TokenType
}

func NewLexer(input string) *Lexer{
    return &Lexer{input, 0}
}

type Lexer struct {
	input string
	pos   int
}

func(l *Lexer) NextToken() *Token {
    if l.pos >= len(l.input){
        fmt.Println("end of the input")
        return &Token{"", -1}
    }
    for l.pos < len(l.input) && unicode.IsSpace(rune(l.input[l.pos])) {
        l.pos++
    }
    current := l.input[l.pos]

    var number string

    if unicode.IsDigit(rune(current)){
        for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])){
            number += string(l.input[l.pos])
            l.pos++
        }
        return &Token{number, NUMBER}
    }
    l.pos++
    switch string(current){
    case "+":
        return &Token{"+", PLUS}
    case "-":
        return &Token{"-", MINUS}
    default:
        return &Token{string(current), INVALID}

    }
}

// example 3 + 4

func (l *Lexer) Tokenize() () {
    // var tokens []Token
    for{
        token := l.NextToken()
        if token.Type == INVALID{
            fmt.Println("Invalid Token: ", token.Value)
            break
        }
        if token.Value == "" {
            break
        }
        fmt.Printf("Token Type: %v Token Value: %v\n", token.Type, token.Value)
    }
}
