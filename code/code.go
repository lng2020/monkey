package code

import (
	"encoding/binary"
	"fmt"
)

const (
	OpConstant Opcode = iota
)

var definition = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

type Instructions []byte

type Opcode byte

type Definition struct {
	Name          string
	OperandWidths []int
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definition[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opconde %d undefined", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definition[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}
