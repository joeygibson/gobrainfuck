package machine

type Compiler struct {
	code       string
	codeLength int
	position   int

	instructions []*Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler{code: code,
		codeLength:   len(code),
		instructions: []*Instruction{},
	}
}

func (c *Compiler) Compile() []*Instruction {
	loopStack := []int{}

	for c.position < c.codeLength {
		current := c.code[c.position]

		switch current {
		case '+':
			c.CompileFoldableInstruction('+', Plus)
		case '-':
			c.CompileFoldableInstruction('-', Minus)
		case '>':
			c.CompileFoldableInstruction('>', Right)
		case '<':
			c.CompileFoldableInstruction('<', Left)
		case '.':
			c.CompileFoldableInstruction('.', PutChar)
		case ',':
			c.CompileFoldableInstruction(',', ReadChar)
		case '[':
			insPos := c.EmitWithArg(JumpIfZero, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			openInstruction := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]

			closeInstructionPos := c.EmitWithArg(JumpIfNotZero, openInstruction)
			c.instructions[openInstruction].Argument = closeInstructionPos
		}

		c.position++
	}

	return c.instructions
}

func (c *Compiler) CompileFoldableInstruction(char byte, instType InstType) {
	count := 1

	for c.position < c.codeLength-1 && c.code[c.position+1] == char {
		count++

		c.position++
	}

	c.EmitWithArg(instType, count)
}

func (c *Compiler) EmitWithArg(instType InstType, arg int) int {
	ins := &Instruction{Type: instType, Argument: arg}

	c.instructions = append(c.instructions, ins)

	return len(c.instructions) - 1
}
