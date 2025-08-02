package token

type TokenType string

type Token struct {
	Type    TokenType //Type, so we can distinguish from “integers” and “right bracket” for example.
	Literal string    //Holds the literal value of the token, so we can reuse it later and the information whether a “number” token is a 5 or a 10 doesn’t get lost.
}

const (
	//Special types
	ILLEGAL = "ILLEGAL" //Token/character we don't know about
	EOF     = "EOF"     //End Of File, tells the parser that it can stop

	//Identifiers and literals
	IDENT = "IDENT" //add, fn, x, y, etc.
	INT   = "INT"   //1234567890

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
