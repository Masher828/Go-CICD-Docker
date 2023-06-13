package greetings

import "testing"

func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := "Manish"
    if name != want {
        t.Fatal("Test Case failedd")
    }
}
