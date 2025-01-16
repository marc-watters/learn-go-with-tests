package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("\ngot: \t%d\nwant:\t%d", got, want)
	}
}
