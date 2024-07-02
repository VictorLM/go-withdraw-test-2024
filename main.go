package main

import (
	"errors"
	"fmt"
)

// available banknotes
var availableBanknotes = []int{200, 100, 50, 20, 10, 5}

// map to store banknotes and amounts
var banknoteAmounts = map[int]int{}

// global variable to store rest amount
var rest int

func main() {
	// show available banknotes to user
	fmt.Println("Cédulas disponíveis: ")
	for _, banknotes := range availableBanknotes {
		fmt.Println(banknotes)
	}
	// ask user for amount to withdraw
	var amountToWithdraw int
	fmt.Print("Entre com o valor a ser sacado: ")
	fmt.Scan(&amountToWithdraw)
	// check if amount is valid
	if amountToWithdraw <= 0 {
		fmt.Println("Valor inválido!")
		return
	}
	// set rest to amount to global variable
	rest = amountToWithdraw
	// loop until rest is 0
	err := loop()
	// check if there was an error
	if err != nil {
		fmt.Println(err)
		return
	}
	// print success message
	fmt.Println("Saque realizado com sucesso!")
	for banknote, amount := range banknoteAmounts {
		fmt.Printf("%d nota(s) de %d \n", amount, banknote)
	}
	fmt.Println("Muito obrigado!")
}

func loop() error {
	// looped variable to check if enter condition
	var looped bool = false
	// loop through available banknotes
	for _, banknote := range availableBanknotes {
		// check if rest value is greater than banknote
		if rest >= banknote {
			// subtract banknote from rest
			rest = (rest - banknote)
			// add banknote to banknoteAmounts map
			banknoteAmounts[banknote] = banknoteAmounts[banknote] + 1
			// set looped variable to true
			looped = true
			// break loop
			break
		}
	}
	// check if rest is 0 or looped
	if rest == 0 {
		return nil
	} else if rest > 0 && looped {
		return loop() // recursion
	}
	// if rest greater than zero after looped and not enter condition, return error
	return errors.New("valor indisponível para saque")
}
