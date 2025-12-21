package fanoutin

import (
	"fmt"
	"sync"
	"time"
)

type Image struct {
	ID   int
	Path string
}

func main() {
	imageIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	jobs := make(chan int, len(imageIDs))

	results := make(chan Image, len(imageIDs))
	
	for res := range results {
		fmt.Printf("Saved Thumbnail for Image %d at %s\n", res.ID, res.Path)
	}
}


func worker(id int, jobs <-chan int, results chan<- Image, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing image %d\n", id, job)
		time.Sleep(500 * time.Millisecond) 
		results <- Image{ID: job, Path: fmt.Sprintf("/thumbs/%d.jpg", job)}
	}
}