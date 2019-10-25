package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("tests Messages func", func(t *testing.T) {
		expected := "hi there, i love goo!"
		if Message("go") != expected {
			t.Errorf("error: expected %s\n got %s", expected, Message("go"))
		}
	})
}
