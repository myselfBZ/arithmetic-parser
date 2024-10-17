package lexer

import (
	"strconv"
)

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

func isNumber(s string) bool {
    _, err :=  strconv.Atoi(s)
    return err == nil
}


func IsToken(token string) ( bool,  TokenType) {
    if isNumber(token) {
        return true, Tokens["number"]
    } 
    if  IsIdentifier(token){
        return true, IDENTIFIER
    }
    if  v, ok := Tokens[token]; ok{
        return true, v
    } 
    return false, 0 
}

func IsIdentifier(s string) (bool) {
    if len(s) == 0 {
        return false 
    }

    if _, ok := Tokens[s]; ok{
        return false
    } 

    if !((s[0] >= 'a' && s[0] <= 'z') || (s[0] >= 'A' && s[0] <= 'Z')) {
        return false
    }

    for _, c := range s{
        if !((c >= 'a' && c <='z') || (c >= 'A' && c <='Z')){
            return false 
        }
    }

    return true
}
