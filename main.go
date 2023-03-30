package main

import "fmt"

func main() {
	n := 0

	_, msg := isPrime(n)

	fmt.Println(msg)
}

func isPrime(number int) (bool, string) {
	if number == 0 || number == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", number)
	}

	if number < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", number, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", number)
}
