package chip8

type Interpreter struct {
	pc              uint16
	registers       [16]uint8
	addressRegister uint16
	stack           []uint16
	memory          [4096]uint8
	rand            Randomizer
	keyboard        Keyboard
	event           KeyEvent
	speakers        Speakers
	screen          *IntScreen
	waitingForKey   bool
	waitRegister    int
	delay           uint8
	sound           uint8
	paused          bool
}

type Randomizer interface {
	Next() uint8
}

type Keyboard interface {
	KeyDown(which uint8) bool
}

type KeyEvent interface {
	LastKey() (keyWasPressed bool, key uint8)
	ClearLastKey()
}

type Speakers interface {
	Beep()
	BeQuiet()
}

type Screen interface {
	IsSet(x, y int) bool
	Size() (w, h int)
}

func NewInterpreter() *Interpreter {
	i := &Interpreter{pc: 0x200, screen: &IntScreen{}}
	i.initMemory()
	return i
}

func (chip8 *Interpreter) initMemory() {
	digits := []uint8{
		0xF0, 0x90, 0x90, 0x90, 0xF0,
		0x20, 0x60, 0x20, 0x20, 0x70,
		0xF0, 0x10, 0xF0, 0x80, 0xF0,
		0xF0, 0x10, 0xF0, 0x10, 0xF0,
		0x90, 0x90, 0xF0, 0x10, 0x10,
		0xF0, 0x80, 0xF0, 0x10, 0xF0,
		0xF0, 0x80, 0xF0, 0x90, 0xF0,
		0xF0, 0x10, 0x20, 0x40, 0x40,
		0xF0, 0x90, 0xF0, 0x90, 0xF0,
		0xF0, 0x90, 0xF0, 0x10, 0xF0,
		0xF0, 0x90, 0xF0, 0x90, 0x90,
		0xE0, 0x90, 0xE0, 0x90, 0xE0,
		0xF0, 0x80, 0x80, 0x80, 0xF0,
		0xE0, 0x90, 0x90, 0x90, 0xE0,
		0xF0, 0x80, 0xF0, 0x80, 0xF0,
		0xF0, 0x80, 0xF0, 0x80, 0x80,
	}
	for i := range digits {
		chip8.memory[i] = digits[i]
	}
}

func (i *Interpreter) SetRandomizer(rand Randomizer) {
	i.rand = rand
}

func (i *Interpreter) SetKeyboard(keyboard Keyboard) {
	i.keyboard = keyboard
}

func (i *Interpreter) SetKeyEvent(event KeyEvent) {
	i.event = event
}

func (i *Interpreter) SetSpeakers(s Speakers) {
	i.speakers = s
}

func (i *Interpreter) TimerTick() {
	if !i.paused {
		if i.delay != 0 {
			i.delay--
		}
		if i.sound > 0 {
			i.sound--
		}
		if i.sound == 0 {
			i.speakers.BeQuiet()
		}
	}
}

func (i *Interpreter) Screen() Screen {
	return i.screen
}

func (i *Interpreter) IsWaitingForKey() bool {
	return i.waitingForKey
}

func (i *Interpreter) ProgramCounter() uint16 {
	return i.pc
}

func (i *Interpreter) MemoryByte(index uint16) uint8 {
	return i.memory[index]
}

func (i *Interpreter) DelayTimer() uint8 {
	return i.delay
}

func (i *Interpreter) SoundTimer() uint8 {
	return i.sound
}

func (i *Interpreter) Register(which int) uint8 {
	return i.registers[which]
}

func (i *Interpreter) AddressRegister() uint16 {
	return i.addressRegister
}

func (chip8 *Interpreter) LoadProgram(prog []byte) {
	chip8.LoadProgramToAddress(prog, 0x200)
}

func (chip8 *Interpreter) LoadProgramToAddress(prog []byte, address uint16) {
	for index, b := range prog {
		chip8.memory[int(address)+index] = b
	}
	chip8.pc = address
}

func (i *Interpreter) ExecuteNext() {
	if !i.paused {
		if i.waitingForKey {
			hasKey, key := i.event.LastKey()
			if !hasKey {
				return
			}
			i.waitingForKey = false
			i.registers[i.waitRegister] = key
		}
		cmd := makeInstruction(i.memory[i.pc], i.memory[i.pc+1])
		i.Execute(cmd)
	}
}

func makeInstruction(high, low uint8) Instruction {
	return Instruction(uint16(high)<<8 + uint16(low))
}

