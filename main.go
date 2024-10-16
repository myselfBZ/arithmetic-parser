package main

import (
	"log"
	"os"
    "github.com/myselfBZ/plang/lexer"
)



func open(path string) string {
    file, err := os.Open(path) 
    if err != nil{
        log.Fatal("couldn't find the file")
    }
    buff := make([]byte, 1024)
    size, err := file.Read(buff)
    if err != nil{
        log.Fatal("error reading from a file, ", err)
    }
    return string(buff[:size])
}

func main() {
    code := open("test.pl") 
    tokens := lexer.Tokenize(code)
    analsis := map[string]int{}
    for _, t := range tokens{
        kind := lexer.TokenKindString(t.T)     
        if _, ok := analsis[kind]; ok{
            analsis[kind]++
        } else {
            analsis[kind] = 1
        }
    }
    
    for v, k := range analsis{
        log.Printf(" Kind \"%s\" : %d", v, k)
    }
}


