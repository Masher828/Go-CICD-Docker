package greetings

import "testing"

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := "Gladys"
	if name != want {
		t.Fatal("Test Case failedd")
	}
}

func TestBye(t *testing.T) {
	name := "Manish"
	want := "Mani"
	if name != want {
		t.Fatal("Test Case failedd")
	}
}
