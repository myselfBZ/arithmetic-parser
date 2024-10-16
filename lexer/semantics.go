package lexer

var Tokens = map[string]TokenType{
    "let":LET,
    "=":EQUAL,
    "number":NUMBER,
    ")":CLOSE_PAREN,
    "(":OPEN_PAREN,
    "+":PLUS,
    "-":MINUS,
}
