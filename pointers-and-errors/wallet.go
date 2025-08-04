package pointersanderrors

import (
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type InsufficientFundsError struct {
	Requested Bitcoin
	Available Bitcoin
}

func (e InsufficientFundsError) Error() string {
	return fmt.Sprintf("cannot withdraw %s, only %s available", e.Requested, e.Available)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError{
			Requested: amount,
			Available: w.balance,
		}
	}
	w.balance -= amount
	return nil
}
