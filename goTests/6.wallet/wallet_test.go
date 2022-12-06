package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		assert.Equal(t, got, want)
	}

	t.Run("deposit", func(t *testing.T) {
		// Arrange
		want := Bitcoin(10)
		wallet := Wallet{}

		// Act
		wallet.Deposit(want)

		// Assert
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		// Arrange
		want := Bitcoin(10)
		wallet := Wallet{balance: Bitcoin(20)}

		// Act
		err := wallet.Withdraw(Bitcoin(10))

		// Assert
		assertBalance(t, wallet, want)
		assert.NoError(t, err)
	})

	t.Run("withdraw", func(t *testing.T) {
		// Arrange
		want := Bitcoin(10)
		wallet := Wallet{balance: want}

		// Act
		err := wallet.Withdraw(Bitcoin(100))

		// Assert
		assertBalance(t, wallet, want)
		assert.ErrorIs(t, err, ErrInsufficientFunds)
	})
}
