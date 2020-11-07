package code

import (
	"encoding/binary"
	"fmt"
)

//! Instructions for monkeyc
type Instructions []byte

//! Opcode for monkeyc
type Opcode byte

//* opcode definition
const (
	//! Opconstant has 3 bytecode
	//! first bytecode is nstructions, and the other two bytecode are constant by encoding BignEndian
	OpConstant Opcode = iota
)

//! Definition is struct of opcode definition
type Definition struct {
	Name          string
	OperandWidths []int //* number of bytes each operand takes up
}

//* definitions is map of Definitions
//* it records Opcodes
var definitions = map[Opcode]*Definition{
	OpConstant: {"Opconstant", []int{2}},
}

//! Lookup is function that looks for an opcode definition
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

//! Make is a function that crates a single bytecode
func Make(op Opcode, oprands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w // allocvate length to wrrte byte code
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range oprands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += 1
	}

	return instruction
}
