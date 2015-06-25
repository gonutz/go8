package games

import "time"

var AdditionProblems = &Game{
	Name:        "Addition Problems",
	Description: "Add the Numbers Quickly",
	HowToPlay: //
	`Add two numbers and enter the
	result using three digits.
	Example: to enter 12, type 012.
	If you are right, you get a C else
	you get E and the right answer.`,
	ClockSpeed:      1 * time.Millisecond,
	ForegroundColor: Color{0, 0, 0},
	BackgroundColor: Color{255, 255, 255},
	Keys: map[GameKey]uint8{
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
		Enter:   14,
		Space:   14,
	},
	ScreenShot: [32]uint64{
		17851160700068560896, 10422473245855825920, 10933750413240827904,
		10674674687415402496, 17851160872035024896, 0, 0, 0, 1365325643776,
		3440839229440, 1365325643776, 1117261922304, 4114104713216, 0, 0, 0,
		1365262729216, 3320546590720, 1365262729216, 1237487452160,
		4114041798656, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	},
	Program: []byte{
		0x00, 0xe0, 0xcd, 0x7f, 0xce, 0x7f, 0x8c, 0xd0,
		0x8c, 0xe4, 0xa2, 0xa2, 0x6a, 0x00, 0x6b, 0x00,
		0xfd, 0x33, 0xf2, 0x65, 0x22, 0x76, 0xa2, 0x88,
		0x7a, 0x07, 0xda, 0xb5, 0xa2, 0xa2, 0x7a, 0x08,
		0xfe, 0x33, 0xf2, 0x65, 0x22, 0x76, 0xa2, 0x8e,
		0x7a, 0x07, 0xda, 0xb4, 0xa2, 0x92, 0x6a, 0x18,
		0x6b, 0x08, 0xda, 0xbf, 0xf0, 0x0a, 0xf1, 0x0a,
		0xf2, 0x0a, 0xda, 0xbf, 0x6a, 0x15, 0x22, 0x76,
		0xa2, 0xa5, 0xf2, 0x55, 0xa2, 0xa2, 0xfc, 0x33,
		0xf5, 0x65, 0x83, 0x05, 0x33, 0x00, 0x12, 0x62,
		0x84, 0x15, 0x34, 0x00, 0x12, 0x62, 0x85, 0x25,
		0x35, 0x00, 0x12, 0x62, 0x66, 0x0c, 0xf6, 0x18,
		0x12, 0x6a, 0x6a, 0x15, 0x6b, 0x10, 0x22, 0x76,
		0x66, 0x0e, 0x6a, 0x26, 0x6b, 0x08, 0xf6, 0x29,
		0xda, 0xb5, 0xf0, 0x0a, 0x12, 0x00, 0xf0, 0x29,
		0xda, 0xb5, 0x7a, 0x05, 0xf1, 0x29, 0xda, 0xb5,
		0x7a, 0x05, 0xf2, 0x29, 0xda, 0xb5, 0x00, 0xee,
		0x20, 0x20, 0xf8, 0x20, 0x20, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff, 0x03, 0x03, 0x03, 0xff,
		0xff, 0xc0, 0xc0, 0xc0, 0xc0, 0xc0, 0x00, 0xc0,
		0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	},
}
