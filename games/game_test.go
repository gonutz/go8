package games

import (
	"testing"
	"time"
)

func TestEmptyGameCanBeFilledWithDefaultData(t *testing.T) {
	g := &Game{}
	g.FillEmptyData()
	if g.Name != "Unknown" {
		t.Error("name", g.Name)
	}
	if g.Description != "No description available" {
		t.Error("description", g.Description)
	}
	if g.HowToPlay != "Find out for yourself\nhow to play this game..." {
		t.Error("how to:", g.HowToPlay)
	}
	if g.StartAddress != 0x200 {
		t.Error("start", g.StartAddress)
	}
	if g.ClockSpeed != 1*time.Millisecond {
		t.Error("clock speed", g.ClockSpeed)
	}
	white := Color{255, 255, 255}
	if g.ForegroundColor != white {
		t.Error("foreground color", g.ForegroundColor)
	}
	black := Color{0, 0, 0}
	if g.BackgroundColor != black {
		t.Error("background color", g.BackgroundColor)
	}
	checkDefaultKeyMap(t, g)
}

func checkDefaultKeyMap(t *testing.T, g *Game) {
	keys := map[GameKey]uint8{
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
	if len(g.Keys) != len(keys) {
		t.Error("keys: wrong lengths", g.Keys)
	}
	for key, value := range keys {
		if g.Keys[key] != value {
			t.Error("wrong key map", g.Keys[key], value)
		}
	}
}

func TestIfForeEqualsBackgroundColorTheyAreSetToBlackAndWhite(t *testing.T) {
	g := &Game{}
	g.ForegroundColor = Color{3, 3, 3}
	g.BackgroundColor = g.ForegroundColor

	g.FillEmptyData()

	white := Color{255, 255, 255}
	if g.ForegroundColor != white {
		t.Error("foreground color", g.ForegroundColor)
	}
	black := Color{0, 0, 0}
	if g.BackgroundColor != black {
		t.Error("background color", g.BackgroundColor)
	}
}

func TestAfterFillingInEmptyData_GameReturnsItself(t *testing.T) {
	g := &Game{}
	g2 := g.FillEmptyData()
	if g != g2 {
		t.Error("not the same")
	}
}

func TestAfterFillingInEmptyData_ScreenShotIsBlack(t *testing.T) {
	g := &Game{}
	g.FillEmptyData()

	for i := 0; i < 32; i++ {
		if g.ScreenShot[i] != 0 {
			t.Error("screen shot row ", i, g.ScreenShot[i])
		}
	}
}
