package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/gonutz/go8/chip8"
	"github.com/gonutz/go8/games"
	"github.com/gonutz/prototype/draw"
)

func init() {
	runtime.LockOSThread()
}

var game *games.Game

func main() {
	game = games.Blitz
	game.FillEmptyData()
	rand.Seed(time.Now().UnixNano())

	vm := chip8.NewInterpreter()
	vm.SetRandomizer(randomizer{})

	speakers := &speakers{}
	vm.SetSpeakers(speakers)

	keyboard := &keyboard{}
	vm.SetKeyboard(keyboard)

	keyEvent := &keyEvent{
		last: -1,
	}
	vm.SetKeyEvent(keyEvent)

	vm.LoadProgram(game.Program)

	clockTicksPerFrame := int(float64(game.ClockSpeed)/float64(time.Second/60) + 0.5)
	if clockTicksPerFrame < 1 {
		clockTicksPerFrame = 1
	}
	instructionsPerFrame := game.InstructionsPerCycle * clockTicksPerFrame
	draw.RunWindow("GO 8", 640, 320, func(window draw.Window) {
		if window.WasKeyPressed(draw.KeyEscape) {
			window.Close()
			return
		}
		if window.WasKeyPressed(draw.KeyP) {
			vm.SetPaused(!vm.Paused())
		}

		speakers.window = window
		keyboard.window = window
		keyEvent.update(window)

		for i := 0; i < instructionsPerFrame; i++ {
			vm.ExecuteNext()
		}
		vm.TimerTick()

		drawScreen(window, vm.Screen())
	})
}

func drawScreen(window draw.Window, screen chip8.Screen) {
	b := game.BackgroundColor
	back := draw.Color{float32(b.R) / 255, float32(b.G) / 255, float32(b.B) / 255, 1}
	f := game.ForegroundColor
	fore := draw.Color{float32(f.R) / 255, float32(f.G) / 255, float32(f.B) / 255, 1}

	windowW, windowH := window.Size()
	window.FillRect(0, 0, windowW, windowH, draw.Gray)
	screenW, screenH := screen.Size()
	tileSize := min(windowW/screenW, windowH/screenH)
	xOffset := (windowW - screenW*tileSize) / 2
	yOffset := (windowH - screenH*tileSize) / 2
	window.FillRect(xOffset, yOffset, screenW*tileSize, screenH*tileSize, back)
	for y := 0; y < screenH; y++ {
		for x := 0; x < screenW; x++ {
			if screen.IsSet(x, y) {
				window.FillRect(
					xOffset+x*tileSize,
					yOffset+y*tileSize,
					tileSize,
					tileSize,
					fore,
				)
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type randomizer struct{}

func (randomizer) Next() uint8 {
	return uint8(rand.Int() % 256)
}

type speakers struct {
	window draw.Window
}

func (s *speakers) Beep() {
	s.window.PlaySoundFile("./beep.wav")
}

func (*speakers) BeQuiet() {}

type keyboard struct {
	window draw.Window
}

func (k *keyboard) KeyDown(key uint8) bool {
	for gameKey, intKey := range game.Keys {
		for _, keyName := range gameKeyToKeyName[gameKey] {
			if key == intKey && k.window.IsKeyDown(keyName) {
				return true
			}
		}
	}
	return false
}

var gameKeyToKeyName = map[games.GameKey][]draw.Key{
	games.Left:    {draw.KeyLeft},
	games.Right:   {draw.KeyRight},
	games.Up:      {draw.KeyUp},
	games.Down:    {draw.KeyDown},
	games.Enter:   {draw.KeyEnter, draw.KeyNumEnter},
	games.Space:   {draw.KeySpace},
	games.Number0: {draw.Key0, draw.KeyNum0},
	games.Number1: {draw.Key1, draw.KeyNum1},
	games.Number2: {draw.Key2, draw.KeyNum2},
	games.Number3: {draw.Key3, draw.KeyNum3},
	games.Number4: {draw.Key4, draw.KeyNum4},
	games.Number5: {draw.Key5, draw.KeyNum5},
	games.Number6: {draw.Key6, draw.KeyNum6},
	games.Number7: {draw.Key7, draw.KeyNum7},
	games.Number8: {draw.Key8, draw.KeyNum8},
	games.Number9: {draw.Key9, draw.KeyNum9},
	games.KeyA:    {draw.KeyA},
	games.KeyB:    {draw.KeyB},
	games.KeyC:    {draw.KeyC},
	games.KeyD:    {draw.KeyD},
	games.KeyE:    {draw.KeyE},
	games.KeyF:    {draw.KeyF},
	games.KeyQ:    {draw.KeyQ},
}

type keyEvent struct {
	last int
}

func (k *keyEvent) LastKey() (keyWasPressed bool, key uint8) {
	return k.last != -1, uint8(k.last)
}

func (k *keyEvent) ClearLastKey() {
	k.last = -1
}

func (k *keyEvent) update(window draw.Window) {
	for gameKey, intKey := range game.Keys {
		for _, keyName := range gameKeyToKeyName[gameKey] {
			if window.WasKeyPressed(keyName) {
				k.last = int(intKey)
			}
		}
	}
}
