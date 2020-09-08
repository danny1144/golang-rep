package snow

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total1 uint64

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Print(total1)
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i < 100; i++ {
		atomic.AddUint64(&total1, i)
	}

}
