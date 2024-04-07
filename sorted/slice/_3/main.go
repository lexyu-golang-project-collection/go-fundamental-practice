package main

import "fmt"

func main() {
	movies := []string{"Avatar", "Titanic", "Inception", "Interstellar", "The Dark Knight"}

	allmovies := movies[:]

	fmt.Println(allmovies)
}
