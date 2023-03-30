package main

import "testing"

type IsPrimeTestTableSetup struct {
	name     string
	testNum  int
	expected bool
	msg      string
}

func Test_isPrime(t *testing.T) {

	primeTests := []IsPrimeTestTableSetup{
		{name: "prime", testNum: 7, expected: true, msg: "7 is a prime number!"},
		{name: "not prime", testNum: 8, expected: false, msg: "8 is not a prime number because it is divisible by 2!"},
	}

	for _, entry := range primeTests {
		result, msg := isPrime(entry.testNum)
		if entry.expected && !result {
			t.Errorf("%s: expected true, but got false", entry.name)
		}

		if !entry.expected && result {
			t.Errorf("%s: expected false, but got true", entry.name)
		}

		if entry.msg != msg {
			t.Errorf("%s: expected %s but got %s", entry.name, entry.msg, msg)
		}
	}
}
