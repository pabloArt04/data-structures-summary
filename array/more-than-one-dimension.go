package main

import "fmt"

func main() {
	// The programming languages add support for multi-dimensional arrays
	// The type [n][m]T is an array of n values of type [m]T
	var data [2][3]uint8

	// Fill the array
	value := uint8(1)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			data[i][j] = value
			value++
		}
	}

	// Data variable is an abstraction, which provides access to data by rows and columns
	fmt.Println(data[0][0], data[0][1], data[0][2])
	fmt.Println(data[1][0], data[1][1], data[1][2])
	fmt.Println()

	// But actually this data is stored in a linear array
	// Example:
	var linealData [2 * 3]uint8
	value = uint8(1)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			linealData[i*3+j] = value
			value++
		}
	}

	// The data is stored in a linear array, using row major order
	fmt.Println(linealData[0], linealData[1], linealData[2])
	fmt.Println(linealData[3], linealData[4], linealData[5])
}
