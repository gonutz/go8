package games

import "time"

var AstroDodge = &Game{
	Name:        "Astro Dodge",
	Description: "Dodge the Asteroids",
	HowToPlay: //
	`Start the game with Space and
	move with the arrow-keys.
	Dodge the asteroids to score.`,
	ClockSpeed:           5 * time.Millisecond,
	InstructionsPerCycle: 4,
	ForegroundColor:      Color{200, 100, 80},
	BackgroundColor:      Color{0, 0, 0},
	Keys: map[GameKey]uint8{
		Space: 5,
		Left:  4,
		Right: 6,
		Up:    2,
		Down:  8,
	},
	Program: []byte{
		0x12, 0x14, 0x52, 0x45, 0x56, 0x49, 0x56, 0x41,
		0x4c, 0x53, 0x54, 0x55, 0x44, 0x49, 0x4f, 0x53,
		0x32, 0x30, 0x30, 0x38, 0x00, 0xe0, 0x6d, 0x20,
		0xfd, 0x15, 0x24, 0x58, 0x24, 0x60, 0x6d, 0x40,
		0xfd, 0x15, 0x24, 0x58, 0x24, 0x60, 0x6d, 0x20,
		0xfd, 0x15, 0x24, 0x58, 0x00, 0xe0, 0xa3, 0xd8,
		0x25, 0x20, 0x24, 0xb8, 0x6d, 0x04, 0x6c, 0x00,
		0x60, 0x05, 0xe0, 0x9e, 0x12, 0x40, 0x12, 0x54,
		0xfd, 0x15, 0x24, 0x58, 0x7c, 0x01, 0x4c, 0x00,
		0x24, 0xb8, 0x4c, 0x04, 0x24, 0xb8, 0x4c, 0x08,
		0x6c, 0x00, 0x12, 0x38, 0xa3, 0xd8, 0x25, 0x20,
		0x00, 0xe0, 0x68, 0x10, 0x69, 0x14, 0x22, 0xe2,
		0x22, 0x88, 0x25, 0x4a, 0x63, 0x2c, 0x64, 0x00,
		0x25, 0x82, 0x63, 0x2c, 0x64, 0x06, 0x25, 0x8e,
		0x6d, 0x00, 0x6e, 0x08, 0x22, 0xb0, 0x22, 0xe8,
		0x7d, 0x08, 0x4d, 0x80, 0x6d, 0x00, 0x7e, 0x08,
		0x4e, 0x80, 0x6e, 0x00, 0x12, 0x74, 0x12, 0x2c,
		0x6c, 0x00, 0xa3, 0x3e, 0xfc, 0x1e, 0xf1, 0x65,
		0xc1, 0x03, 0xf1, 0x55, 0xa3, 0x58, 0xd0, 0x18,
		0x7c, 0x03, 0x3c, 0x09, 0x12, 0x8a, 0x00, 0xee,
		0x23, 0x1a, 0x25, 0x52, 0x23, 0x1a, 0xc0, 0x0f,
		0x80, 0x04, 0x80, 0x04, 0x61, 0x00, 0x00, 0xee,
		0x6c, 0x00, 0x65, 0x00, 0xa3, 0x3e, 0xfc, 0x1e,
		0xf2, 0x65, 0xa3, 0x58, 0xfd, 0x1e, 0xd0, 0x18,
		0x81, 0x24, 0x83, 0x10, 0x64, 0x1e, 0x83, 0x45,
		0x4f, 0x01, 0x22, 0xa0, 0xa3, 0x58, 0xfe, 0x1e,
		0xd0, 0x18, 0xa3, 0x3e, 0xfc, 0x1e, 0xf1, 0x55,
		0x75, 0x18, 0x7c, 0x03, 0x3c, 0x09, 0x12, 0xb4,
		0x00, 0xee, 0xa3, 0x50, 0xd8, 0x98, 0x00, 0xee,
		0x22, 0xe2, 0x60, 0x04, 0xe0, 0x9e, 0x12, 0xf4,
		0x38, 0x00, 0x78, 0xfe, 0x60, 0x06, 0xe0, 0x9e,
		0x12, 0xfe, 0x38, 0x38, 0x78, 0x02, 0x60, 0x02,
		0xe0, 0x9e, 0x13, 0x08, 0x39, 0x10, 0x79, 0xff,
		0x60, 0x08, 0xe0, 0x9e, 0x13, 0x12, 0x39, 0x18,
		0x79, 0x01, 0x22, 0xe2, 0x4f, 0x01, 0x13, 0x28,
		0x00, 0xee, 0x63, 0x2c, 0x64, 0x00, 0x25, 0x82,
		0x63, 0x2c, 0x64, 0x06, 0x25, 0x8e, 0x00, 0xee,
		0x00, 0xe0, 0x60, 0x00, 0x61, 0x04, 0x24, 0xee,
		0x63, 0x16, 0x64, 0x16, 0x25, 0x82, 0x60, 0x05,
		0xe0, 0x9e, 0x13, 0x36, 0x12, 0x2c, 0x00, 0x0e,
		0x01, 0x18, 0xb4, 0x02, 0x30, 0xe3, 0x03, 0x40,
		0x75, 0x01, 0x40, 0x60, 0x02, 0x40, 0x36, 0x03,
		0x18, 0x18, 0x34, 0x24, 0x7e, 0xff, 0xe7, 0x99,
		0x00, 0x40, 0x38, 0x14, 0x2a, 0x75, 0x3a, 0x14,
		0x00, 0x00, 0x28, 0x70, 0x3e, 0x07, 0x2a, 0x00,
		0x00, 0x40, 0x08, 0x5c, 0x7a, 0x75, 0x0a, 0x04,
		0x00, 0x50, 0x28, 0x54, 0x3e, 0x75, 0x2e, 0x01,
		0x20, 0x70, 0x78, 0x7c, 0x3c, 0x75, 0x6a, 0x54,
		0x00, 0x64, 0x78, 0x78, 0x7e, 0x7d, 0xe8, 0x50,
		0x08, 0x44, 0x2a, 0x50, 0x3a, 0x5d, 0xe8, 0x40,
		0x08, 0x54, 0x0a, 0x11, 0x2a, 0x5c, 0x68, 0x40,
		0x00, 0x04, 0x2a, 0x57, 0x2a, 0x16, 0x28, 0x50,
		0x00, 0x00, 0x2a, 0x13, 0x6a, 0x56, 0x08, 0x00,
		0x00, 0x00, 0x08, 0x71, 0x2a, 0x52, 0x30, 0x00,
		0x00, 0x04, 0x0a, 0x14, 0x3a, 0x60, 0x68, 0x00,
		0x00, 0x04, 0x0a, 0x1c, 0x3e, 0x70, 0x68, 0x50,
		0x20, 0x50, 0x2e, 0x5f, 0x2e, 0x5c, 0x28, 0x50,
		0x20, 0x5c, 0x3a, 0x57, 0x3e, 0x5e, 0x28, 0x50,
		0x00, 0x58, 0x38, 0x77, 0x2e, 0x7f, 0x3e, 0x54,
		0x78, 0x7e, 0xf7, 0xf1, 0xbe, 0x9e, 0x38, 0x66,
		0xf8, 0xcd, 0x73, 0x9b, 0xe3, 0x36, 0xcf, 0xde,
		0x78, 0x66, 0xf7, 0x1d, 0xbe, 0x9b, 0x3a, 0x36,
		0xf9, 0xcd, 0x73, 0xb3, 0xe3, 0x6c, 0xcf, 0xd8,
		0xcd, 0x33, 0x83, 0x78, 0x33, 0xd9, 0x66, 0x9c,
		0xcd, 0xf8, 0x9b, 0xe3, 0x36, 0x87, 0x6c, 0x3c,
		0xfd, 0x7e, 0xf7, 0x79, 0xbf, 0x9f, 0x7c, 0x66,
		0xfc, 0xcd, 0xfb, 0x9b, 0xf7, 0x36, 0xef, 0xde,
		0xfd, 0x66, 0xf7, 0x1d, 0xbf, 0x9b, 0x7c, 0x36,
		0xfc, 0xcd, 0xfb, 0xb3, 0xf7, 0xec, 0xef, 0xd8,
		0xcd, 0x00, 0xc3, 0x00, 0x33, 0x00, 0x66, 0x00,
		0xcd, 0x00, 0x9b, 0x00, 0x36, 0x00, 0x0c, 0x00,
		0x85, 0x66, 0x00, 0x19, 0x01, 0x9b, 0x46, 0x66,
		0x05, 0xcd, 0x88, 0x9b, 0x14, 0x6c, 0x20, 0xd8,
		0xcd, 0x33, 0x83, 0x7c, 0x33, 0xd9, 0x6e, 0xbe,
		0xcd, 0xfd, 0xdb, 0xf3, 0x76, 0xcf, 0x6c, 0xbc,
		0xf0, 0x07, 0x30, 0x00, 0x14, 0x58, 0x00, 0xee,
		0x6d, 0x04, 0x61, 0x0c, 0x60, 0x1c, 0x62, 0x12,
		0xa6, 0x2f, 0xf2, 0x1e, 0xd0, 0x16, 0xfd, 0x15,
		0x24, 0x58, 0x60, 0x14, 0x62, 0x0c, 0xa6, 0x2f,
		0xf2, 0x1e, 0xd0, 0x16, 0x60, 0x24, 0x62, 0x18,
		0xa6, 0x2f, 0xf2, 0x1e, 0xd0, 0x16, 0xfd, 0x15,
		0x24, 0x58, 0x60, 0x0c, 0x62, 0x06, 0xa6, 0x2f,
		0xf2, 0x1e, 0xd0, 0x16, 0x60, 0x2c, 0x62, 0x1e,
		0xa6, 0x2f, 0xf2, 0x1e, 0xd0, 0x16, 0xfd, 0x15,
		0x24, 0x58, 0xa6, 0x2f, 0x60, 0x04, 0xd0, 0x16,
		0x60, 0x34, 0x62, 0x24, 0xa6, 0x2f, 0xf2, 0x1e,
		0xd0, 0x16, 0xfd, 0x15, 0x24, 0x58, 0x00, 0xee,
		0x62, 0x06, 0x60, 0x00, 0x61, 0x17, 0xa5, 0x9f,
		0xd0, 0x16, 0x70, 0x08, 0xf2, 0x1e, 0xd0, 0x16,
		0x70, 0x08, 0xf2, 0x1e, 0xd0, 0x16, 0x70, 0x08,
		0xf2, 0x1e, 0xd0, 0x16, 0x70, 0x08, 0xf2, 0x1e,
		0xd0, 0x16, 0x70, 0x08, 0xf2, 0x1e, 0xd0, 0x16,
		0x70, 0x08, 0xf2, 0x1e, 0xd0, 0x16, 0x70, 0x08,
		0xf2, 0x1e, 0xd0, 0x16, 0x00, 0xee, 0x62, 0x0c,
		0xa5, 0xcf, 0xd0, 0x1c, 0x70, 0x08, 0xf2, 0x1e,
		0xd0, 0x1c, 0x70, 0x08, 0xf2, 0x1e, 0xd0, 0x1c,
		0x70, 0x08, 0xf2, 0x1e, 0xd0, 0x1c, 0x70, 0x08,
		0xf2, 0x1e, 0xd0, 0x1c, 0x70, 0x08, 0xf2, 0x1e,
		0xd0, 0x1c, 0x70, 0x08, 0xf2, 0x1e, 0xd0, 0x1c,
		0x70, 0x08, 0xf2, 0x1e, 0xd0, 0x1c, 0x00, 0xee,
		0x64, 0x01, 0x65, 0x07, 0x62, 0x00, 0x63, 0x00,
		0x60, 0x00, 0x81, 0x30, 0x71, 0x03, 0xd0, 0x11,
		0x71, 0x08, 0xf4, 0x1e, 0xd0, 0x11, 0xf4, 0x1e,
		0x70, 0x08, 0x30, 0x40, 0x15, 0x2a, 0x73, 0x03,
		0x83, 0x52, 0x72, 0x01, 0x32, 0x08, 0x15, 0x28,
		0x00, 0xee, 0x60, 0x00, 0xa5, 0x9a, 0xf0, 0x55,
		0x00, 0xee, 0xa5, 0x9a, 0xf1, 0x65, 0x70, 0x01,
		0x82, 0x00, 0x82, 0x15, 0x4f, 0x01, 0x81, 0x00,
		0xf1, 0x55, 0x00, 0xee, 0xa5, 0x9c, 0xf2, 0x65,
		0xf0, 0x29, 0xd3, 0x45, 0x73, 0x05, 0xf1, 0x29,
		0xd3, 0x45, 0x73, 0x05, 0xf2, 0x29, 0xd3, 0x45,
		0x73, 0x05, 0x62, 0x00, 0xf2, 0x29, 0xd3, 0x45,
		0x00, 0xee, 0xa5, 0x9a, 0xf0, 0x65, 0xa5, 0x9c,
		0xf0, 0x33, 0x25, 0x64, 0x00, 0xee, 0xa5, 0x9b,
		0xf0, 0x65, 0xa5, 0x9c, 0xf0, 0x33, 0x25, 0x64,
		0x00, 0xee, 0x00, 0x64, 0x00, 0x00, 0x00, 0x00,
		0x3c, 0x36, 0x3c, 0x30, 0x30, 0x00, 0xf3, 0xdb,
		0xf3, 0xdb, 0xdb, 0x00, 0xe7, 0x0c, 0xc7, 0x01,
		0xef, 0x00, 0x9e, 0x30, 0x1c, 0x86, 0x3c, 0x00,
		0x1e, 0x30, 0x1c, 0x06, 0x3c, 0x00, 0xf3, 0x66,
		0x67, 0x66, 0x66, 0x00, 0x9e, 0xdb, 0xde, 0xdb,
		0xdb, 0x00, 0x78, 0x30, 0x30, 0x30, 0x30, 0x01,
		0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03,
		0x03, 0x03, 0x01, 0xe7, 0x0d, 0x0d, 0x6f, 0x6d,
		0x6d, 0x6d, 0x6d, 0x6d, 0x6d, 0x6d, 0xed, 0x3f,
		0xb5, 0xb5, 0xb5, 0xb5, 0xb5, 0xb5, 0xb5, 0xb5,
		0xb5, 0xb5, 0xb5, 0x3e, 0xb0, 0xb0, 0xbc, 0xb0,
		0xb0, 0xb0, 0xb0, 0xb0, 0xb0, 0xb0, 0xbe, 0x1c,
		0x36, 0x36, 0x36, 0x36, 0x36, 0x36, 0x36, 0x36,
		0x36, 0x36, 0x1c, 0xdb, 0xdb, 0xdb, 0xdb, 0xdb,
		0xdb, 0xdb, 0xdb, 0xdb, 0xdb, 0x7b, 0x3b, 0xef,
		0x0d, 0x0d, 0xcf, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d,
		0x0d, 0x0d, 0xed, 0x00, 0x80, 0x80, 0x00, 0x80,
		0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x00,
		0x00, 0x0c, 0x11, 0x11, 0x10, 0x00, 0x00, 0x95,
		0x55, 0x95, 0xcd, 0x00, 0x00, 0x53, 0x55, 0x55,
		0x33, 0x40, 0x40, 0x44, 0x42, 0x41, 0x46, 0x00,
		0x40, 0x6a, 0x4a, 0x4a, 0x46, 0x00, 0x20, 0x69,
		0xaa, 0xaa, 0x69, 0x00, 0x00, 0x20, 0x90, 0x88,
		0x30,
	},
}
