package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	book.Pages = pages
}

func main() {

	/*
		Create 3 Book Structs with the following data:

		Book 1:
		Title: "The Great Gatsby"
		Author: "F. Scott Fitzgerald"
		Pages: 180

		Book 2
		Title: "To Kill a Mockingbird"
		Author: "Harper Lee"
		Pages: 281

		Book 3
		Title: "Pride and Prejudice"
		Author: "Jane Austen"
		Pages: 279
	*/

	book1 := &Book{"The Great Gatsby", "F. Scott Fitzgerald", 180}
	book2 := &Book{"To Kill a Mockingbird", "Harper Lee", 281}
	book3 := &Book{"Pride and Prejudice", "Jane Austen", 279}

	/*
		Update the information for Books as following:

		Book 1: Updates Page Count to 210
		Book 2: Updates Page Count to 250
		Book 3: Updates Page Count to 295
	*/
	updatePages(book1, 210)
	updatePages(book2, 250)
	updatePages(book3, 295)

	/*
		Print all the struct objects
		fmt.Println(book)
	*/
	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book3)

}
