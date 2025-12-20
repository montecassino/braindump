/***
dining philosophers problem is a classic synchronization problem that leads to deadlock and starvation
since a single philosopher needs to eat with two forks (left and right) and there are only 5 forks,
resource sharing is inevitable.

without a certain technique or medium, all philosophers will try to eat at the same time and try to get fork from one another
like Philosopher A has Fork A but it also needs Fork B to eat but Fork B is being held by Philosopher B and he also needs to eat using Fork A.
This problem leads to a circular dependency of waiting aka "deadlock"

my solution for this is to use the Pigeonhole principle (N minus 1) where total forks (5) - 1 = 4 seats at dining area.
***/

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

func (p Philosopher) eat(wg *sync.WaitGroup, waiterChan chan bool, eatingRounds int) {
	defer wg.Done()

	for i := 0; i < eatingRounds; i++ {		
		fmt.Printf("Philosopher %d is thinking...\n", p.id)
		time.Sleep(time.Millisecond * 10)

		waiterChan <- true

		p.leftFork.Lock()
		time.Sleep(time.Millisecond * 10)
		p.rightFork.Lock()
		
		fmt.Printf("Philosopher %d is eating...\n", p.id)
		time.Sleep(time.Millisecond * 10)

		p.leftFork.Unlock()
		p.rightFork.Unlock()

		<- waiterChan
	}
}

func RunDiningPhilosophers(philosopherCount int, eatingRounds int) {
	count := philosopherCount
	forks := make([]*Fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(Fork)
	}

	philosophers := make([]*Philosopher, count)
	waiterChan := make(chan bool, count - 1)
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{
			id: i, 
			leftFork: forks[i], 
			rightFork: forks[(i+1)%count],
		}
		wg.Add(1)
		go philosophers[i].eat(&wg, waiterChan, eatingRounds)
	}

	wg.Wait()
}