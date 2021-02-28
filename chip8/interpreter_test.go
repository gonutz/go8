package chip8

import "testing"

var chip8 *Interpreter

func setUp() {
	chip8 = NewInterpreter()
	chip8.SetSpeakers(dummySpeakers{})
	chip8.SetKeyEvent(dummyKeyEvent{})
}

type dummySpeakers struct{}

func (dummySpeakers) Beep()    {}
func (dummySpeakers) BeQuiet() {}

type dummyKeyEvent struct{}

func (dummyKeyEvent) LastKey() (bool, uint8) { return false, 0 }
func (dummyKeyEvent) ClearLastKey()          {}

func TestAfterCreationAllRegistersAreZero(t *testing.T) {
	setUp()
	checkAllRegistersAreZero(t, chip8)
}

func checkAllRegistersAreZero(t *testing.T, chip8 *Interpreter) {
	for i := 0; i < 16; i++ {
		if chip8.Register(i) != uint8(0) {
			t.Error("register", i, "not 0")
		}
	}
	if chip8.AddressRegister() != 0 {
		t.Error("address register not 0")
	}
}

func TestAfterCreationTheProgramCounterIsAtHex200(t *testing.T) {
	setUp()
	pc := chip8.ProgramCounter()
	if pc != 0x200 {
		t.Error("pc was not 512 but", pc)
	}
}

func TestCallingRcpFunctionIsSkipped(t *testing.T) {
	setUp()
	chip8.Execute(CallRcp(123))
	checkAllRegistersAreZero(t, chip8)
	pc := chip8.ProgramCounter()
	if pc != 0x200+2 {
		t.Error("pc was", pc)
	}
}

func TestRegistersCanBeSetToValues(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(4, 123))
	if chip8.Register(4) != 123 {
		t.Error("register not set")
	}
	if chip8.ProgramCounter() != 0x200+2 {
		t.Error("wrong pc")
	}
}

func TestJumpingSetsProgramCounterToAddress(t *testing.T) {
	setUp()
	chip8.Execute(JumpTo(95))
	if chip8.ProgramCounter() != 95 {
		t.Error("wrong pc")
	}
}

func TestJumpWithOffsetJumpsToAddressPlusRegister0(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(0, 5))
	chip8.Execute(JumpWithOffset(550))
	if chip8.ProgramCounter() != 555 {
		t.Error("wrong pc", chip8.ProgramCounter())
	}
}

func TestInstructionCanBeSkippedIfRegisterEqualsValue(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(7, 77))
	pc := chip8.ProgramCounter()
	chip8.Execute(SkipIfEqual(7, 66))
	if chip8.ProgramCounter() != pc+2 {
		t.Error("was skipped")
	}

	pc = chip8.ProgramCounter()
	chip8.Execute(SkipIfEqual(7, 77))
	if chip8.ProgramCounter() != pc+4 {
		t.Error("was not skipped")
	}
}

func TestInstructionCanBeSkippedIfRegisterDoesNotEqualValue(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(7, 77))
	pc := chip8.ProgramCounter()
	chip8.Execute(SkipIfNotEqual(7, 66))
	if chip8.ProgramCounter() != pc+4 {
		t.Error("was not skipped")
	}

	pc = chip8.ProgramCounter()
	chip8.Execute(SkipIfNotEqual(7, 77))
	if chip8.ProgramCounter() != pc+2 {
		t.Error("was skipped")
	}
}

func TestInstructionCanBeSkippedIfTwoRegistersAreEqual(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(3, 33))
	chip8.Execute(SetRegisterTo(7, 77))
	chip8.Execute(SetRegisterTo(12, 77))

	pc := chip8.ProgramCounter()
	chip8.Execute(SkipIfRegistersEqual(3, 7))
	if chip8.ProgramCounter() != pc+2 {
		t.Error("was skipped")
	}

	pc = chip8.ProgramCounter()
	chip8.Execute(SkipIfRegistersEqual(7, 12))
	if chip8.ProgramCounter() != pc+4 {
		t.Error("was not skipped")
	}
}

func TestInstructionCanBeSkippedIfTwoRegistersAreUnequal(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(3, 33))
	chip8.Execute(SetRegisterTo(7, 77))
	chip8.Execute(SetRegisterTo(12, 77))

	pc := chip8.ProgramCounter()
	chip8.Execute(SkipIfRegistersUnequal(7, 12))
	if chip8.ProgramCounter() != pc+2 {
		t.Error("was skipped")
	}

	pc = chip8.ProgramCounter()
	chip8.Execute(SkipIfRegistersUnequal(3, 7))
	if chip8.ProgramCounter() != pc+4 {
		t.Error("was not skipped")
	}
}

