package consts

type Operation int8

const (
	ADD Operation = iota
	SUB
	DIV
	MUL
	EQ
	NEQ
	LEQ
	LT
	GT
	GEQ
	UNKOWN_FOR_EXCEPTION
)
