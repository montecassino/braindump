package deadlock

import (
	"testing"
	"time"
)

func TestDiningPhilosophers(t *testing.T) {
	philosopherCount := 5
	eatingRounds := 3

	done := make(chan bool)

	go func() {
		RunDiningPhilosophers(philosopherCount, eatingRounds)
		done <- true
	}()

	select {
	case <-done:
		t.Log("All philosophers finished eating without deadlock.")
	case <-time.After(10 * time.Second):
		t.Error("Test timed out! Possible Deadlock detected.")
	}
}