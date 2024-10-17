package main

import (
	"log"
	"os"
    "github.com/myselfBZ/plang/parser"
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
    for i := 0; i < len(tokens); i+=4{
        node, err := parser.VarDeclare(tokens[i:i+4]) 
        if err != nil{
            log.Fatal(err)
        }
        log.Println(node)
    }
}


