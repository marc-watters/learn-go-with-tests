package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want BitCoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("\ngot: \t%s\nwant:\t%s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()

		if err == nil {
			t.Error("expected an error")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		amount := BitCoin(10)
		wallet.Deposit(amount)
		assertBalance(t, wallet, amount)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}

		err := wallet.Withdraw(BitCoin(10))
		if err != nil {
			t.Errorf("error withdrawing from wallet: %s", err)
		}

		want := BitCoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		balance := BitCoin(20)
		wallet := Wallet{balance}
		err := wallet.Withdraw(200)

		assertBalance(t, wallet, balance)
		assertError(t, err)
	})
}
