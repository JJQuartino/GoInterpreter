package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  //Current position in input (points to current char)
	readPosition int  //Current reading position in input (after current char)
	char         byte //Current char under examination
}

/*
Both position and readPosition are used as indexes, e.g.: l.input[l.readPosition].
They point into the input string so we can "peek" futher into our input.
readPoisition always points to the "next" character in the input and position
points to the character in the input that corresponds to the ch byte.
*/

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// This is a receiver. It's used for declaring methods.

func (l *Lexer) readChar() {
	/*
		TODO: make readChar() fully support Unicode.
	*/
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
