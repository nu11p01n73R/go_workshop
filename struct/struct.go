package main

import "fmt"

type Movie struct {
	Name string
	Year int
}

func main() {
	movie1 := Movie{"Fight Club", 1999}
	fmt.Println(movie1)

	movie2 := Movie{}
	movie2.Name = "The Godfather"
	movie2.Year = 1975

	fmt.Println("Name : ", movie2.Name)
	fmt.Println("Year : ", movie2.Year)

	movie3 := movie2
	fmt.Println("Movie 3 Name : ", movie3.Name)
	fmt.Println("Movie 3 Year : ", movie3.Year)

	movie3.Year = 1970
	fmt.Println(movie3, movie2)

	movie4 := &movie2
	movie4.Year = 1960
	fmt.Println(movie2, movie4)
}
