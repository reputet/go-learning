package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	value_old := 0
	value_new := 1
	return func() int {
		buffer := value_old
		value_old = value_new
		value_new = value_new + buffer
		if (value_new == 1) {
			return 1
		}
		return value_new
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
