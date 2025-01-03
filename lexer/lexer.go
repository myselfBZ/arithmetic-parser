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
    MULTIPLY
    DIVIDE
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
    case "*":
        return &Token{"*", MULTIPLY}
    case "/":
        return &Token{"/", DIVIDE}
    default:
        return &Token{string(current), INVALID}

    }
}

func (l *Lexer) Tokenize() []*Token {
    var tokens []*Token
    for{
        token := l.NextToken()
        if token.Type == INVALID || token.Type == -1{
            break
        }
        tokens = append(tokens, token)
        fmt.Printf("Token Type: %v Token Value: %v\n", token.Type, token.Value)
    }
    return tokens
}
