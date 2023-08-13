package mypackage

import (
	"fmt"
	"sync"
	"time"

	g "github.com/lexyu/go-beginner/mypackage/global"
)

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		g.PL("Not Enough Money In Account")
	} else {
		fmt.Printf("%d withdrawn : Balance : %d\n", v, a.balance)
		a.balance -= v
	}
}

func Mutex() {
	var acct Account
	acct.balance = 100
	g.PL("Balance :", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(3 * time.Second)
}