package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("- Can get balance", func(t *testing.T) {
		wallet := Wallet{balance: 10}

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("- Can deposit funds", func(t *testing.T) {
		wallet := Wallet{balance: 10}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(20)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
