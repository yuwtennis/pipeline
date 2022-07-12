package test

import "testing"

func TestHelloName(t *testing.T) {
	name := "John"
	expect := "John"

	if name != expect {
		t.Fatal("Name not correct")
	}
}
