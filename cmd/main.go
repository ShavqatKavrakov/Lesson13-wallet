package main

import (
	"fmt"

	"github.com/ShavqatKavrakov/wallet/pkg/wallet"
)

func main() {
	scv := &wallet.Service{}
	_, err := scv.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
	}
	//scv.Deposit(1, 10)
	//scv.RegisterAccount("+992000000002")
}
