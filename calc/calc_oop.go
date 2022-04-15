package calc

import (
	"fmt"

	cts "czerwonka.com/calculator/consts"
)

type Calc struct {
	A int
	B int
}

type CalcPrinter struct {
	C Calc
}

func (cr CalcPrinter) PrintResult(o cts.Operation) {
	switch o {
	case cts.ADD:
		fmt.Println(cr.C.Add())
	case cts.SUB:
		fmt.Println(cr.C.Sub())
	case cts.MUL:
		fmt.Println(cr.C.Mul())
	case cts.DIV:
		fmt.Println(cr.C.Div())
	case cts.EQ:
		fmt.Println(cr.C.Eqaul())
	case cts.NEQ:
		fmt.Println(cr.C.NotEqual())
	case cts.LEQ:
		fmt.Println(cr.C.Leq())
	case cts.LT:
		fmt.Println(cr.C.Lt())
	case cts.GT:
		fmt.Println(cr.C.Gt())
	case cts.GEQ:
		fmt.Println(cr.C.Geq())
	default:
		panic("UNEXPECTED TYPE")
	}
}

func (c Calc) Add() int {
	return c.A + c.B
}

func (c Calc) Div() int {
	return c.A / c.B
}

func (c Calc) Mul() int {
	return c.A * c.B
}

func (c Calc) Sub() int {
	return c.A - c.B
}

func (c Calc) Mod() int {
	return c.A % c.B
}

func (c Calc) Eqaul() bool {
	return c.A == c.B
}

func (c Calc) NotEqual() bool {
	return c.A != c.B
}

func (c Calc) Lt() bool {
	return c.A < c.B
}

func (c Calc) Gt() bool {
	return c.A > c.B
}

func (c Calc) Leq() bool {
	return c.A <= c.B
}

func (c Calc) Geq() bool {
	return c.A >= c.B
}
