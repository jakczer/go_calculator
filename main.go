package main

import (
	"czerwonka.com/calculator/calc"
	"czerwonka.com/calculator/consts"
)

func main() {
	op := consts.UNKOWN_FOR_EXCEPTION
	bc := calc.CalcPrinter{calc.Calc{1, 3}}
	bc.PrintResult(op)
}
