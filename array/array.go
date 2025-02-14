package main

import "fmt"

func main() {
	// An array is a fixed-size sequence of elements of a single type
	// The type [n]T is an array of n values of type T
	var data [5]uint8

	// Set the value of the array
	// The index is 0-based
	data[0] = 1
	data[1] = 2
	data[2] = 3
	data[3] = 4
	data[4] = 5

	// The same as
	// for i := 0; i < len(data); i++ {
	// 	data[i] = i
	// }

	// the arrays use contiguous memory
	// That means the memory is allocated in a single block
	for i := 0; i < 5; i++ {
		fmt.Println(&data[i])
	}
	// the difference between the memory addresses is 1 byte, because the type is uint8

	// if you try to access an index out of the array bounds, you will get a runtime error
	// fmt.Println(data[5])
}
