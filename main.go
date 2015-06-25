package main

import (
	"github.com/gonutz/go8/chip8"
	"github.com/gonutz/go8/games"
	"github.com/gonutz/prototype/draw"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	vm := chip8.NewInterpreter()
	vm.SetRandomizer(randomizer{})

	speakers := &speakers{}
	vm.SetSpeakers(speakers)

	keyNames := []string{
		"0", "1", "2", "3",
		"4", "5", "6", "7",
		"8", "9", "a", "s",
		"d", "f", "g", "h",
	}
	keyboard := &keyboard{
		window:   nil,
		keyNames: keyNames,
	}
	vm.SetKeyboard(keyboard)

	keyEvent := &keyEvent{
		keyNames: keyNames,
		last:     -1,
	}
	vm.SetKeyEvent(keyEvent)

	vm.LoadProgram(games.Blitz.Program)

	draw.RunWindow("GO 8", 640, 320, func(window *draw.Window) {
		if window.WasKeyPressed("escape") {
			window.Close()
			return
		}

		speakers.window = window
		keyboard.window = window
		keyEvent.update(window)

		const instructionsPerFrame = 100
		for i := 0; i < instructionsPerFrame; i++ {
			vm.ExecuteNext()
		}
		vm.TimerTick()

		drawScreen(window, vm.Screen())
	})
}

func drawScreen(window *draw.Window, screen chip8.Screen) {
	window.FillRect(0, 0, 640, 480, draw.Black)
	w, h := screen.Size()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if screen.IsSet(x, y) {
				window.FillRect(x*10, y*10, 10, 10, draw.White)
			}
		}
	}
}

type randomizer struct{}

func (randomizer) Next() uint8 {
	return uint8(rand.Int() % 256)
}

type speakers struct {
	window *draw.Window
}

func (s *speakers) Beep() {
	s.window.PlaySoundFile("./beep.wav")
}

func (*speakers) BeQuiet() {
}

type keyboard struct {
	window   *draw.Window
	keyNames []string
}

func (k *keyboard) KeyDown(key uint8) bool {
	if int(key) >= len(k.keyNames) {
		return false
	}
	return k.window.IsKeyDown(k.keyNames[key])
}

type keyEvent struct {
	keyNames []string
	last     int
}

func (k *keyEvent) LastKey() (keyWasPressed bool, key uint8) {
	return k.last != -1, uint8(k.last)
}

func (k *keyEvent) ClearLastKey() {
	k.last = -1
}

func (k *keyEvent) update(window *draw.Window) {
	for i, name := range k.keyNames {
		if window.WasKeyPressed(name) {
			k.last = i
		}
	}
}
