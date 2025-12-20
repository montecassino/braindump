package deadlock

import (
	"fmt"
	"sync"
	"time"
)

type Fork struct{ sync.Mutex }

type Philosopher struct {
	id                  int
	leftFork *Fork
	rightFork *Fork
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		// TODO: Implement the logic to pick up forks, eat, 
		// and put them down safely.
		
		fmt.Printf("Philosopher %d is thinking...\n", p.id)
		time.Sleep(time.Millisecond * 500)

		// Hint: How do you prevent everyone from 
		// grabbing the left fork at the exact same time?
		
		fmt.Printf("Philosopher %d is eating...\n", p.id)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	count := 5
	forks := make([]*Fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(Fork)
	}

	philosophers := make([]*Philosopher, count)
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{
			id: i, 
			leftFork: forks[i], 
			rightFork: forks[(i+1)%count],
		}
		wg.Add(1)
		go philosophers[i].eat(&wg)
	}

	wg.Wait()
}