func TestTheAddressRegisterCanBeSet(t *testing.T) {
	setUp()
	chip8.Execute(SetAddressRegisterTo(456))
	if chip8.AddressRegister() != 456 {
		t.Error("wrong address register value", chip8.AddressRegister())
	}
}

func TestRandomNumbersCanBeGenerated(t *testing.T) {
	setUp()
	chip8.SetRandomizer(sequenceRand(100, 200))

	chip8.Execute(RandomizeAnd(5, 20))
	if chip8.Register(5) != 100&20 {
		t.Error("wrong rand", chip8.Register(5))
	}

	chip8.Execute(RandomizeAnd(6, 30))
	if chip8.Register(6) != 200&30 {
		t.Error("wrong rand", chip8.Register(5))
	}
}

type sequence struct {
	stream  []uint8
	current int
}

func (s *sequence) Next() uint8 {
	s.current++
	return s.stream[s.current-1]
}

func sequenceRand(numbers ...uint8) *sequence {
	return &sequence{stream: numbers}
}

func TestInstructionCanBeSkippedIfKeyIsPressed(t *testing.T) {
	setUp()
	k := &stubKeyboard{}
	chip8.SetKeyboard(k)
	chip8.Execute(SetRegisterTo(5, 10))

	k.keys[10] = false
	pc := chip8.ProgramCounter()
	chip8.Execute(SkipIfKeyDown(5))
	if chip8.ProgramCounter() != pc+2 {
		t.Error("skipped")
	}

	k.keys[10] = true
	pc = chip8.ProgramCounter()
	chip8.Execute(SkipIfKeyDown(5))
	if chip8.ProgramCounter() != pc+4 {
		t.Error("not skipped")
	}
}

type stubKeyboard struct{ keys [16]bool }

func (k *stubKeyboard) KeyDown(which uint8) bool { return k.keys[which] }

func TestInstructionCanBeSkippedIfKeyIsUp(t *testing.T) {
	setUp()
	k := &stubKeyboard{}
	chip8.SetKeyboard(k)
	chip8.Execute(SetRegisterTo(7, 2))

	k.keys[2] = true
	pc := chip8.ProgramCounter()
	chip8.Execute(SkipIfKeyUp(7))
	if chip8.ProgramCounter() != pc+2 {
		t.Error("skipped")
	}

	k.keys[2] = false
	pc = chip8.ProgramCounter()
	chip8.Execute(SkipIfKeyUp(7))
	if chip8.ProgramCounter() != pc+4 {
		t.Error("not skipped")
	}
}

func TestSubroutinesCanBeRunAndReturnedFrom(t *testing.T) {
	setUp()
	original := chip8.ProgramCounter()
	chip8.Execute(Call(300))
	if chip8.ProgramCounter() != 300 {
		t.Error("pc not 300")
	}

	chip8.Execute(Call(400))
	if chip8.ProgramCounter() != 400 {
		t.Error("pc not 400")
	}

	chip8.Execute(Return)
	if chip8.ProgramCounter() != 300+2 {
		t.Error("pc not back after 400", chip8.ProgramCounter())
	}

	chip8.Execute(Return)
	if chip8.ProgramCounter() != original+2 {
		t.Error("pc not back after 300", original+2, chip8.ProgramCounter())
	}
}

func TestValueCanBeAddedToRegister(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(3, 50))
	chip8.Execute(IncrementRegisterBy(3, 40))
	if chip8.Register(3) != 90 {
		t.Error("wrong sum", chip8.Register(3))
	}
}

func TestIncrementingRegisterMayOverflow(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(3, 255))
	chip8.Execute(IncrementRegisterBy(3, 2))
	if chip8.Register(3) != 1 {
		t.Error("wrong sum", chip8.Register(3))
	}
}

func TestTwoRegistersCanBeAdded(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(0, 50))
	chip8.Execute(SetRegisterTo(1, 60))
	chip8.Execute(AddRegister(0, 1))
	if chip8.Register(0) != 110 {
		t.Error("wrong sum", chip8.Register(0))
	}
}

