package chip8

type IntScreen struct {
	Rows [32]uint64
}

func (s *IntScreen) IsSet(x, y int) bool {
	mask := uint64(1) << uint(63-x)
	return s.Rows[y]&mask != 0
}

func (s *IntScreen) Size() (int, int) {
	return 64, 32
}
