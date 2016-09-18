package main

import "fmt"

func main() {
	var aSlice = []string{"a", "hello", "hai"}
	fmt.Println("aSlice:", aSlice)

	bSlice := make([]int, 5)
	bSlice[1] = 1
	bSlice[3] = 3
	fmt.Println("bSlice", bSlice)

	bSlice = append(bSlice, 6, 7)
	fmt.Println(bSlice, len(bSlice))

	//Slicing of arrays and slices
	fmt.Println(bSlice[1:4])
	fmt.Println(bSlice[1:])
	fmt.Println(bSlice[:4])
	fmt.Println(bSlice[:])

	aArray := [3]int{1, 2, 3}
	bArray := aArray
	bArray[2] = 20
	fmt.Println(aArray, bArray)

	cSlice := []int{1, 2, 3}
	dSlice := cSlice
	dSlice[2] = 20
	fmt.Println(cSlice, dSlice)

	for index, val := range bSlice {
		fmt.Println(index, val)
	}
}
