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

	l.skipWhitespace()

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
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

/*
Reads in an identifier and advances our lexer’s positions until it encounters a non-letter-character
*/
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	//TODO: generalize this by passing in the character-identifying functions as arguments.
	//TODO: make Monkey suport floats, hex and octal notation optional.
	positon := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[positon:l.position]
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