func TestWhenAddingTwoRegisterOverflow_CarryBitIsSet(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(0, 254))
	chip8.Execute(SetRegisterTo(1, 1))
	chip8.Execute(AddRegister(0, 1))
	if chip8.Register(0) != 255 || chip8.Register(0xF) != 0 {
		t.Error("added without overflow")
	}
	chip8.Execute(AddRegister(0, 1))
	if chip8.Register(0) != 0 || chip8.Register(0xF) != 1 {
		t.Error("added with overflow")
	}
	chip8.Execute(AddRegister(0, 1))
	if chip8.Register(0) != 1 || chip8.Register(0xF) != 0 {
		t.Error("added again without overflow")
	}
}

func TestRegisterCanBeCopiedToOtherRegister(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(15, 200))
	chip8.Execute(SetRegisterTo(4, 11))
	chip8.Execute(CopyRegister(4, 15))
	if chip8.Register(4) != 200 {
		t.Error("not copied", chip8.Register(4))
	}
}

func TestRegisterCanBeOrAssigned(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(1, 5))
	chip8.Execute(SetRegisterTo(2, 10))
	chip8.Execute(OrRegister(1, 2))
	if chip8.Register(1) != 5|10 {
		t.Error("not or'ed", chip8.Register(1))
	}
}

func TestRegisterCanBeAndAssigned(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(1, 5))
	chip8.Execute(SetRegisterTo(2, 10))
	chip8.Execute(AndRegister(1, 2))
	if chip8.Register(1) != 5&10 {
		t.Error("not and'ed", chip8.Register(1))
	}
}

func TestRegisterCanBeXorAssigned(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(1, 5))
	chip8.Execute(SetRegisterTo(2, 10))
	chip8.Execute(XorRegister(1, 2))
	if chip8.Register(1) != 5^10 {
		t.Error("not xor'ed", chip8.Register(1))
	}
}

func TestRegistersCanBeSubtractedWhichSetsCarryBit(t *testing.T) {
	setUp()

	chip8.Execute(SetRegisterTo(3, 20))
	chip8.Execute(SetRegisterTo(2, 40))
	chip8.Execute(SubtractRegister(3, 2))
	if chip8.Register(3) != 256-20 {
		t.Error("20-40 !=", chip8.Register(3))
	}
	if chip8.Register(15) != 0 {
		t.Error("carry flag should be 0", chip8.Register(15))
	}

	chip8.Execute(SetRegisterTo(3, 50))
	chip8.Execute(SetRegisterTo(2, 20))
	chip8.Execute(SubtractRegister(3, 2))
	if chip8.Register(3) != 30 {
		t.Error("50-20 !=", chip8.Register(3))
	}
	if chip8.Register(15) != 1 {
		t.Error("carry flag should be 1", chip8.Register(15))
	}
}

func TestXMinusXYieldsZeroAndNoCarryBit(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(4, 1))
	chip8.Execute(SetRegisterTo(5, 1))
	chip8.Execute(SubtractRegister(4, 5))
	if chip8.Register(4) != 0 {
		t.Error("1 - 1 was", chip8.Register(4))
	}
	if chip8.Register(15) != 1 {
		t.Error("wrong carry", chip8.Register(15))
	}
}

func TestRegistersCanBeReverselySubtractedWhichSetsCarryBit(t *testing.T) {
	setUp()

	chip8.Execute(SetRegisterTo(3, 40))
	chip8.Execute(SetRegisterTo(2, 20))
	chip8.Execute(SubtractReversed(3, 2))
	if chip8.Register(3) != 256-20 {
		t.Error("20-40 !=", chip8.Register(3))
	}
	if chip8.Register(15) != 0 {
		t.Error("carry flag should be 0", chip8.Register(15))
	}

	chip8.Execute(SetRegisterTo(3, 20))
	chip8.Execute(SetRegisterTo(2, 50))
	chip8.Execute(SubtractReversed(3, 2))
	if chip8.Register(3) != 30 {
		t.Error("50-20 !=", chip8.Register(3))
	}
	if chip8.Register(15) != 1 {
		t.Error("carry flag should be 1", chip8.Register(15))
	}
}

func TestAddressRegisterCanBeIncrementedByRegister(t *testing.T) {
	setUp()
	chip8.Execute(SetAddressRegisterTo(45))
	chip8.Execute(SetRegisterTo(3, 25))
	chip8.Execute(IncrementAddressRegister(3))
	if chip8.AddressRegister() != 70 {
		t.Error("wrong sum in address register", chip8.AddressRegister())
	}
}

