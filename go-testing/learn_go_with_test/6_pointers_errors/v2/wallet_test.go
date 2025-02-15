package pointers_errors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	want := Bitcoin(5)

	if got != want {
		t.Errorf("got %s wnat %s", got, want)
	}
}
