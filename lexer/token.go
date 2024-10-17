package lexer

import (
	"errors"
	"log"
	"strings"
)


type TokenType int

var (
    InvalidToken = errors.New("invalid token")
)

const(
    LET TokenType = iota
    EQUAL
    NUMBER
    // OPERATORS
    PLUS
    MINUS
    // SPEACIAL SINGS
   OPEN_PAREN 
   CLOSE_PAREN 
   IDENTIFIER
)

type Token struct{
    T    TokenType 
    Value   string
}

func TokenKindString(token TokenType) string{
    switch token {
    case LET:
        return "let"
    case EQUAL:
        return "="
    case NUMBER:
        return "number"
    case OPEN_PAREN:
        return "("
    case CLOSE_PAREN:
        return ")"
    case PLUS:
        return "+"
    case MINUS:
        return "-"
    case IDENTIFIER:
        return "identifier"
    default:
        return "" 
    }
}



func (t Token) Debug(){
    log.Printf("Token kind: %s. Value: %s", TokenKindString(t.T), t.Value) 
}


func Tokenize(code string) []Token {
    trimmed := trim(code) 
	words := strings.Fields(trimmed)
    tokens := []Token{}
    for _, w := range words{
        ok, t := IsToken(w)
        if !ok{
            log.Println("invalid token: ", w)
            continue
        } 
        token := Token{
            T: t,
            Value: w,
        }
        tokens = append(tokens, token)
    }
    return tokens
}

func trim(code string) string  {
    return strings.TrimSpace(code) 
}
