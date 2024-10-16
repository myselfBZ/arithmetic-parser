package lexer

import (
	"errors"
	"log"
	"strconv"
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
    default:
        return "" 
    }
}

func isNumber(s string) bool {
    _, err :=  strconv.Atoi(s)
    return err == nil
}

func IsToken(token string) (bool, TokenType) {
    if isNumber(token) {
        return true, Tokens["number"]
    }
    if v, ok := Tokens[token]; ok{
        return true, v
    }
    return false, 0 
}

func (t Token) Debug(){
    log.Printf("Token kind: %s. Value: %s", TokenKindString(t.T), t.Value) 
}

/* i probably have to use regex
func specialSymbol(s string) (string, bool) {
    for _, s := range s{
        strS := string(s) 
        if strS == "=" || strS == "(" || strS == ")" || strS == "+" || strS == "-"{
            return strS, true
        }
    }
    return strS, false 
}
*/

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
