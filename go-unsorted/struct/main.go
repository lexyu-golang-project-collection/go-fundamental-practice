package main

import "fmt"

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

type O1 struct {
	x int
}

type O2 struct {
	x int
}

type Movie struct {
	name   string
	rating float32
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
		name:   s,
		rating: r,
	}
	return
}

func calcArea(c *Circle) {
	const PI float64 = 3.14
	var area float64
	area = PI * c.radius * c.radius
	c.area = area
}

func main() {
	var s Student
	fmt.Printf("%+v\n", s)

	s1 := new(Student) // pointer
	fmt.Printf("%+v\n", s1)

	s2 := Student{
		name:   "test",
		rollNo: 100,
	}
	fmt.Printf("%+v\n", s2)

	s3 := Student{"t", 1, []int{1, 2, 3}, map[string]int{"A": 100, "B": 80}}
	fmt.Printf("%+v\n", s3)

	c := Circle{x: 5, y: 5, radius: 5}
	calcArea(&c)
	fmt.Printf("%+v\n", c)

	t1 := O1{x: 5}
	t2 := O2{x: 6}
	t3 := O2{x: 5}
	if t1 != O1(t2) {
		fmt.Println("N")
	} else {
		fmt.Println("Y")
	}
	if t1 == O1(t3) {
		fmt.Println("Y")
	}

	fmt.Printf("%+v", getMovie("xyz", 3.5))

	mov := getMovie("xyz", 2.1)
	mov1 := getMovie("abc", 3.3)
	movies := make([]Movie, 5)
	movies = append(movies, mov)
	movies = append(movies, mov1)
	for _, value := range movies {
		fmt.Println(value)
	}

	mov3 := Movie{"xyz", 2.1}
	mov4 := Movie{"abc", 2.1}
	if mov3.rating == mov4.rating || mov3 != mov4 {
		fmt.Println("condition met")
	} else if mov3.rating == mov4.rating {
		fmt.Println("condition_2 met")
	}
}