func TestDigitSpriteLocationsPointToCorrectSprites(t *testing.T) {
	setUp()
	checkDigit(t, 0, 0xF0, 0x90, 0x90, 0x90, 0xF0)
	checkDigit(t, 1, 0x20, 0x60, 0x20, 0x20, 0x70)
	checkDigit(t, 2, 0xF0, 0x10, 0xF0, 0x80, 0xF0)
	checkDigit(t, 3, 0xF0, 0x10, 0xF0, 0x10, 0xF0)
	checkDigit(t, 4, 0x90, 0x90, 0xF0, 0x10, 0x10)
	checkDigit(t, 5, 0xF0, 0x80, 0xF0, 0x10, 0xF0)
	checkDigit(t, 6, 0xF0, 0x80, 0xF0, 0x90, 0xF0)
	checkDigit(t, 7, 0xF0, 0x10, 0x20, 0x40, 0x40)
	checkDigit(t, 8, 0xF0, 0x90, 0xF0, 0x90, 0xF0)
	checkDigit(t, 9, 0xF0, 0x90, 0xF0, 0x10, 0xF0)
	checkDigit(t, 10, 0xF0, 0x90, 0xF0, 0x90, 0x90)
	checkDigit(t, 11, 0xE0, 0x90, 0xE0, 0x90, 0xE0)
	checkDigit(t, 12, 0xF0, 0x80, 0x80, 0x80, 0xF0)
	checkDigit(t, 13, 0xE0, 0x90, 0x90, 0x90, 0xE0)
	checkDigit(t, 14, 0xF0, 0x80, 0xF0, 0x80, 0xF0)
	checkDigit(t, 15, 0xF0, 0x80, 0xF0, 0x80, 0x80)
}

func checkDigit(t *testing.T, digit uint16, expected ...uint8) {
	chip8.Execute(SetRegisterTo(0, digit))
	chip8.Execute(LoadDigitSprite(0))
	mem := chip8.AddressRegister()
	for i := uint16(0); i < 5; i++ {
		if chip8.MemoryByte(mem+i) != expected[i] {
			t.Error("wrong digit", digit, "at", i, "was", chip8.MemoryByte(mem+i))
		}
	}
}

func TestAllDigitsLieInAddressesBelowHex200(t *testing.T) {
	setUp()
	for i := uint16(0); i < 16; i++ {
		chip8.Execute(LoadDigitSprite(i))
		if chip8.AddressRegister()+4 >= 0x200 {
			t.Error("digit", i, "starts at address", chip8.AddressRegister())
		}
	}
}

func TestBinaryDecimalsCanBeStoredInMemory(t *testing.T) {
	setUp()
	chip8.Execute(SetAddressRegisterTo(500))
	chip8.Execute(SetRegisterTo(4, 235))
	chip8.Execute(LoadDecimalsOf(4))
	if chip8.MemoryByte(500) != 2 {
		t.Error("hundreds", chip8.MemoryByte(500))
	}
	if chip8.MemoryByte(501) != 3 {
		t.Error("tens", chip8.MemoryByte(501))
	}
	if chip8.MemoryByte(502) != 5 {
		t.Error("ones", chip8.MemoryByte(502))
	}
}

func TestRegistersCanBeWrittenToMemory(t *testing.T) {
	setUp()
	chip8.Execute(SetAddressRegisterTo(600))
	chip8.Execute(SetRegisterTo(0, 1))
	chip8.Execute(SetRegisterTo(1, 2))
	chip8.Execute(SetRegisterTo(2, 3))
	chip8.Execute(SetRegisterTo(3, 4))
	chip8.Execute(CopyToMemory(2))

	m1 := chip8.MemoryByte(600)
	m2 := chip8.MemoryByte(601)
	m3 := chip8.MemoryByte(602)
	m4 := chip8.MemoryByte(603)
	if m1 != 1 || m2 != 2 || m3 != 3 {
		t.Error("registers not copied", m1, m2, m3)
	}
	if m4 == 4 {
		t.Error("copied too much", m4)
	}
}

