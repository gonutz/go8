package chip8

type Instruction uint16

const (
	ClearScreen Instruction = 0x00E0
	Return                  = 0x00EE
)

func CallRcp(address uint16) Instruction {
	return Instruction(address)
}

func JumpTo(address uint16) Instruction {
	return Instruction(address + 1<<12)
}

func Call(address uint16) Instruction {
	return Instruction(address + 2<<12)
}

func SkipIfEqual(register, equalTo uint16) Instruction {
	return Instruction(3<<12 + register<<8 + equalTo)
}

func SkipIfNotEqual(register, unequalTo uint16) Instruction {
	return Instruction(4<<12 + register<<8 + unequalTo)
}

func SkipIfRegistersEqual(r1, r2 uint16) Instruction {
	return Instruction(5<<12 + r1<<8 + r2<<4)
}

func SetRegisterTo(register, value uint16) Instruction {
	return Instruction(6<<12 + register<<8 + value)
}

func IncrementRegisterBy(register, increment uint16) Instruction {
	return Instruction(7<<12 + register<<8 + increment)
}

func CopyRegister(to, from uint16) Instruction {
	return Instruction(8<<12 + to<<8 + from<<4)
}

func OrRegister(to, or uint16) Instruction {
	return Instruction(8<<12 + to<<8 + or<<4 + 1)
}

func AndRegister(to, and uint16) Instruction {
	return Instruction(8<<12 + to<<8 + and<<4 + 2)
}

func XorRegister(to, xor uint16) Instruction {
	return Instruction(8<<12 + to<<8 + xor<<4 + 3)
}

func AddRegister(to, plus uint16) Instruction {
	return Instruction(8<<12 + to<<8 + plus<<4 + 4)
}

func SubtractRegister(to, minus uint16) Instruction {
	return Instruction(8<<12 + to<<8 + minus<<4 + 5)
}

func SubtractReversed(to, left uint16) Instruction {
	return Instruction(8<<12 + to<<8 + left<<4 + 7)
}

func ShiftRight(which uint16) Instruction {
	return Instruction(8<<12 + which<<8 + 6)
}

func ShiftLeft(which uint16) Instruction {
	return Instruction(8<<12 + which<<8 + 0xE)
}

func SkipIfRegistersUnequal(r1, r2 uint16) Instruction {
	return Instruction(9<<12 + r1<<8 + r2<<4)
}

func SetAddressRegisterTo(value uint16) Instruction {
	return Instruction(0xA<<12 + value)
}

func JumpWithOffset(to uint16) Instruction {
	return Instruction(0xB<<12 + to)
}

func RandomizeAnd(register, and uint16) Instruction {
	return Instruction(0xC<<12 + register<<8 + and)
}

func DrawSprite(xRegister, yRegister, bytes uint16) Instruction {
	return Instruction(0xD<<12 + xRegister<<8 + yRegister<<4 + bytes)
}

func SkipIfKeyDown(which uint16) Instruction {
	return Instruction(0xE09E + which<<8)
}

func SkipIfKeyUp(which uint16) Instruction {
	return Instruction(0xE0A1 + which<<8)
}

func LoadDelayTimerInto(register uint16) Instruction {
	return Instruction(0xF007 + register<<8)
}

func WaitForKeyPress(which uint16) Instruction {
	return Instruction(0xF00A + which<<8)
}

func SetDelayTimer(to uint16) Instruction {
	return Instruction(0xF015 + to<<8)
}

func SetSoundTimer(to uint16) Instruction {
	return Instruction(0xF018 + to<<8)
}

func IncrementAddressRegister(by uint16) Instruction {
	return Instruction(0xF01E + by<<8)
}

func LoadDigitSprite(digit uint16) Instruction {
	return Instruction(0xF029 + digit<<8)
}

func LoadDecimalsOf(register uint16) Instruction {
	return Instruction(0xF033 + register<<8)
}

func CopyToMemory(upTo uint16) Instruction {
	return Instruction(0xF055 + upTo<<8)
}

func LoadFromMemory(upTo uint16) Instruction {
	return Instruction(0xF065 + upTo<<8)
}
