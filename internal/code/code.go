package code

import (
	"encoding/binary"
	"fmt"
)

const (
	OpConstant Opcode = iota
)

type Instructions []byte

type Opcode byte

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	// check if the opcode is defined in the definitions map
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	// calculate the length of the instruction
	// the byte length of the instruction is the sum of the length of the operands and the opcode, which is 1 byte
	instructionLen := 1
	for _, width := range def.OperandWidths {
		instructionLen += width
	}

	instruction := make([]byte, instructionLen)
	// set the opcode
	instruction[0] = byte(op)

	// set the operands
	offset := 1
	for i, operand := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
		}
		offset += width
	}

	return instruction
}
