package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	IDENT  TokenType = "IDENT"
	INT    TokenType = "INT"
	STRING TokenType = "STRING"

	TRUE  TokenType = "TRUE"
	FALSE TokenType = "FALSE"

	IF     TokenType = "IF"
	ELSE   TokenType = "ELSE"
	RETURN TokenType = "RETURN"

	ASSIGN TokenType = "="
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="
	PLUS   TokenType = "+"
	MINUS  TokenType = "-"
	TIMES  TokenType = "*"
	SLASH  TokenType = "/"
	BANG   TokenType = "!"
	LT     TokenType = "<"
	GT     TokenType = ">"

	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"

	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
