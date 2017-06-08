package games

import "time"

type Game struct {
	Name                 string
	Description          string
	HowToPlay            string
	ClockSpeed           time.Duration
	InstructionsPerCycle int
	ForegroundColor      Color
	BackgroundColor      Color
	StartAddress         uint16
	Keys                 map[GameKey]uint8
	Program              []byte
	ScreenShot           [32]uint64
}

type Color struct{ R, G, B uint8 }

type GameKey int

const (
	Left GameKey = iota
	Right
	Up
	Down
	Enter
	Space
	Number0
	Number1
	Number2
	Number3
	Number4
	Number5
	Number6
	Number7
	Number8
	Number9
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyQ

	NumberOfGameKeys // this has to be last in the enumeration
)

func (g *Game) FillEmptyData() *Game {
	if g.Name == "" {
		g.Name = "Unknown"
	}
	if g.Description == "" {
		g.Description = "No description available"
	}
	if g.HowToPlay == "" {
		g.HowToPlay = "Find out for yourself\nhow to play this game..."
	}
	if g.StartAddress == 0 {
		g.StartAddress = 0x200
	}
	if g.ClockSpeed == 0 {
		g.ClockSpeed = 1 * time.Millisecond
	}
	if g.InstructionsPerCycle == 0 {
		g.InstructionsPerCycle = 1
	}
	if g.BackgroundColor == g.ForegroundColor {
		g.BackgroundColor = Color{0, 0, 0}
		g.ForegroundColor = Color{255, 255, 255}
	}
	if g.Keys == nil {
		fillKeys(g)
	}
	return g
}

func fillKeys(g *Game) {
	g.Keys = map[GameKey]uint8{
		Number0: 0,
		Number1: 1,
		Number2: 2,
		Number3: 3,
		Number4: 4,
		Number5: 5,
		Number6: 6,
		Number7: 7,
		Number8: 8,
		Number9: 9,
		KeyA:    10,
		KeyB:    11,
		KeyC:    12,
		KeyD:    13,
		KeyE:    14,
		KeyF:    15,
	}
}
