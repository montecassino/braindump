// fanout/fanin, race pattern (1st response wins)

package golang

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Result struct {
	Engine string
	Data   string
}


func MockSearch(ctx context.Context, name string) (string, error) {
	delay := time.Duration(rand.Intn(500)+100) * time.Millisecond
	
	select {
	case <-time.After(delay):
		return fmt.Sprintf("Result from %s", name), nil
	case <-ctx.Done():
		fmt.Printf("Provider %s: Received cancellation, stopping work.\n", name)
		return "", ctx.Err()
	}
}

func AggregateSearch(ctx context.Context) (Result, error) {
	iCtx, cancel := context.WithTimeout(ctx, 300 * time.Millisecond)

	defer cancel()

	resChan := make(chan Result, 3)

	for i := 0; i < 3; i++ {
		go func(){
			res, err := MockSearch(iCtx, "test name" + strconv.Itoa(i))

			if (err != nil) {
				return
			}

			select {
			case resChan <- Result{Engine: strconv.Itoa(i), Data: res}:
			// listen to time deadline
			case <- iCtx.Done():
				return
			}

	
		}()
	}

	select {
	case res := <- resChan:
		return res, nil
	case <- iCtx.Done():
        return Result{}, fmt.Errorf("search timed out: %w", iCtx.Err())
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Starting aggregated search...")
	start := time.Now()

	res, err := AggregateSearch(context.Background())

	if err != nil {
		fmt.Printf("Search failed: %v\n", err)
	} else {
		fmt.Printf("Final Winning Result: %s (Took %v)\n", res.Data, time.Since(start))
	}

	time.Sleep(1 * time.Second)
}