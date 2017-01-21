package gameboy

import "testing"

func TestCBOps(t *testing.T) {
	tests := []struct {
		op  uint8
		af1 uint16
		af2 uint16
	}{
		{0x07, 0x0000, 0x0080}, {0x07, 0x02B0, 0x0400}, {0x07, 0xFFB0, 0xFF10}, {0x07, 0xFEB0, 0xFD10},
		{0x0F, 0x0000, 0x0080}, {0x0F, 0x02B0, 0x0100}, {0x0F, 0xFFB0, 0xFF10}, {0x0F, 0xFEB0, 0x7F00},

		{0x17, 0x0000, 0x0080}, {0x17, 0x02B0, 0x0500}, {0x17, 0xFFB0, 0xFF10}, {0x17, 0xFEB0, 0xFD10},
		{0x1F, 0x0000, 0x0080}, {0x1F, 0x02B0, 0x8100}, {0x1F, 0xFFB0, 0xFF10}, {0x1F, 0xFEB0, 0xFF00},

		{0x27, 0x0000, 0x0080}, {0x27, 0x02B0, 0x0400}, {0x27, 0xFFB0, 0xFE10}, {0x27, 0xFEB0, 0xFC10},
		{0x2F, 0x0000, 0x0080}, {0x2F, 0x02B0, 0x0100}, {0x2F, 0xFFB0, 0xFF10}, {0x2F, 0xFEB0, 0xFF00},
	}

	rom := make([]byte, 0x2000)
	for _, test := range tests {
		rom[0x100] = 0xCB
		rom[0x101] = test.op
		rom[0x102] = 0x10

		gb := NewMachine(ROM(rom), false)

		gb.cpu.setAF(test.af1)
		gb.StepUntilStop()

		if gb.cpu.af() != test.af2 {
			t.Errorf("(op=%02x, af=%04x) expected af=%04x, got af=%04x", test.op, test.af1, test.af2, gb.cpu.af())
		}
	}
}
