package parser

import "roll/src/token"

type Kind int

const (
	Number Kind = iota
	Add
	Sub
	Mul
	Div
	Dice
	KeepHigh
	KeepLow
	DropHigh
	DropLow
)

func FromTokenKind(kind token.Kind) Kind {
	switch kind {
	case token.Number:
		return Number
	case token.Add:
		return Add
	case token.Sub:
		return Sub
	case token.Mul:
		return Mul
	case token.Div:
		return Div
	case token.KeepHigh:
		return KeepHigh
	case token.KeepLow:
		return KeepLow
	case token.DropHigh:
		return DropHigh
	case token.DropLow:
		return DropLow
	case token.Dice:
		return Dice
	case token.LParen:
		panic("unexpected LParen kind")
	case token.RParen:
		panic("unexpected RParen kind")
	case token.EOF:
		panic("unexpected EOF kind")
	case token.Error:
		panic("unexpected Error kind")
	}
	panic("unreachable")
}

func (n Kind) String() string {
	switch n {
	case Number:
		return "Number"
	case Add:
		return "Add"
	case Sub:
		return "Sub"
	case Mul:
		return "Mul"
	case Div:
		return "Div"
	case Dice:
		return "Dice"
	case KeepHigh:
		return "KeepHigh"
	case KeepLow:
		return "KeepLow"
	case DropHigh:
		return "DropHigh"
	case DropLow:
		return "DropLow"
	}
	return "Unknown  Kind"
}
