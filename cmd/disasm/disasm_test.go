package main

import "testing"

func TestDisasmValues(t *testing.T) {
	examples := []struct {
		value    uint16
		nextWord uint16
		width    int
		disasm   string
	}{
		{0x00, 0x0000, 0, "A"},
		{0x01, 0x0000, 0, "B"},
		{0x02, 0x0000, 0, "C"},
		{0x03, 0x0000, 0, "X"},
		{0x04, 0x0000, 0, "Y"},
		{0x05, 0x0000, 0, "Z"},
		{0x06, 0x0000, 0, "I"},
		{0x07, 0x0000, 0, "J"},
		{0x08, 0x0000, 0, "[A]"},
		{0x09, 0x0000, 0, "[B]"},
		{0x0a, 0x0000, 0, "[C]"},
		{0x0b, 0x0000, 0, "[X]"},
		{0x0c, 0x0000, 0, "[Y]"},
		{0x0d, 0x0000, 0, "[Z]"},
		{0x0e, 0x0000, 0, "[I]"},
		{0x0f, 0x0000, 0, "[J]"},
		{0x10, 0x0000, 1, "[0x0000+A]"},
		{0x11, 0x0001, 1, "[0x0001+B]"},
		{0x12, 0x0020, 1, "[0x0020+C]"},
		{0x13, 0x0300, 1, "[0x0300+X]"},
		{0x14, 0x4000, 1, "[0x4000+Y]"},
		{0x15, 0x0005, 1, "[0x0005+Z]"},
		{0x16, 0x0060, 1, "[0x0060+I]"},
		{0x17, 0x0700, 1, "[0x0700+J]"},
		{0x18, 0x0000, 0, "POP"},
		{0x19, 0x0000, 0, "PEEK"},
		{0x1a, 0x0000, 0, "PUSH"},
		{0x1b, 0x0000, 0, "SP"},
		{0x1c, 0x0000, 0, "PC"},
		{0x1d, 0x0000, 0, "O"},
		{0x1e, 0x0123, 1, "[0x0123]"},
		{0x1f, 0x3210, 1, "0x3210"},
		{0x20, 0x0000, 0, "0x00"},
		{0x21, 0x0000, 0, "0x01"},
		{0x22, 0x0000, 0, "0x02"},
		{0x23, 0x0000, 0, "0x03"},
		{0x24, 0x0000, 0, "0x04"},
		{0x25, 0x0000, 0, "0x05"},
		{0x26, 0x0000, 0, "0x06"},
		{0x27, 0x0000, 0, "0x07"},
		{0x28, 0x0000, 0, "0x08"},
		{0x29, 0x0000, 0, "0x09"},
		{0x2a, 0x0000, 0, "0x0a"},
		{0x2b, 0x0000, 0, "0x0b"},
		{0x2c, 0x0000, 0, "0x0c"},
		{0x2d, 0x0000, 0, "0x0d"},
		{0x2e, 0x0000, 0, "0x0e"},
		{0x2f, 0x0000, 0, "0x0f"},
		{0x30, 0x0000, 0, "0x10"},
		{0x31, 0x0000, 0, "0x11"},
		{0x32, 0x0000, 0, "0x12"},
		{0x33, 0x0000, 0, "0x13"},
		{0x34, 0x0000, 0, "0x14"},
		{0x35, 0x0000, 0, "0x15"},
		{0x36, 0x0000, 0, "0x16"},
		{0x37, 0x0000, 0, "0x17"},
		{0x38, 0x0000, 0, "0x18"},
		{0x39, 0x0000, 0, "0x19"},
		{0x3a, 0x0000, 0, "0x1a"},
		{0x3b, 0x0000, 0, "0x1b"},
		{0x3c, 0x0000, 0, "0x1c"},
		{0x3d, 0x0000, 0, "0x1d"},
		{0x3e, 0x0000, 0, "0x1e"},
		{0x3f, 0x0000, 0, "0x1f"},
	}

	for _, example := range examples {
		disasm, width := DisasmValue(example.value, example.nextWord)
		if disasm != example.disasm || width != example.width {
			t.Errorf("0x%02x was %d '%s' not %d '%s'", example.value, width, disasm, example.width, example.disasm)
		}
	}
}

func TestDisasmInstructions(t *testing.T) {
	examples := []struct {
		words    []uint16
		width    int
		disasm   string
	}{
		{[]uint16{0x0001, 0x0000, 0x0000}, 1, "SET A, A"},
		{[]uint16{0x0002, 0x0000, 0x0000}, 1, "ADD A, A"},
		{[]uint16{0x0003, 0x0000, 0x0000}, 1, "SUB A, A"},
		{[]uint16{0x0004, 0x0000, 0x0000}, 1, "MUL A, A"},
		{[]uint16{0x0005, 0x0000, 0x0000}, 1, "DIV A, A"},
		{[]uint16{0x0006, 0x0000, 0x0000}, 1, "MOD A, A"},
		{[]uint16{0x0007, 0x0000, 0x0000}, 1, "SHL A, A"},
		{[]uint16{0x0008, 0x0000, 0x0000}, 1, "SHR A, A"},
		{[]uint16{0x0009, 0x0000, 0x0000}, 1, "AND A, A"},
		{[]uint16{0x000a, 0x0000, 0x0000}, 1, "BOR A, A"},
		{[]uint16{0x000b, 0x0000, 0x0000}, 1, "XOR A, A"},
		{[]uint16{0x000c, 0x0000, 0x0000}, 1, "IFE A, A"},
		{[]uint16{0x000d, 0x0000, 0x0000}, 1, "IFN A, A"},
		{[]uint16{0x000e, 0x0000, 0x0000}, 1, "IFG A, A"},
		{[]uint16{0x000f, 0x0000, 0x0000}, 1, "IFB A, A"},

		{[]uint16{instruction(0x1, 0x00, 0x01), 0x0000, 0x0000}, 1, "SET A, B"},
		{[]uint16{instruction(0x1, 0x08, 0x09), 0x0000, 0x0000}, 1, "SET [A], [B]"},

		{[]uint16{instruction(0x1, 0x10, 0x01), 0x0123, 0x4567}, 2, "SET [0x0123+A], B"},
		{[]uint16{instruction(0x1, 0x00, 0x11), 0x0123, 0x4567}, 2, "SET A, [0x0123+B]"},

		{[]uint16{instruction(0x1, 0x10, 0x11), 0x0123, 0x4567}, 3, "SET [0x0123+A], [0x4567+B]"},
		{[]uint16{instruction(0x1, 0x1e, 0x1e), 0x0123, 0x4567}, 3, "SET [0x0123], [0x4567]"},
		{[]uint16{instruction(0x1, 0x1f, 0x1f), 0x0123, 0x4567}, 3, "SET 0x0123, 0x4567"},

		{[]uint16{instruction(0x0, 0x01, 0x23), 0x0123, 0x4567}, 1, "JSR 0x23"},
	}

	for _, example := range examples {
		width, disasm := DisasmInstruction(example.words)
		if disasm != example.disasm || width != example.width {
			t.Errorf("0x%04x 0x%04x 0x%04x was %d '%s' not %d '%s'", example.words[0], example.words[1], example.words[2], width, disasm, example.width, example.disasm)
		}
	}
}

func instruction(op, a, b uint16) uint16 {
	return (op & 0xf) + (a & 0x3f) << 4 + (b & 0x3f) << 10
}
