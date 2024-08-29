package vm

import (
	"github.com/natac13/monkey-compiler/internal/code"
	"github.com/natac13/monkey-compiler/internal/object"
)

type Frame struct {
	cl *object.Closure
	// pointer index of this frame
	ip int
	// basePointer is the index of the bottom of the current frame
	basePointer int
}

func NewFrame(cl *object.Closure, basePointer int) *Frame {
	return &Frame{cl: cl, ip: -1, basePointer: basePointer}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
