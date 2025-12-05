package main

import (
	"fmt"
	"strconv"
	"math/big"
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

func (fe FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if other == nil {
		return nil, fmt.Errorf("other field element is required")
	}

	if other.prime != fe.prime {
		return nil, fmt.Errorf("cannot add field elements with different primes: %d vs %d", fe.prime, other.prime)
	}

	res := (fe.num + other.num) % fe.prime
	return NewFieldElement(res, fe.prime) 
}

func (fe FieldElement) Subtract(other *FieldElement) (*FieldElement, error) {
	if other == nil {
		return nil, fmt.Errorf("other field element is required")
	}

	if other.prime != fe.prime {
		return nil, fmt.Errorf("cannot subtract field elements with different primes: %d vs %d", fe.prime, other.prime)
	}

	res := (fe.num - other.num) % fe.prime
	return NewFieldElement(res, fe.prime) 
}

func (fe FieldElement) Multiply(other *FieldElement) (*FieldElement, error) {
	res := (fe.num * other.num) % fe.prime
	return NewFieldElement(res, fe.prime) 
}

func (fe FieldElement) Pow(exp int) (*FieldElement, error) {
	if fe.num == 0 {
		if exp <= 0 {
			return nil, fmt.Errorf("undefined: 0^%d (non-positive exponent on zero)", exp)
		}
		// 0^positive = 0
		return NewFieldElement(0, fe.prime)
	}

	pMinus1 := fe.prime - 1

	n := exp % pMinus1
	if n < 0 {
		n += pMinus1
	}

	base := big.NewInt(int64(fe.num))
	power := big.NewInt(int64(n))
	mod := big.NewInt(int64(fe.prime))
	res := new(big.Int).Exp(base, power, mod)

	return NewFieldElement(int(res.Int64()), fe.prime)
}

func (fe FieldElement) Divide(other *FieldElement) (*FieldElement, error) {
	if other == nil {
		return nil, fmt.Errorf("other field element is required")
	}
	if fe.prime != other.prime {
		return nil, fmt.Errorf("cannot divide field elements with different primes: %d vs %d", fe.prime, other.prime)
	}
	if other.num == 0 {
		return nil, fmt.Errorf("division by zero")
	}

	// a / b = a * b^(-1)
	inverse, err := other.Pow(-1) // Now works correctly!
	if err != nil {
		return nil, err
	}
	return fe.Multiply(inverse)
}