func TestRegistersCanBeReadFromMemory(t *testing.T) {
	setUp()
	chip8.memory[700] = 1
	chip8.memory[701] = 2
	chip8.memory[702] = 3
	chip8.memory[703] = 4
	chip8.Execute(SetAddressRegisterTo(700))
	chip8.Execute(LoadFromMemory(2))
	r1 := chip8.Register(0)
	r2 := chip8.Register(1)
	r3 := chip8.Register(2)
	r4 := chip8.Register(3)
	if r1 != 1 || r2 != 2 || r3 != 3 {
		t.Error("registers not loaded", r1, r2, r3)
	}
	if r4 == 4 {
		t.Error("loaded too much", r4)
	}
}

func TestDelayTimerCanBeSet(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(5, 10))
	chip8.Execute(SetDelayTimer(5))
	if chip8.DelayTimer() != 10 {
		t.Error("delay timer not set")
	}
}

func TestDelayTimerCanBeReadToRegister(t *testing.T) {
	setUp()
	chip8.delay = 20
	chip8.Execute(LoadDelayTimerInto(2))
	if chip8.Register(2) != 20 {
		t.Error("delay timer not read")
	}
}

func TestSoundTimerCanBeSet(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(6, 15))
	chip8.Execute(SetSoundTimer(6))
	if chip8.SoundTimer() != 15 {
		t.Error("sound timer not set")
	}
}

func TestOnTimerTickTimersAreDecrementedUntilZero(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(0, 3))
	chip8.Execute(SetRegisterTo(1, 2))
	chip8.Execute(SetDelayTimer(0))
	chip8.Execute(SetSoundTimer(1))

	chip8.TimerTick()
	if chip8.DelayTimer() != 2 {
		t.Error("delay timer not 2 but", chip8.DelayTimer())
	}
	if chip8.SoundTimer() != 1 {
		t.Error("sound timer not 1 but", chip8.SoundTimer())
	}

	chip8.TimerTick()
	if chip8.DelayTimer() != 1 {
		t.Error("delay timer not 1 but", chip8.DelayTimer())
	}
	if chip8.SoundTimer() != 0 {
		t.Error("sound timer not 0 but", chip8.SoundTimer())
	}

	chip8.TimerTick()
	if chip8.DelayTimer() != 0 {
		t.Error("delay timer not 0 but", chip8.DelayTimer())
	}
	if chip8.SoundTimer() != 0 {
		t.Error("sound timer not 0 but", chip8.SoundTimer())
	}

	chip8.TimerTick()
	if chip8.DelayTimer() != 0 {
		t.Error("delay timer not 0 but", chip8.DelayTimer())
	}
}

func TestRightShiftAlsoSetsFlag(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(3, 5))

	chip8.Execute(ShiftRight(3))
	if chip8.Register(3) != 2 {
		t.Error("not shifted", chip8.Register(3))
	}
	if chip8.Register(15) != 1 {
		t.Error("wrong carry flag", chip8.Register(15))
	}

	chip8.Execute(ShiftRight(3))
	if chip8.Register(3) != 1 {
		t.Error("not shifted", chip8.Register(3))
	}
	if chip8.Register(15) != 0 {
		t.Error("wrong carry flag", chip8.Register(15))
	}
}

func TestLeftShiftAlsoSetsFlag(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(4, 0xA0))

	chip8.Execute(ShiftLeft(4))
	if chip8.Register(4) != 0x40 {
		t.Error("not shifted", chip8.Register(4))
	}
	if chip8.Register(15) != 1 {
		t.Error("wrong carry flag", chip8.Register(15))
	}

	chip8.Execute(ShiftLeft(4))
	if chip8.Register(4) != 0x80 {
		t.Error("not shifted", chip8.Register(4))
	}
	if chip8.Register(15) != 0 {
		t.Error("wrong carry flag", chip8.Register(15))
	}
}

func TestInitiallyNoScreenPixelIsSet(t *testing.T) {
	setUp()
	checkOtherPixelsAreNotSet(t, []point{})
}

type point struct{ x, y int }

func TestSpriteCanBeDrawnFromMemory(t *testing.T) {
	setUp()
	chip8.memory[700] = 0x11 // 00010001
	chip8.memory[701] = 0x0D // 00001101
	chip8.memory[702] = 0x31 // 00110001
	chip8.Execute(SetAddressRegisterTo(700))
	var x, y uint16 = 25, 2
	chip8.Execute(SetRegisterTo(1, x))
	chip8.Execute(SetRegisterTo(2, y))
	chip8.Execute(DrawSprite(1, 2, 3))

	set := []point{
		{28, 2}, {32, 2},
		{29, 3}, {30, 3}, {32, 3},
		{27, 4}, {28, 4}, {32, 4},
	}
	checkPixelsAreSet(t, set)
	checkOtherPixelsAreNotSet(t, set)
}

