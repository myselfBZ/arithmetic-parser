package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/myselfBZ/plang/lexer"
)

type Parser struct {
	tokens []*lexer.Token
	pos    int
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) Parse() (int, error){
    return p.parseExpression()
}

func(p *Parser) parseExpression() (int, error){
    left, err := p.parseFactor()
    if err != nil{
        return 0, err
    }

    for p.pos < len(p.tokens){
        current := p.tokens[p.pos]
        if current.Type != lexer.PLUS && current.Type != lexer.MINUS{
            break
        }

        p.pos++ 
        right, err := p.parseTerm()
        if err != nil{
            return 0, err
        }

        switch current.Type{
        case lexer.PLUS:
            left+=right
        case lexer.MINUS:
            left-=right
        }
    }
    return left, nil
}

func (p *Parser) parseFactor() (int, error) {
    left, err := p.parseTerm()
    if err != nil {
        return 0, err
    }

    for p.pos < len(p.tokens) {
        current := p.tokens[p.pos]

        if current.Type != lexer.MULTIPLY && current.Type != lexer.DIVIDE {
            break
        }

        p.pos++
        right, err := p.parseTerm() // Parse the next term
        if err != nil {
            return 0, err
        }

        switch current.Type {
        case lexer.MULTIPLY:
            left *= right
        case lexer.DIVIDE:
            if right == 0 {
                return 0, fmt.Errorf("division by zero")
            }
            left /= right
        }
    }

    return left, nil
}

func (p *Parser) parseTerm() (int, error) {
    if p.pos >= len(p.tokens){
        return 0, errors.New("malformed expression")
    }
	token := p.tokens[p.pos]
	p.pos++

	switch token.Type {
	case lexer.NUMBER:
		return strconv.Atoi(token.Value)
	default:
		return 0, fmt.Errorf("unexpected token: %v", token)
	}
}
