package method

import "fmt"

type PhoneFeatures interface {
	Call()
}

type Phone struct {
	UserName string
}

func NewPhone(userName string) Phone {
	return Phone{UserName: userName}
}

// Fix Option 1 - constructor return pointer
// func NewPhone(userName string) *Phone {
// 	return &Phone{UserName: userName}
// }

func (p *Phone) Call() {
	fmt.Println(p.UserName, " Call...")
}

// Fix Option 2 - Value Receiver
// func (p Phone) Call() {
// 	fmt.Println(p.UserName, " Call...")
// }
