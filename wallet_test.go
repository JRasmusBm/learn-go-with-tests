package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, got, want Bitcoin) {
		t.Helper()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Errorf("Wanted an error but received none")
		}

		if got != want {
			t.Errorf("Error message: got %#v want %#v", got, want)
		}
	}

	t.Run("Can get balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		assertBalance(t, wallet.Balance(), Bitcoin(10))

	})

	t.Run("Correctly formats balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		got := wallet.Balance().String()
		want := "10 BTC"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Can deposit funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		wallet.Deposit(10)

		assertBalance(t, wallet.Balance(), Bitcoin(20))
	})

	t.Run("Withdraw when balance is sufficient", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(2)

		t.Run("Subtracts the amount from the old balance", func(t *testing.T) {
			assertBalance(t, wallet.Balance(), Bitcoin(8))
		})

		t.Run("Does not return an error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error got %#v", err)
			}
		})
	})

	t.Run("Withdraw when balance is insufficient", func(t *testing.T) {
		initialBalance := Bitcoin(10)

		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(12)

		t.Run("Balance stays the same", func(t *testing.T) {
			assertBalance(t, wallet.Balance(), initialBalance)
		})

		t.Run("An error is returned", func(t *testing.T) {
			assertError(t, err, InsufficientFundsError)
		})
	})
}
