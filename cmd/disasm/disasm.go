package main

import (
	"fmt"
)

func DisasmInstruction(words []uint16) (width int, disasm string) {
	opcode := words[0] & 0xf
	a := (words[0] >> 4) & 0x3f
	b := (words[0] >> 10) & 0x3f

	switch {
	case opcode == 0:
		switch a {
		case 0x01:
			return 1, fmt.Sprintf("JSR 0x%02x", b)
		}
	case opcode != 0:
		count := 1
		ad, aw := DisasmValue(a, words[count])
		count += aw
		bd, bw := DisasmValue(b, words[count])
		count += bw
		return count, fmt.Sprintf("%s %s, %s", BasicOpcodeName(opcode), ad, bd)
	}
	return 1, "" // TODO
}

func BasicOpcodeName(opcode uint16) string {
	switch opcode {
	case 0x1:
		return "SET"
	case 0x2:
		return "ADD"
	case 0x3:
		return "SUB"
	case 0x4:
		return "MUL"
	case 0x5:
		return "DIV"
	case 0x6:
		return "MOD"
	case 0x7:
		return "SHL"
	case 0x8:
		return "SHR"
	case 0x9:
		return "AND"
	case 0xa:
		return "BOR"
	case 0xb:
		return "XOR"
	case 0xc:
		return "IFE"
	case 0xd:
		return "IFN"
	case 0xe:
		return "IFG"
	case 0xf:
		return "IFB"
	}

	return "" // TODO
}

func DisasmValue(value, nextWord uint16) (disasm string, nextUsed int) {
	registers := []string{"A", "B", "C", "X", "Y", "Z", "I", "J"}
	switch {
	case value >= 0x00 && value <= 0x07:
		return registers[value], 0
	case value >= 0x08 && value <= 0x0f:
		return "[" + registers[value-0x08] + "]", 0
	case value >= 0x10 && value <= 0x17:
		return fmt.Sprintf("[0x%04x+%s]", nextWord, registers[value-0x10]), 1
	case value == 0x18:
		return "POP", 0
	case value == 0x19:
		return "PEEK", 0
	case value == 0x1a:
		return "PUSH", 0
	case value == 0x1b:
		return "SP", 0
	case value == 0x1c:
		return "PC", 0
	case value == 0x1d:
		return "O", 0
	case value == 0x1e:
		return fmt.Sprintf("[0x%04x]", nextWord), 1
	case value == 0x1f:
		return fmt.Sprintf("0x%04x", nextWord), 1
	case value >= 0x20 && value <= 0x3f:
		return fmt.Sprintf("0x%02x", value-0x20), 0
	}

	return "", 0 // TODO: Error or error value?
}
