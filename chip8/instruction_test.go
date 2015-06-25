package chip8

import "testing"

type constantTestCase struct {
	name             string
	actual, expected Instruction
}

func TestConstantInstructionsHaveCorrectValues(t *testing.T) {
	tests := []constantTestCase{
		{"cls", ClearScreen, Instruction(0x00E0)},
		{"ret", Return, Instruction(0x00EE)},
	}

	for _, test := range tests {
		if test.actual != test.expected {
			t.Error(test.name, "was", test.actual)
		}
	}
}

func TestInstructionsWithParametersHaveCorrectValues(t *testing.T) {
	tests := []constantTestCase{
		{"RCP", CallRcp(0x123), Instruction(0x0123)},
		{"jump", JumpTo(0x123), Instruction(0x1123)},
		{"call", Call(0x0444), Instruction(0x2444)},
		{"skip if equal", SkipIfEqual(5, 0xAB), Instruction(0x35AB)},
		{"skip if not equal", SkipIfNotEqual(5, 0xAB), Instruction(0x45AB)},
		{"skip if registers equal", SkipIfRegistersEqual(8, 9), Instruction(0x5890)},
		{"set register", SetRegisterTo(10, 0xFC), Instruction(0x6AFC)},
		{"inc", IncrementRegisterBy(11, 0x12), Instruction(0x7B12)},
		{"copy register", CopyRegister(1, 2), Instruction(0x8120)},
		{"or", OrRegister(3, 7), Instruction(0x8371)},
		{"and", AndRegister(3, 7), Instruction(0x8372)},
		{"xor", XorRegister(3, 7), Instruction(0x8373)},
		{"add", AddRegister(3, 7), Instruction(0x8374)},
		{"sub", SubtractRegister(3, 7), Instruction(0x8375)},
		{"sub reversed", SubtractReversed(3, 7), Instruction(0x8377)},
		{"shr", ShiftRight(3), Instruction(0x8306)},
		{"shl", ShiftLeft(3), Instruction(0x830E)},
		{"skip register unequal", SkipIfRegistersUnequal(3, 7), Instruction(0x9370)},
		{"set adr", SetAddressRegisterTo(0x321), Instruction(0xA321)},
		{"jump offset", JumpWithOffset(0xABC), Instruction(0xBABC)},
		{"rand", RandomizeAnd(3, 0xBB), Instruction(0xC3BB)},
		{"sprite", DrawSprite(1, 2, 3), Instruction(0xD123)},
		{"skip key down", SkipIfKeyDown(1), Instruction(0xE19E)},
		{"skip key up", SkipIfKeyUp(2), Instruction(0xE2A1)},
		{"delay timer", LoadDelayTimerInto(3), Instruction(0xF307)},
		{"wait key", WaitForKeyPress(4), Instruction(0xF40A)},
		{"set delay", SetDelayTimer(9), Instruction(0xF915)},
		{"set sound", SetSoundTimer(7), Instruction(0xF718)},
		{"inc adr", IncrementAddressRegister(3), Instruction(0xF31E)},
		{"adr to sprite", LoadDigitSprite(1), Instruction(0xF129)},
		{"bcd", LoadDecimalsOf(5), Instruction(0xF533)},
		{"copy to mem", CopyToMemory(0xA), Instruction(0xFA55)},
		{"load from mem", LoadFromMemory(0xA), Instruction(0xFA65)},
	}

	for _, test := range tests {
		if test.actual != test.expected {
			t.Error(test.name, "was", test.actual)
		}
	}
}
