package main

import (
	"testing"

	"github.com/atotto/clipboard"
)

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(12)
	}
}

func TestRandomString(t *testing.T) {
	s := RandomString(12)
	t.Log(s)
}

func TestClipboard(t *testing.T) {
	s, err := clipboard.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
