package l1

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type NetWork struct {
}

type Contract struct {
	Name string
}

func GenContract(count int) []Contract {
	contracts := make([]Contract, count)
	return contracts
}

func TestTransformNetWork(t *testing.T) {
	ch := make(chan int, 1)

	go func() {
		wg := sync.WaitGroup{}
		for v := range 10 {
			wg.Add(1)
			go func() {
				if v == 4 {
					time.Sleep(time.Second)
				}
				ch <- v
				wg.Done()
			}()
		}
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("handle is ok")
}
