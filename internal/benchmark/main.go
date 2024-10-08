package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/natac13/monkey-compiler/internal/compiler"
	"github.com/natac13/monkey-compiler/internal/evaluator"
	"github.com/natac13/monkey-compiler/internal/lexer"
	"github.com/natac13/monkey-compiler/internal/object"
	"github.com/natac13/monkey-compiler/internal/parser"
	"github.com/natac13/monkey-compiler/internal/vm"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")
var fibNum = flag.Int("fib", 15, "fibonacci number to calculate")

// var input = `
// let fibonacci = fn(x) {
// 	if (x == 0) {
// 		0
// 	} else {
// 		if (x == 1) {
// 			return 1;
// 		} else {
// 			fibonacci(x - 1) + fibonacci(x - 2);
// 		}
// 	}
// };
// fibonacci(30);
// `

func main() {
	flag.Parse()

	var duration time.Duration
	var result object.Object

	input := fmt.Sprintf(`
		let fibonacci = fn(x) {
			if (x == 0) {
				0
			} else {
				if (x == 1) {
					return 1;
				} else {
					fibonacci(x - 1) + fibonacci(x - 2);
				}
			}
		};
		fibonacci(%d);
	`, *fibNum)

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if *engine == "vm" {
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Printf("compiler error: %s", err)
			return
		}

		machine := vm.New(comp.ByteCode())
		start := time.Now()

		err = machine.Run()
		if err != nil {
			fmt.Printf("vm error: %s", err)
			return
		}

		duration = time.Since(start)
		result = machine.LastPoppedStackElem()
	} else {
		env := object.NewEnvironment()
		start := time.Now()
		result = evaluator.Eval(program, env)
		duration = time.Since(start)
	}

	fmt.Printf(
		"engine=%s, result=%s, duration=%s\n",
		*engine,
		result.Inspect(),
		duration,
	)
}
