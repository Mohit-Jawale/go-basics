package pointersanderrors

import "fmt"

type bitcoin int

type Wallet struct {
	balance bitcoin
}

func (w *Wallet) Deposit(amount bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() bitcoin {
	return w.balance
}
