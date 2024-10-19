// Copyright 2024 tabuyos. All rights reserved.
//
// @author tabuyos
// @since 2024/10/10
// @description file desc
package chan_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_Chan(t *testing.T) {
	ch := make(chan string)

	go func() {
		for v := range ch {
			fmt.Printf("1 >: %s\n", v)
		}
	}()

	go func() {
		for v := range ch {
			fmt.Printf("2 >: %s\n", v)
		}
	}()

	for v := range 10 {
		ch <- strconv.Itoa(v)
	}

	time.Sleep(1 * time.Second)
}
