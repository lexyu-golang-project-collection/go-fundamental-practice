package main

import "fmt"

func main() {
	movies := []string{"Avatar", "Titanic", "Inception", "Interstellar", "The Dark Knight"}

	part_movies := movies[:2]

	fmt.Println(part_movies)
}
