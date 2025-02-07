package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(bitcoin(10))

	got := wallet.Balance()

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	want := bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
