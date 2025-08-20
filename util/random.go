package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomText generates a random text of length n
func RandomText(n int, min int, max int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		sentenceLength := min + rand.Intn(max-min+1)
		for j := 0; j < sentenceLength; j++ {
			c := alphabet[rand.Intn(k)]
			sb.WriteByte(c)
		}
		const cd = '.'
		sb.WriteByte(cd)
	}

	return sb.String()
}
