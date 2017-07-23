package machine

type InstType byte

const (
	Plus          InstType = '+'
	Minus         InstType = '-'
	Right         InstType = '>'
	Left          InstType = '<'
	PutChar       InstType = '.'
	ReadChar      InstType = ','
	JumpIfZero    InstType = '['
	JumpIfNotZero InstType = ']'
)

type Instruction struct {
	Type     InstType
	Argument int
}
