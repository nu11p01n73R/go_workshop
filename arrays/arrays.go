package main

import "fmt"

func main() {
	var intArr [5]int
	fmt.Println(intArr)

	intArr[0] = 1
	intArr[2] = 3
	fmt.Println(intArr)

	bArr := [4]int{1, 2, 3, 4}
	fmt.Println(bArr[2:3])

	strArr := [3]string{"apple", "orange", "grapes"}
	fmt.Println(strArr)
}
