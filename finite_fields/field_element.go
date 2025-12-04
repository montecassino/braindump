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

func (fe FieldElement) Pow(power uint) (*FieldElement, error) {
	base := big.NewInt(int64(fe.num))
	exp := big.NewInt(int64(power))
	mod := big.NewInt(int64(fe.prime))
	
	res := new(big.Int).Exp(base, exp, mod)
	return NewFieldElement(int(res.Int64()), fe.prime)
}

func (fe FieldElement) Divide(other *FieldElement) (*FieldElement, error) {
	if other == nil {
		return nil, fmt.Errorf("other field element is required")
	}

	if other.prime != fe.prime {
		return nil, fmt.Errorf("cannot multiply field elements with different primes: %d vs %d", fe.prime, other.prime)
	}

	if other.num == 0 {
		return nil, fmt.Error("division by zero")
	}

	exponent := uint(fe.prime - 2)
	inverse, err := other.Pow(exponent)
    if err != nil {
        return nil, fmt.Errorf("failed to compute inverse: %w", err)
    }

	product, err := fe.Multiply(inverse)
	if err != nil {
        return nil, fmt.Errorf("multiplication failed: %w", err)
    }

    return result, nil
}