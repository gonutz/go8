package games

import "time"

var Blitz = &Game{
	Name:        "Blitz",
	Description: "Destroy Towers With Airbombs",
	HowToPlay: //
	`Drop bombs with the Space key.
	Destroy all towers before you
	crash into one of them as your
	altitude decreases.`,
	ClockSpeed:           1 * time.Millisecond,
	InstructionsPerCycle: 10,
	BackgroundColor:      Color{120, 200, 255},
	ForegroundColor:      Color{115, 50, 0},
	Keys: map[GameKey]uint8{
		Left:    6,
		Right:   6,
		Enter:   6,
		Space:   5,
		Down:    5,
		Number5: 5,
	},
	ScreenShot: [32]uint64{
		0, 524288, 1015808, 1032192, 0, 0, 0, 12582912, 0, 0, 0, 0, 0, 0, 0,
		12582924, 12582924, 12582924, 12582924, 12582924, 12582924,
		864691180007325708, 864691180007325708, 864691180007325708,
		864691180007325708, 864691386165755916, 864691386165755916,
		864691386165755916, 864691386165755916, 864691386165755916,
		864691386165755916, 864691386165755916,
	},
	Program: []byte{
		0x12, 0x17, 0x42, 0x4c, 0x49, 0x54, 0x5a, 0x20,
		0x42, 0x79, 0x20, 0x44, 0x61, 0x76, 0x69, 0x64,
		0x20, 0x57, 0x49, 0x4e, 0x54, 0x45, 0x52, 0xa3,
		0x41, 0x60, 0x04, 0x61, 0x09, 0x62, 0x0e, 0x67,
		0x04, 0xd0, 0x1e, 0xf2, 0x1e, 0x70, 0x0c, 0x30,
		0x40, 0x12, 0x21, 0xf0, 0x0a, 0x00, 0xe0, 0x22,
		0xd9, 0xf0, 0x0a, 0x00, 0xe0, 0x8e, 0x70, 0xa3,
		0x1e, 0x6b, 0x1f, 0xcc, 0x1f, 0x8c, 0xc4, 0xdc,
		0xb2, 0x3f, 0x01, 0x12, 0x49, 0xdc, 0xb2, 0x12,
		0x39, 0xca, 0x07, 0x7a, 0x01, 0x7b, 0xfe, 0xdc,
		0xb2, 0x7a, 0xff, 0x3a, 0x00, 0x12, 0x4d, 0x7e,
		0xff, 0x3e, 0x00, 0x12, 0x39, 0x6b, 0x00, 0x8c,
		0x70, 0x6d, 0x00, 0x6e, 0x00, 0xa3, 0x1b, 0xdd,
		0xe3, 0x3f, 0x00, 0x12, 0xc1, 0x3b, 0x00, 0x12,
		0x81, 0x60, 0x05, 0xe0, 0x9e, 0x12, 0x87, 0x6b,
		0x01, 0x88, 0xd0, 0x78, 0x02, 0x89, 0xe0, 0x79,
		0x03, 0xa3, 0x1e, 0xd8, 0x91, 0x81, 0xf0, 0x60,
		0x05, 0xf0, 0x15, 0xf0, 0x07, 0x30, 0x00, 0x12,
		0x8b, 0x3b, 0x01, 0x12, 0xab, 0xa3, 0x1e, 0x31,
		0x01, 0xd8, 0x91, 0x79, 0x01, 0x39, 0x20, 0x12,
		0xab, 0x6b, 0x00, 0x31, 0x00, 0x7c, 0xff, 0x4c,
		0x00, 0x12, 0xbb, 0xa3, 0x1b, 0xdd, 0xe3, 0x7d,
		0x02, 0x3d, 0x40, 0x12, 0xb9, 0x6d, 0x00, 0x7e,
		0x01, 0x12, 0x65, 0x00, 0xe0, 0x77, 0x02, 0x12,
		0x2d, 0xa3, 0x1b, 0xdd, 0xe3, 0x60, 0x14, 0x61,
		0x02, 0x62, 0x0b, 0xa3, 0x20, 0xd0, 0x1b, 0xf2,
		0x1e, 0x70, 0x08, 0x30, 0x2c, 0x12, 0xcd, 0x12,
		0xd7, 0x60, 0x0a, 0x61, 0x0d, 0x62, 0x05, 0xa3,
		0x07, 0xd0, 0x15, 0xf2, 0x1e, 0x70, 0x08, 0x30,
		0x2a, 0x12, 0xe1, 0x80, 0x70, 0x70, 0xfe, 0x80,
		0x06, 0xa3, 0x87, 0xf0, 0x33, 0xf2, 0x65, 0x60,
		0x2d, 0xf1, 0x29, 0x61, 0x0d, 0xd0, 0x15, 0x70,
		0x05, 0xf2, 0x29, 0xd0, 0x15, 0x00, 0xee, 0x83,
		0x82, 0x83, 0x82, 0xfb, 0xe8, 0x08, 0x88, 0x05,
		0xe2, 0xbe, 0xa0, 0xb8, 0x20, 0x3e, 0x80, 0x80,
		0x80, 0x80, 0xf8, 0x80, 0xf8, 0xfc, 0xc0, 0xc0,
		0xf9, 0x81, 0xdb, 0xcb, 0xfb, 0x00, 0xfa, 0x8a,
		0x9a, 0x99, 0xf8, 0xef, 0x2a, 0xe8, 0x29, 0x29,
		0x00, 0x6f, 0x68, 0x2e, 0x4c, 0x8f, 0xbe, 0xa0,
		0xb8, 0xb0, 0xbe, 0x00, 0xbe, 0x22, 0x3e, 0x34,
		0xb2, 0xd8, 0xd8, 0x00, 0xc3, 0xc3, 0x00, 0xd8,
		0xd8, 0x00, 0xc3, 0xc3, 0x00, 0xd8, 0xd8, 0xc0,
		0xc0, 0x00, 0xc0, 0xc0, 0x00, 0xc0, 0xc0, 0x00,
		0xc0, 0xc0, 0x00, 0xdb, 0xdb, 0xdb, 0xdb, 0x00,
		0x18, 0x18, 0x00, 0x18, 0x18, 0x00, 0x18, 0x18,
		0x00, 0xdb, 0xdb, 0xdb, 0xdb, 0x00, 0x18, 0x18,
		0x00, 0x18, 0x18, 0x00, 0x18, 0x18, 0x00, 0x18,
		0x18, 0xdb, 0xdb, 0x00, 0x03, 0x03, 0x00, 0x18,
		0x18, 0x00, 0xc0, 0xc0, 0x00, 0xdb, 0xdb,
	},
}
