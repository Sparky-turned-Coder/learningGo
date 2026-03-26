package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// Don't edit above this line

	if finalCost = bulkMessageCost; isPremiumUser {
		finalCost = finalCost - finalCost*discountRate
		if finalCost <= accountBalance {
			fmt.Println(purchaseSuccessMessage)
		} else {
			fmt.Println(insufficientFundMessage)
		}
	}
	accountBalance = accountBalance - finalCost

	// don't edit below this line

	fmt.Println("Account balance: ", accountBalance)
}

// Steps we performed on this exercise:  (We wrote the logic for our if statements)

// 1. Set finalCost to the bulkMessageCost.
// 2. If the user is a premium user, apply the discountRate to the finalCost.
//  	- For example, a discountRate of 0.10 means the discounted price per message would be 90% of the original price.
// 3. If the user has enough money in their accountBalance:
//		Deduct finalCost from their accountBalance.
// 		Print the purchaseSuccessMessage
// 4. If not, just print the insufficientFundMessage.
