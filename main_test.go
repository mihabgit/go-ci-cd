package main

import "testing"

func TestGreet(t *testing.T) {
	expected := "Hello, Go!"
	actual := Greet("Go")
	if actual != expected {
		t.Errorf("Greet(\"Go\") = %q, want %q", actual, expected)
	}
}