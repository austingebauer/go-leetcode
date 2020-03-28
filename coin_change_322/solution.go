package coin_change_322

import "math"

// Version 3: Bottom-up approach using filled slice with one above amount,
// which is one above max in scenario with denomination 1 * amount.
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// Fill the array with values larger than amount
	// At most, denomination 1 coins would supply change for 'n' 'n' times
	mins := make([]int, amount+1)
	for i := range mins {
		mins[i] = amount + 1
	}

	// It takes 0 coins to make up the amount of 0
	mins[0] = 0

	for a := 1; a <= amount; a++ {
		for c := 0; c < len(coins); c++ {
			// coin is larger than current amount
			if coins[c] > a {
				continue
			}

			// mins[a] is initially large enough to after being filled into slice above
			mins[a] = int(math.Min(float64(mins[a]), float64(mins[a-coins[c]]+1)))
		}
	}

	if mins[amount] > amount {
		return -1
	}

	return mins[amount]
}

// Version 2: Bottom-up approach using MaxInt32 for min comparison
func coinChangeBottomUp1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	mins := make([]int, amount+1)
	mins[0] = 0

	for a := 1; a <= amount; a++ {
		min := math.MaxInt32
		for c := 0; c < len(coins); c++ {
			// coin is larger than current amount
			if coins[c] > a {
				continue
			}

			min = int(math.Min(float64(min), float64(mins[a-coins[c]])))
		}

		mins[a] = min + 1
	}

	if mins[amount] == math.MaxInt32+1 {
		return -1
	}

	return mins[amount]
}

// Version 1: Top-down approach using recursion
func coinChangeTopDown(coins []int, amount int) int {
	if len(coins) == 0 || amount < 1 {
		return -1
	}
	return minChange(coins, amount, make([]int, amount))
}

func minChange(coins []int, rem int, memo []int) int {
	if rem < 0 {
		return -1
	}

	if rem == 0 {
		return 0
	}

	if memo[rem-1] != 0 {
		return memo[rem-1]
	}

	min := math.MaxInt32
	for i := 0; i < len(coins); i++ {
		change := minChange(coins, rem-coins[i], memo)
		if change >= 0 {
			min = int(math.Min(float64(min), float64(change+1)))
		}
	}

	if min == math.MaxInt32 {
		return -1
	}

	memo[rem-1] = min
	return memo[rem-1]
}
