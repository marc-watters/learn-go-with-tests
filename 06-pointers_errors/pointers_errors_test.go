package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(BitCoin(10))

		got := wallet.Balance()
		want := BitCoin(10)

		wallet.Deposit(BitCoin(10))

		if got != want {
			t.Errorf("\ngot: \t%s\nwant:\t%s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}

		wallet.Withdraw(BitCoin(10))

		got := wallet.Balance()
		want := BitCoin(10)

		if got != want {
			t.Errorf("\ngot: \t%s\nwant:\t%s", got, want)
		}
	})
}
