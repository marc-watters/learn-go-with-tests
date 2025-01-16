package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		amount := BitCoin(10)
		wallet.Deposit(amount)
		assertBalance(t, wallet, amount)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		err := wallet.Withdraw(BitCoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		balance := BitCoin(20)
		wallet := Wallet{balance}
		err := wallet.Withdraw(200)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, balance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want BitCoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("\ngot: \t%s\nwant:\t%s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("did not an expect an error")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected an error")
	}

	if got != want {
		t.Errorf("\ngot: \t%s\nwant:\t%s", got, want)
	}
}
