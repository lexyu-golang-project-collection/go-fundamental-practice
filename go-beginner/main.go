package main

import (
	"fmt"

	method "github.com/lexyu/go-beginner/mypackage/method"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	// myfunc.Empty()
	// myfunc.Hello()
	// fmt.Println("---------------------------")
	// myfunc.Casting()
	// fmt.Println("---------------------------")
	// myfunc.Condtional()
	// fmt.Println("---------------------------")
	// myfunc.Runes()
	// fmt.Println("---------------------------")
	// myfunc.Time()
	// fmt.Println("---------------------------")
	// myfunc.Math()
	// fmt.Println("---------------------------")
	// myfunc.Range()
	// fmt.Println("---------------------------")
	// myfunc.Array()
	// fmt.Println("---------------------------")
	// myfunc.Slice()
	// fmt.Println("---------------------------")
	// myfunc.Funcs()
	// fmt.Println("---------------------------")
	// myfunc.Pointers()
	// fmt.Println("---------------------------")
	// myfunc.FileIO()
	// fmt.Println("---------------------------")
	// myfunc.CLI()
	// fmt.Println("---------------------------")
	// myfunc.Maps()
	// fmt.Println("---------------------------")
	// myfunc.Generics()
	// fmt.Println("---------------------------")
	// myfunc.Composition()
	// fmt.Println("---------------------------")
	// myfunc2.OOP()
	// oop.Empty()
	// fmt.Println("---------------------------")
	// myfunc.Defined_Types_Associate_Methods()
	// fmt.Println("---------------------------")
	// myfunc.Interfaces()
	// fmt.Println("---------------------------")
	// concurrency.Goroutines()
	// concurrency.Empty()
	// fmt.Println("---------------------------")
	// myfunc.Channels()
	// fmt.Println("---------------------------")
	// myfunc.Mutex()
	// fmt.Println("---------------------------")
	// myfunc.Closures()
	fmt.Println("---------------------------")
	// myfunc.Recursion(10)
	// fmt.Println("---------------------------")
	// regex.Regex()
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	// myfunc
	// fmt.Println("---------------------------")
	employee1 := method.Employee{1, "L", 1000}
	employee1.AddTileValueReciever("Mr.")
	fmt.Println("Value Reciever Employee Name ", employee1.Name)
	fmt.Println("---------------------------")
	employee2 := &method.Employee{2, "LL", 100}
	employee2.AddTilePointerReciever("Mr.")
	fmt.Println("Pointer Reciever Employee Name ", employee2.Name)
}
