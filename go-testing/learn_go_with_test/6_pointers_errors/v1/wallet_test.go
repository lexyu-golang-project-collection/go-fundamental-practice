package pointers_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	result := wallet.Balance()

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	var expected uint = 10

	if result != expected {
		t.Errorf("result %d expected %d", result, expected)
	}
}
