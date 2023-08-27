package test

import (
	"fmt"
	"testing"
)

func TestBing(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := make([]int, len(a))
	vResults := make(chan struct {
		Index int
		Info  int
	}, len(a))
	for i := 0; i < len(a); i++ {
		go func(i int) {
			vResults <- struct {
				Index int
				Info  int
			}{Index: i, Info: i}
		}(i)
	}
	for i := 0; i < len(a); i++ {
		result := <-vResults
		b[result.Index] = result.Info
	}
	fmt.Println(b)
}
