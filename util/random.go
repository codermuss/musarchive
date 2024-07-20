package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnoprstuvwxyz"

func init() {
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomTitle() string {
	return RandomString(10)
}

func RandomDescription() string {
	return RandomString(255)
}

func RandomImage() string {
	return "https://picsum.photos/200/300"
}

// RandomMoney generates a random amount of money
func RandomMoney() float64 {
	return float64(RandomInt(0, 1000))
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
