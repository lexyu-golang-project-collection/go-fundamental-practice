package method

import "fmt"

type Worker interface {
	WorkValueReceiver()
	WorkPointerReceiver()
}

type Employee struct {
	Id   int
	Name string
	Age  int
}

func (e Employee) AddTileValueReciever(title string) {
	e.Name = title + e.Name
}

func (e *Employee) AddTilePointerReciever(title string) {
	e.Name = title + e.Name
}

func (e Employee) WorkValueReceiver() {
	fmt.Println(e.Name + " works (value receiver)")
}

func (e *Employee) WorkPointerReceiver() {
	fmt.Println(e.Name + " works  (pointer receiver)")
}
