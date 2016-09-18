package main

import "fmt"

func main() {
	aMap := make(map[string]string)

	aMap["first_name"] = "Johny"
	aMap["last_name"] = "James"

	fmt.Println("aMap:", aMap)
	fmt.Println("aMap length", len(aMap))

	firstName := aMap["first_name"]
	fmt.Println("first name:", firstName)

	_, pre := aMap["test_key"]
	fmt.Println("test_key in aMap :", pre)

	_, pre = aMap["first_name"]
	fmt.Println("first name in aMap :", pre)

	bMap := map[string]int{"a": 1, "b": 2}
	fmt.Println("bMap: ", bMap)
}
