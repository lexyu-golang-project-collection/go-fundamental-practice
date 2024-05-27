package main

import "fmt"

func main() {
	grades := []int{10, 20, 30}

	movies := []string{"Avatar", "Titanic", "Inception", "Interstellar", "The Dark Knight"}

	slice := []int{1, 2, 3, 4, 5}[:3]

	fmt.Println(grades)
	fmt.Println(movies)
	fmt.Println(slice)

	fmt.Println("============================================================")

	allmovies := movies[:]
	fmt.Println(allmovies)

	fmt.Println("============================================================")

	part_movies := movies[:2]
	fmt.Println(part_movies)
}
