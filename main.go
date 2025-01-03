package main

import (
	"log"
	"os"
	"strings"

	"github.com/myselfBZ/plang/lexer"
	// "github.com/myselfBZ/plang/lexer"
)

func open(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("couldn't find the file")
	}
	buff := make([]byte, 1024)
	size, err := file.Read(buff)
	if err != nil {
		log.Fatal("error reading from a file, ", err)
	}
	return string(buff[:size])
}

func main() {
    src := open("test.pl")
    trimmedSrc := strings.TrimSpace(src)
    l := lexer.NewLexer(trimmedSrc)
    l.Tokenize()
}
