package compiler

import (
	"monkeyc/cmd/ast"
	"monkeyc/cmd/code"
	"monkeyc/cmd/object"
)

//! Compiler is compiler struct
type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

//! New is constructor of Compiler
func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

//! Compile transforms an ast node to byte code instructions
func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

//! Bytecode transforms a compiler struct to a Bytecode struct
func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

//! Bytecode is what we wil pass to the VM
type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}
