package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

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
		{name: "zero", testNum: 0, expected: false, msg: "0 is not prime, by definition!"},
		{name: "one", testNum: 1, expected: false, msg: "1 is not prime, by definition!"},
		{name: "negative number", testNum: -20, expected: false, msg: "Negative numbers are not prime, by definition!"},
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

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	read, write, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = write

	prompt()

	// close our writer
	_ = write.Close()

	// reset os.Stdout to what is was before
	os.Stdout = oldOut

	// read the output of our propt() func from our read pipe
	out, _ := io.ReadAll(read)

	if string(out) != "-> " {
		t.Errorf("incorrect propt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	read, write, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = write

	intro()

	// close our writer
	_ = write.Close()

	// reset os.Stdout to what is was before
	os.Stdout = oldOut

	// read the output of our propt() func from our read pipe
	out, _ := io.ReadAll(read)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}
