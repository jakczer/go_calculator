package calc_test

import (
	"testing"

	"czerwonka.com/calculator/calc"
)

func TestAdd(t *testing.T) {
	if calc.Add(1, 2) != 3 {
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	if calc.Div(2, 2) != 1 {
		t.Fail()
	}
}
