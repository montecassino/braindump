package contextbasedcancellation

import (
	"context"
	"fmt"
	"time"
)

type Result struct {
	Source string
	Data   string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500 * time.Millisecond)

	defer cancel()

	results := FetchAllResults(ctx)

	fmt.Printf("Search complete. Got %d results:\n", len(results))
	for _, res := range results {
		fmt.Printf("- %s: %s\n", res.Source, res.Data)
	}
}

func FetchAllResults(ctx context.Context) []Result {
    resultsChan := make(chan Result, 3) 

    go mockFetch(ctx, "Internal DB", 50*time.Millisecond, resultsChan)
    go mockFetch(ctx, "API A", 200*time.Millisecond, resultsChan)
    go mockFetch(ctx, "API B", 5000*time.Millisecond, resultsChan)

    var finalResults []Result

    for i := 0; i < 3; i++ {
        select {
        case res := <-resultsChan:
            finalResults = append(finalResults, res)
        case <-ctx.Done():
            return finalResults
        }
    }
    return finalResults
}

func mockFetch(ctx context.Context, source string, delay time.Duration, out chan<- Result) {
	select {
	case <-time.After(delay):
		out <- Result{Source: source, Data: "Success"}
	case <- ctx.Done():
		return
	}
}