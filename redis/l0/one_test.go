// Copyright 2024 tabuyos. All rights reserved.
//
// @author tabuyos
// @since 2024/09/26
// @description file desc
package l0_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func Test_One(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Ping(context.Background()).Err()

	if err != nil {
		fmt.Println(err)
	}

	fn, ch := MakeChannel()

	go func() {
		for msgCh := range client.Subscribe(context.Background(), "test-one").Channel() {
			fn(msgCh.Payload)
		}
	}()

	go func() {
		fmt.Println("starting")
		for msg := range ch {
			fmt.Println(msg)
		}
	}()

	time.Sleep(3 * time.Second)

	for i := range 5 {
		client.Publish(context.Background(), "test-one", fmt.Sprintf("hello-%d", i))
	}
}

func MakeChannel() (func(string), <-chan any) {
	ch := make(chan any)
	fn := func(msg string) {
		ch <- msg
	}
	return fn, ch
}
