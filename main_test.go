package main

import "testing"

func Test_isPrime(t *testing.T) {
	param := 0
	result, msg := isPrime(param)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", param)
	}

	if msg != "0 is not prime, by definition" {
		t.Errorf("wrong message returned")
	}

	param = 7
	result, msg = isPrime(param)
	if !result {
		t.Errorf("with %d as test parameter, got false, but expected true", param)
	}

	if msg != "7 is a prime number!" {
		t.Errorf("wrong message returned")
	}
}
