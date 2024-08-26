package vm

import (
	"github.com/natac13/monkey-compiler/internal/code"
	"github.com/natac13/monkey-compiler/internal/object"
)

type Frame struct {
	// points to the compiled function referenced by the frame
	fn *object.CompiledFunction
	// pointer index of this frame
	ip int
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
