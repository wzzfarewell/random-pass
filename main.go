package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 12, "length of random pass")
	flag.Parse()

	pass := RandomString(n)
	clipboard.WriteAll(pass)
}

// RandomString generates a random string of length n.
// The string is composed of letters, numbers and symbols. Symbols are chosen from: _-.*?
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-.*?")
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(100))))
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[r.Intn(len(letters))]
	}
	return string(s)
}
