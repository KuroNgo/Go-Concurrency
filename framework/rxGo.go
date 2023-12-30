package framework

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

// RxGoFirst A Hot Observable
func RxGoFirst() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	// creates a cold observable from a channel with FromChannel
	observable := rxgo.FromChannel(ch)

	// First Observer
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	// Second Observer
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

}

// RxGoSecond A Cold Observable
func RxGoSecond() {
	// on the other hand, let's create a cold Observable using Defer operator
	observable := rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}})

	// First Observable
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	// Second Observable
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
