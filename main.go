package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	fmt.Println("This is a First Observable : ")
	rxGo_First()
	fmt.Println("This is a second Observable: ")
	rxGo_Second()

}

// A Hot Observable
func rxGo_First() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
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

// A Cold Observable
func rxGo_Second() {
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
