package token

type Kind int

const (
	Number Kind = iota
	Add
	Sub
	Mul
	Div
	KeepHigh
	KeepLow
	DropHigh
	DropLow
	Dice
	LParen
	RParen
	EOF
	Error
)

func getKind(text string) Kind {
	switch text {
	case "+":
		return Add
	case "-":
		return Sub
	case "*":
		return Mul
	case "/":
		return Div
	case "d":
		return Dice
	case "kh":
		return KeepHigh
	case "kl":
		return KeepLow
	case "dh":
		return DropHigh
	case "dl":
		return DropLow
	case "(":
		return LParen
	case ")":
		return RParen
	case "":
		return EOF
	}

	if isNumber(text) {
		return Number
	}

	return Error
}

func (t Kind) String() string {
	switch t {
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
	case LParen:
		return "LParen"
	case RParen:
		return "RParen"
	case EOF:
		return "EOF"
	case Error:
		return "Error"
	}
	return "Unknown"
}
