package calc_test

import (
	"testing"

	"czerwonka.com/calculator/calc"
)

func TestOopAdd(t *testing.T) {
	if (calc.Calc{1, 2}.Add() != 3) {
		t.Fatal("add not working")
	}
}

func TestOopDiv(t *testing.T) {
	if (calc.Calc{2, 2}.Div() != 1) {
		t.Fatal("div not working")
	}
}
