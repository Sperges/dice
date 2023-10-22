package token

import (
	"text/scanner"
	"unicode"
)

type Token struct {
	Text     string
	Kind     Kind
	Position scanner.Position
}

func (t *Token) IsOp() bool {
	switch t.Kind {
	case Number:
		return false
	case Add:
		return true
	case Sub:
		return true
	case Mul:
		return true
	case Div:
		return true
	case KeepHigh:
		return true
	case KeepLow:
		return true
	case DropHigh:
		return true
	case DropLow:
		return true
	case Dice:
		return true
	case LParen:
		return false
	case RParen:
		return false
	case EOF:
		return false
	case Error:
		return false
	}
	return false
}

func isNumber(text string) bool {
	for _, r := range text {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
