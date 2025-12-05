package main

import (
	"fmt"
)

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {

	mystack := stack[string]{
		elements: []string{"harsh"},
	}

	fmt.Println(mystack)

	nums := []int{1, 2, 3}
	names := []string{"harsh", "shiva"}
	printSlice(names)
	printSlice(nums)
}
