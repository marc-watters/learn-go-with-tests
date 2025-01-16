package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(BitCoin(10))

	got := wallet.Balance()
	want := BitCoin(10)

	if got != want {
		t.Errorf("\ngot: \t%d\nwant:\t%d", got, want)
	}
}
