package main

import (
	"fmt"
	"strconv"
)

type FieldElement struct {
	num   int
	prime int
}

func NewFieldElement(num, prime int) (*FieldElement, error) {
	if num < 0 || num >= prime {
		return nil, fmt.Errorf("num %d is not in field range 0 to %d", num, prime-1)
	}
	return &FieldElement{
		num:   num,
		prime: prime,
	}, nil
}

func (fe FieldElement) String() string {
	return strconv.Itoa(fe.prime) + "_" + strconv.Itoa(fe.num)
}

func (fe FieldElement) Equal(other *FieldElement) bool {
	if other == nil {
		return false
	}
	return fe.num == other.num && fe.prime == other.prime
}