func checkPixelsAreSet(t *testing.T, set []point) {
	for _, p := range set {
		if !chip8.Screen().IsSet(p.x, p.y) {
			t.Fatal(p.x, p.y, "not set")
		}
	}
}

func checkOtherPixelsAreNotSet(t *testing.T, set []point) {
	w, h := chip8.Screen().Size()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if !isInSet(set, x, y) && chip8.Screen().IsSet(x, y) {
				t.Fatal(x, y, "was set")
			}
		}
	}
}

func isInSet(set []point, x, y int) bool {
	for _, p := range set {
		if x == p.x && y == p.y {
			return true
		}
	}
	return false
}

func TestDrawingSpriteXorsThePixels(t *testing.T) {
	setUp()
	chip8.memory[700] = 0x11
	chip8.memory[701] = 0x0D
	chip8.memory[702] = 0x31
	chip8.Execute(SetAddressRegisterTo(700))
	var x, y uint16 = 25, 2
	chip8.Execute(SetRegisterTo(1, x))
	chip8.Execute(SetRegisterTo(2, y))

	chip8.Execute(DrawSprite(1, 2, 3))
	chip8.Execute(DrawSprite(1, 2, 3))

	checkOtherPixelsAreNotSet(t, []point{})
}

func TestDrawingSpriteSetsCollisionFlagIfAnyPixelWasErased(t *testing.T) {
	setUp()
	chip8.memory[700] = 0x11 // 00010001
	chip8.memory[701] = 0x0C // 00001100
	chip8.memory[702] = 0x04 // 00000100
	chip8.memory[703] = 0x80 // 10000000
	var x, y uint16 = 25, 2
	chip8.Execute(SetRegisterTo(1, x))
	chip8.Execute(SetRegisterTo(2, y))

	chip8.Execute(SetAddressRegisterTo(700))
	chip8.Execute(DrawSprite(1, 2, 1))
	if chip8.Register(15) != 0 {
		t.Error("1) flag must not be set")
	}

	chip8.Execute(SetAddressRegisterTo(701))
	chip8.Execute(DrawSprite(1, 2, 1))
	if chip8.Register(15) != 0 {
		t.Error("2) flag must not be set")
	}

	chip8.Execute(SetAddressRegisterTo(702))
	chip8.Execute(DrawSprite(1, 2, 1))
	if chip8.Register(15) != 1 {
		t.Error("3) flag should be set")
	}

	chip8.Execute(SetAddressRegisterTo(703))
	chip8.Execute(DrawSprite(1, 2, 1))
	if chip8.Register(15) != 0 {
		t.Error("4) flag must not be set")
	}
}

