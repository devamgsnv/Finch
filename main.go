package main

import "fmt"

func main() {
	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	balance := calculateBalance(transactions)
	fmt.Printf("Your balance: %.2f", balance)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Enter transaction (n for exit): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
