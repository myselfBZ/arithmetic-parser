package parser

import (
	"errors"
	"github.com/myselfBZ/plang/lexer"
)

var( 
    malformedExpression = errors.New("malformed expression")
    unexpectedToken = errors.New("unexpected token")
    useOfReservedKeyWord = errors.New("use of reserved keyword")
)



func VarDeclare(tokens []lexer.Token) (*ASTNode, error) {
    if  len(tokens) != 4 {
        return nil, malformedExpression
    }

    if len(tokens) < 4{
        return nil, malformedExpression 
    }

    if tokens[0].T  != lexer.LET{
        return  nil, unexpectedToken 
    }   

    if _, ok := lexer.Tokens[tokens[1].Value]; ok && tokens[1].T != lexer.IDENTIFIER { 
        return nil, useOfReservedKeyWord 
    }

    if tokens[2].T != lexer.EQUAL{
        return nil, unexpectedToken
    }

    return &ASTNode{
        Kind: "Assignment",
        Variable: tokens[1].Value,
        Value: tokens[3].Value,
    }, nil

}


