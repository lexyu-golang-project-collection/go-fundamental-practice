package pointers_errors

import "fmt"

type Wallet struct {
	balance uint
}

func (w *Wallet) Deposit(amount uint) {
	w.balance += amount
}

func (w *Wallet) Balance() uint {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	return w.balance
}