func TestDrawingSpriteWrapsAroundLeftAndRightBounds(t *testing.T) {
	setUp()
	chip8.memory[700] = 0xFF // 11111111
	chip8.memory[701] = 0xFF // 11111111
	var x, y uint16 = 61, 2
	chip8.Execute(SetRegisterTo(1, x))
	chip8.Execute(SetRegisterTo(2, y))

	chip8.Execute(SetAddressRegisterTo(700))
	chip8.Execute(DrawSprite(1, 2, 2))

	set := []point{
		{61, 2}, {62, 2}, {63, 2}, {0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
		{61, 3}, {62, 3}, {63, 3}, {0, 3}, {1, 3}, {2, 3}, {3, 3}, {4, 3},
	}
	checkPixelsAreSet(t, set)
	checkOtherPixelsAreNotSet(t, set)
}

func TestDrawingWrapsAroundMoreThanOnce(t *testing.T) {
	setUp()
	chip8.memory[700] = 0x94 // 10010100
	chip8.memory[701] = 0x94 // 10010100
	var x, y uint16 = 61 + 64, 30
	chip8.Execute(SetRegisterTo(1, x))
	chip8.Execute(SetRegisterTo(2, y))

	chip8.Execute(SetAddressRegisterTo(700))
	chip8.Execute(DrawSprite(1, 2, 2))

	set := []point{
		{61, 30}, {0, 30}, {2, 30},
		{61, 31}, {0, 31}, {2, 31},
	}
	checkPixelsAreSet(t, set)
	checkOtherPixelsAreNotSet(t, set)
}

func TestDrawingBelowLastLineIsSkipped(t *testing.T) {
	setUp()
	chip8.memory[700] = 0x80 // 10000000
	chip8.memory[701] = 0x80 // 10000000
	var x, y uint16 = 0, 31
	chip8.Execute(SetRegisterTo(1, x))
	chip8.Execute(SetRegisterTo(2, y))

	chip8.Execute(SetAddressRegisterTo(700))
	chip8.Execute(DrawSprite(1, 2, 2))

	set := []point{{0, 31}}
	checkPixelsAreSet(t, set)
	checkOtherPixelsAreNotSet(t, set)
}

func TestAfterClearScreenNoPixelIsSet(t *testing.T) {
	setUp()
	chip8.memory[700] = 0xFF
	chip8.memory[701] = 0xFF
	chip8.Execute(SetRegisterTo(1, 0))
	chip8.Execute(SetRegisterTo(2, 0))
	chip8.Execute(SetAddressRegisterTo(700))
	chip8.Execute(DrawSprite(1, 2, 2))

	chip8.Execute(ClearScreen)
	checkOtherPixelsAreNotSet(t, []point{})
}

func TestLoadingProgramCopiesItToHex200(t *testing.T) {
	setUp()
	chip8.Execute(CallRcp(0))
	chip8.LoadProgram([]uint8{1, 2, 3, 4, 5})
	for i := 0; i < 5; i++ {
		if chip8.MemoryByte(uint16(0x200+i)) != uint8(i+1) {
			t.Error("instruction", i, "not loaded")
		}
	}
	if chip8.ProgramCounter() != 0x200 {
		t.Error("pc not reset")
	}
}

func TestProgramCanBeLoadedIntoSpecifiedLocation(t *testing.T) {
	setUp()
	chip8.Execute(CallRcp(0))
	chip8.LoadProgramToAddress([]uint8{1, 2, 3, 4, 5}, 345)
	for i := 0; i < 5; i++ {
		if chip8.MemoryByte(uint16(345+i)) != uint8(i+1) {
			t.Error("instruction", i, "not loaded")
		}
	}
	if chip8.ProgramCounter() != 345 {
		t.Error("pc not reset")
	}
}

func TestInstructionAtProgramCounterCanBeExecuted(t *testing.T) {
	setUp()
	prog := toBytes(
		SetRegisterTo(2, 14),
		SetAddressRegisterTo(300),
		CopyToMemory(2))
	chip8.LoadProgramToAddress(prog, 500)

	chip8.ExecuteNext()
	chip8.ExecuteNext()
	chip8.ExecuteNext()

	if chip8.Register(2) != 14 {
		t.Error("first instruction not executed", chip8.Register(2))
	}
	if chip8.MemoryByte(302) != 14 {
		t.Error("third instruction not executed")
	}
	if chip8.ProgramCounter() != 506 {
		t.Error("wrong pc", chip8.ProgramCounter())
	}
}

func toBytes(instructions ...Instruction) []uint8 {
	bytes := make([]uint8, len(instructions)*2)
	for i, cmd := range instructions {
		bytes[i*2] = uint8((uint16(cmd) & 0xFF00) >> 8)
		bytes[i*2+1] = uint8(uint16(cmd) & 0x00FF)
	}
	return bytes
}

func TestAProgramIsAByteSequence(t *testing.T) {
	setUp()
	chip8.LoadProgram([]uint8{0x15, 0x43}) // jump to 0x0543
	chip8.ExecuteNext()
	if chip8.ProgramCounter() != 0x543 {
		t.Error("byte program not executed")
	}
}

func TestWaitingForKeyStopsExecution(t *testing.T) {
	setUp()
	stubEvent := &stubKeyEvent{}
	chip8.SetKeyEvent(stubEvent)
	prog := toBytes(
		WaitForKeyPress(5),
		SetAddressRegisterTo(123),
		SetAddressRegisterTo(111))
	chip8.LoadProgram(prog)

	stubEvent.hasKey = false
	if chip8.IsWaitingForKey() {
		t.Error("should not wait yet")
	}

	chip8.ExecuteNext()
	if !chip8.IsWaitingForKey() {
		t.Error("should be waiting for key")
	}

	pc := chip8.ProgramCounter()
	chip8.ExecuteNext()
	if chip8.ProgramCounter() != pc {
		t.Error("should not have executed while waiting")
	}

	stubEvent.hasKey = true
	stubEvent.key = 10
	chip8.ExecuteNext()
	if chip8.Register(5) != 10 {
		t.Error("key not stored")
	}
	if chip8.AddressRegister() != 123 {
		t.Error("second instruction not executed")
	}

	stubEvent.key = 3
	chip8.ExecuteNext()
	if chip8.Register(5) != 10 {
		t.Error("key register was overridden again")
	}
}

type stubKeyEvent struct {
	hasKey bool
	key    uint8
}

func (e *stubKeyEvent) LastKey() (hasOne bool, key uint8) {
	return e.hasKey, e.key
}

func (_ *stubKeyEvent) ClearLastKey() {}

func TestWaitingForKeyClearsTheLastKeyEventBuffer(t *testing.T) {
	setUp()
	spyEvent := &spyKeyEvent{}
	chip8.SetKeyEvent(spyEvent)
	prog := toBytes(WaitForKeyPress(5))
	chip8.LoadProgram(prog)
	chip8.ExecuteNext()
	if !spyEvent.cleared {
		t.Error("key event buffer not cleared")
	}
}

type spyKeyEvent struct {
	cleared bool
}

func (*spyKeyEvent) LastKey() (hasOne bool, key uint8) { return false, 0 }

func (e *spyKeyEvent) ClearLastKey() { e.cleared = true }

func TestBeepingRunsWhileSoundTimerIsAboveZero(t *testing.T) {
	setUp()
	spy := &spySpeakers{}
	chip8.SetSpeakers(spy)
	chip8.Execute(SetRegisterTo(0, 0))
	chip8.Execute(SetRegisterTo(2, 2))

	chip8.Execute(SetSoundTimer(0))
	if spy.beeping {
		t.Error("should not beep")
	}

	chip8.Execute(SetSoundTimer(2))
	if !spy.beeping {
		t.Error("should beep")
	}

	chip8.TimerTick()
	if !spy.beeping {
		t.Error("should still beep")
	}

	chip8.TimerTick()
	if spy.beeping {
		t.Error("should have stopped beeping")
	}
}

type spySpeakers struct {
	beeping bool
}

func (s *spySpeakers) Beep() {
	s.beeping = true
}

func (s *spySpeakers) BeQuiet() {
	s.beeping = false
}

func TestWhilePaused_CommandsAreNotExecuted(t *testing.T) {
	setUp()
	chip8.LoadProgram(toBytes(SetRegisterTo(1, 2)))

	chip8.SetPaused(true)
	if !chip8.Paused() {
		t.Error("not paused")
	}

	chip8.ExecuteNext()
	if chip8.Register(1) == 2 {
		t.Error("executed command while paused")
	}

	chip8.SetPaused(false)
	if chip8.Paused() {
		t.Error("paused")
	}

	chip8.ExecuteNext()
	if chip8.Register(1) != 2 {
		t.Error("command not executed after resume")
	}
}

func TestWhilePaused_TimersAreNotDecremented(t *testing.T) {
	setUp()
	chip8.Execute(SetRegisterTo(3, 3))
	chip8.Execute(SetDelayTimer(3))
	chip8.Execute(SetSoundTimer(3))

	chip8.SetPaused(true)
	chip8.TimerTick()

	if chip8.DelayTimer() != 3 {
		t.Error("delay timer ticked while paused")
	}
	if chip8.SoundTimer() != 3 {
		t.Error("sound timer ticked while paused")
	}

	chip8.SetPaused(false)
	chip8.TimerTick()

	if chip8.DelayTimer() != 2 {
		t.Error("delay timer not ticked after resume")
	}
	if chip8.SoundTimer() != 2 {
		t.Error("sound timer not ticked after resume")
	}
}

func TestPausingPausesSoundToo(t *testing.T) {
	setUp()
	spy := &spySpeakers{}
	chip8.SetSpeakers(spy)
	chip8.Execute(SetRegisterTo(5, 5))
	chip8.Execute(SetSoundTimer(5))
	chip8.SetPaused(true)
	if spy.beeping {
		t.Error("should not beep")
	}
	chip8.SetPaused(false)
	if !spy.beeping {
		t.Error("should beep again")
	}
}

func TestOnResume_KeyEventsFromPauseAreDismissed(t *testing.T) {
	setUp()
	spy := &spyKeyEvent{}
	chip8.SetKeyEvent(spy)
	chip8.SetPaused(true)
	chip8.SetPaused(false)
	if !spy.cleared {
		t.Error("should have cleared key buffer after resume")
	}
}
