package games

import "time"

var Airplane = &Game{
	Name:        "Airplane",
	Description: "Air-Supply for the Masses",
	HowToPlay: //
	`Use the Space key to drop a parcel
	but do not hit any of the other
	airplanes!`,
	ClockSpeed:           10 * time.Millisecond,
	InstructionsPerCycle: 4,
	Keys: map[GameKey]uint8{
		Space: 8,
		Enter: 8,
	},
	BackgroundColor: Color{90, 170, 225},
	ForegroundColor: Color{115, 50, 0},
	Program: []byte{
		0x6a, 0x00, 0x6b, 0x04, 0x6c, 0x01, 0x6d, 0x00,
		0x6e, 0x02, 0x23, 0x26, 0x23, 0x20, 0x60, 0x30,
		0x61, 0x01, 0xf0, 0x15, 0xf0, 0x07, 0xf1, 0x18,
		0x30, 0x00, 0x12, 0x14, 0x22, 0x42, 0x23, 0x20,
		0x7d, 0x01, 0x23, 0x20, 0x60, 0x08, 0xe0, 0xa1,
		0x23, 0x0a, 0x4a, 0x00, 0x12, 0x3e, 0xa3, 0x62,
		0xd8, 0x91, 0x79, 0x01, 0xd8, 0x91, 0x4f, 0x01,
		0x12, 0xf4, 0x49, 0x18, 0x12, 0xe4, 0x22, 0xb2,
		0x12, 0x1e, 0x4c, 0x01, 0x22, 0x6c, 0x4c, 0x02,
		0x22, 0x7a, 0x4c, 0x03, 0x22, 0x88, 0x4c, 0x04,
		0x22, 0x96, 0x4c, 0x05, 0x22, 0xa4, 0xa3, 0x59,
		0xd6, 0x72, 0x44, 0x00, 0x00, 0xee, 0xa3, 0x57,
		0xd4, 0x52, 0x42, 0x00, 0x00, 0xee, 0xa3, 0x5b,
		0xd2, 0x32, 0x00, 0xee, 0x66, 0x28, 0x67, 0x09,
		0x64, 0x00, 0x65, 0x00, 0x62, 0x00, 0x63, 0x00,
		0x00, 0xee, 0x66, 0x28, 0x67, 0x0e, 0x64, 0x28,
		0x65, 0x14, 0x62, 0x00, 0x63, 0x00, 0x00, 0xee,
		0x66, 0x28, 0x67, 0x07, 0x64, 0x28, 0x65, 0x0c,
		0x62, 0x16, 0x63, 0x11, 0x00, 0xee, 0x66, 0x28,
		0x67, 0x07, 0x64, 0x28, 0x65, 0x0e, 0x62, 0x16,
		0x63, 0x14, 0x00, 0xee, 0x66, 0x28, 0x67, 0x05,
		0x64, 0x28, 0x65, 0x10, 0x62, 0x16, 0x63, 0x0b,
		0x00, 0xee, 0xa3, 0x59, 0xd6, 0x72, 0x76, 0xfe,
		0xd6, 0x72, 0x44, 0x00, 0x00, 0xee, 0xa3, 0x57,
		0xd4, 0x52, 0x74, 0x02, 0x44, 0x44, 0x74, 0xc0,
		0xd4, 0x52, 0x42, 0x00, 0x00, 0xee, 0xa3, 0x5b,
		0xd2, 0x32, 0x72, 0x02, 0x4c, 0x04, 0x72, 0x02,
		0x4c, 0x05, 0x72, 0x02, 0x42, 0x44, 0x72, 0xc0,
		0xd2, 0x32, 0x00, 0xee, 0x7c, 0x01, 0x6d, 0x00,
		0x6e, 0x02, 0x00, 0xe0, 0x4c, 0x06, 0x6c, 0x01,
		0x6a, 0x00, 0x12, 0x0a, 0x60, 0x06, 0xf0, 0x18,
		0x7b, 0xff, 0x4b, 0x00, 0x13, 0x08, 0x6d, 0x00,
		0x6e, 0x02, 0x00, 0xe0, 0x6a, 0x00, 0x12, 0x0a,
		0x13, 0x08, 0x4a, 0x01, 0x00, 0xee, 0x60, 0x02,
		0xf0, 0x18, 0x6a, 0x01, 0x88, 0xd0, 0x78, 0x01,
		0x89, 0xe0, 0x79, 0x01, 0xd8, 0x91, 0x00, 0xee,
		0xa3, 0x54, 0xdd, 0xe2, 0x00, 0xee, 0x64, 0x19,
		0x63, 0x00, 0xa3, 0x56, 0xd3, 0x41, 0x73, 0x08,
		0x33, 0x40, 0x13, 0x2c, 0x63, 0x1e, 0x64, 0x1b,
		0xfc, 0x29, 0xd3, 0x45, 0x4b, 0x04, 0xa3, 0x5f,
		0x4b, 0x03, 0xa3, 0x60, 0x4b, 0x02, 0xa3, 0x61,
		0x4b, 0x01, 0xa3, 0x62, 0x63, 0x01, 0x74, 0x02,
		0xd3, 0x41, 0x00, 0xee, 0x80, 0xf8, 0xff, 0x80,
		0xe0, 0x10, 0x70, 0x88, 0xee, 0x11, 0x77, 0xaa,
		0xa8, 0xa0, 0x80, 0x00,
	},
}