func (i *Interpreter) Execute(command Instruction) {
	nibble3 := int(command & 0xF000 >> 12)
	nibble2 := int(command & 0x0F00 >> 8)
	nibble1 := int(command & 0x00F0 >> 4)
	nibble0 := int(command & 0x000F)
	byte0 := uint8(command & 0x00FF)
	address := uint16(command & 0x0FFF)

	i.pc += 2

	switch nibble3 {
	case 0:
		switch command {
		case Return:
			i.returnFromCurrentFunction()
		case ClearScreen:
			i.clearScreen()
		}
	case 1:
		i.jumpTo(address)
	case 2:
		i.callFunctionAt(address)
	case 3:
		i.skipNextInstructionIfEqual(i.Register(nibble2), byte0)
	case 4:
		i.skipNextInstructionIfNotEqual(i.Register(nibble2), byte0)
	case 5:
		i.skipNextInstructionIfRegistersEqual(nibble2, nibble1)
	case 6:
		i.registers[nibble2] = byte0
	case 7:
		i.registers[nibble2] += byte0
	case 8:
		i.executeArithmeticInstruction(nibble0, nibble2, nibble1)
	case 9:
		i.skipNextInstructionIfRegistersNotEqual(nibble2, nibble1)
	case 10:
		i.addressRegister = address
	case 11:
		i.jumpTo(address + uint16(i.registers[0]))
	case 12:
		i.registers[nibble2] = i.rand.Next() & byte0
	case 13:
		i.drawSprite(nibble2, nibble1, int(nibble0))
	case 14:
		if byte0 == 0x9E && i.keyboard.KeyDown(i.Register(nibble2)) {
			i.pc += 2
		}
		if byte0 == 0xA1 && !i.keyboard.KeyDown(i.Register(nibble2)) {
			i.pc += 2
		}
	case 15:
		switch byte0 {
		case 0x7:
			i.registers[nibble2] = i.delay
		case 0xA:
			i.waitingForKey = true
			i.waitRegister = int(nibble2)
			i.event.ClearLastKey()
		case 0x15:
			i.delay = i.registers[nibble2]
		case 0x18:
			i.setSoundTimerTo(i.registers[nibble2])
		case 0x1E:
			i.addressRegister += uint16(i.registers[nibble2])
		case 0x29:
			i.addressRegister = startAddressOfDigit(i.registers[nibble2])
		case 0x33:
			value := int(i.registers[nibble2])
			ones := value % 10
			tens := (value / 10) % 10
			hundreds := value / 100
			addr := i.addressRegister
			i.memory[addr] = uint8(hundreds)
			i.memory[addr+1] = uint8(tens)
			i.memory[addr+2] = uint8(ones)
		case 0x55:
			addr := i.addressRegister
			for x := 0; x <= nibble2; x++ {
				i.memory[addr+uint16(x)] = i.registers[x]
			}
		case 0x65:
			addr := i.addressRegister
			for x := 0; x <= nibble2; x++ {
				i.registers[x] = i.memory[addr+uint16(x)]
			}
		}
	}
}

func (i *Interpreter) returnFromCurrentFunction() {
	i.pc = i.stack[0]
	i.stack = i.stack[1:]
}

func (i *Interpreter) clearScreen() {
	for index := range i.screen.Rows {
		i.screen.Rows[index] = 0
	}
}

func (i *Interpreter) jumpTo(address uint16) {
	i.pc = address
}

func (i *Interpreter) callFunctionAt(address uint16) {
	i.stack = append([]uint16{i.pc}, i.stack...)
	i.pc = address
}

func (i *Interpreter) skipNextInstructionIfEqual(a, b uint8) {
	if a == b {
		i.pc += 2
	}
}

func (i *Interpreter) skipNextInstructionIfNotEqual(a, b uint8) {
	if a != b {
		i.pc += 2
	}
}

func (i *Interpreter) skipNextInstructionIfRegistersEqual(r1, r2 int) {
	if i.Register(r1) == i.Register(r2) {
		i.pc += 2
	}
}

func (i *Interpreter) skipNextInstructionIfRegistersNotEqual(r1, r2 int) {
	if i.Register(r1) != i.Register(r2) {
		i.pc += 2
	}
}

func (i *Interpreter) executeArithmeticInstruction(which, dest, source int) {
	switch which {
	case 0:
		i.registers[dest] = i.registers[source]
	case 1:
		i.registers[dest] |= i.registers[source]
	case 2:
		i.registers[dest] &= i.registers[source]
	case 3:
		i.registers[dest] ^= i.registers[source]
	case 4:
		i.add(dest, source)
	case 5:
		i.subtract(dest, dest, source)
	case 6:
		i.registers[15] = i.registers[dest] & 1
		i.registers[dest] /= 2
	case 7:
		i.subtract(dest, source, dest)
	case 14:
		i.registers[15] = i.registers[dest] & 0x80 >> 7
		i.registers[dest] *= 2
	}
}

func (i *Interpreter) add(dest, source int) {
	i.registers[15] = 0
	if int(i.registers[dest])+int(i.registers[source]) > 255 {
		i.registers[0xF] = 1
	}
	i.registers[dest] += i.registers[source]
}

func (i *Interpreter) subtract(into, rLeft, rRight int) {
	i.registers[15] = 0
	if i.registers[rLeft] >= i.registers[rRight] {
		i.registers[15] = 1
	}
	i.registers[into] = i.registers[rLeft] - i.registers[rRight]
}

func (i *Interpreter) drawSprite(xRegister, yRegister, height int) {
	x := int(i.registers[xRegister]) % 64
	y := int(i.registers[yRegister])
	addr := int(i.addressRegister)
	i.registers[15] = 0
	for index := 0; index < height; index++ {
		line := uint64(i.memory[addr+index])
		mask := lineMask(x, line)
		if y+index < 32 {
			if i.screen.Rows[y+index]&mask != 0 {
				i.registers[15] = 1
			}
			i.screen.Rows[y+index] ^= mask
		}
	}
}

func lineMask(x int, line uint64) uint64 {
	if x <= 64-8 {
		return line << uint(64-8-x)
	} else {
		return line>>uint(x-(64-8)) + line<<uint(64-x+(64-8))
	}
}

func (i *Interpreter) setSoundTimerTo(value uint8) {
	i.sound = value
	if i.sound > 0 {
		i.speakers.Beep()
	}
}

func startAddressOfDigit(digit uint8) uint16 {
	return 5 * uint16(digit)
}

func (chip8 *Interpreter) SetPaused(paused bool) {
	chip8.paused = paused
	chip8.pauseOrUnpauseSound(paused)
	chip8.event.ClearLastKey()
}

func (chip8 *Interpreter) pauseOrUnpauseSound(paused bool) {
	if paused {
		chip8.speakers.BeQuiet()
	} else if chip8.sound > 0 {
		chip8.speakers.Beep()
	}
}

func (chip8 *Interpreter) Paused() bool {
	return chip8.paused
}
