package main

import "fmt"

type Data struct {
	value int
}

func (d Data) print() {
	fmt.Printf("%+v\n", d)
}

func (d *Data) print2() {
	fmt.Printf("%+v\n", *d)
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) incrementAge() {
	p.Age++
}

type data []string

func (d *data) Add(item string) {
	*d = append(*d, item)
}

// built-in types (string, bool, int...etc, prefer using value receiver)
// recevier(map, function, channel type are already pointer types)
// recevier(slice) use a pointer
func main() {
	// Demo1()
	// Demo2()
	// Demo3()
	Demo4()
}

func Demo1() {
	data := Data{value: 1}
	f := data.print
	data.value = 20
	f()
	fmt.Printf("%+v\n", data)
	f()
}

func Demo2() {
	data := Data{value: 1}
	f := data.print2
	data.value = 20
	f()
	fmt.Printf("%+v\n", data)
	f()
}

func Demo3() {
	users := []Person{
		{"J", 30},
		{"K", 25},
	}
	for _, user := range users {
		user.incrementAge()
	}
	fmt.Printf("%+v\n", users)
}

func Demo4() {
	users := []Person{
		{"J", 30},
		{"K", 25},
	}
	for i := range users {
		users[i].incrementAge()
	}
	fmt.Printf("%+v\n", users)
}
