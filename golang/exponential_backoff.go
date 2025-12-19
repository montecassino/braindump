package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds do not retry")
	ErrTransient         = errors.New("temporary server error please retry")
	ErrTimeout           = errors.New("request timed out")
)

func UnreliableAPI(ctx context.Context) error {
	select {
	case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
	case <-ctx.Done():
		return ctx.Err()
	}

	outcomes := []error{nil, ErrTransient, ErrTransient, ErrInsufficientFunds}
	return outcomes[rand.Intn(len(outcomes))]
}

func ProcessPaymentWithRetry(ctx context.Context) error {
	iCtx, cancel := context.WithTimeout(ctx, 2500 * time.Millisecond)
	defer cancel()

	maxTry := 3
	var lastErr error 

	for i := 0; i < maxTry; i++ {
		if i > 0 {
			backoff := time.Duration(i) * 500 * time.Millisecond
			select {
			case <-time.After(backoff):
			case <-iCtx.Done():
				return fmt.Errorf("operation timed out during backoff: %w", iCtx.Err())
			}
		}

		err := UnreliableAPI(iCtx)

		if err == nil {
			return nil
		}

		if errors.Is(err, ErrInsufficientFunds) {
			return err
		}

		lastErr = err 
	}

	return fmt.Errorf("all retries failed. last error: %w", lastErr)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ProcessPaymentWithRetry(ctx)
	if err != nil {
		fmt.Printf("Payment Failed: %v\n", err)
	} else {
		fmt.Println("Payment Successful!")
	}
}