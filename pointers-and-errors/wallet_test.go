package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		got := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, got)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertInsufficientFundsError(t, err, Bitcoin(100), Bitcoin(20))
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertInsufficientFundsError(t testing.TB, err error, wantRequested, wantAvailable Bitcoin) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	insufficientErr, ok := err.(InsufficientFundsError)
	if !ok {
		t.Fatalf("got error %T %q, want InsufficientFundsError", err, err)
	}

	if insufficientErr.Requested != wantRequested {
		t.Errorf("got requested %q, want %q", insufficientErr.Requested, wantRequested)
	}

	if insufficientErr.Available != wantAvailable {
		t.Errorf("got available %q, want %q", insufficientErr.Available, wantAvailable)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
