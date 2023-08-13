package method

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
