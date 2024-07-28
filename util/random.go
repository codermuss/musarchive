package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

func RandomUsername() string {
	return RandomString(12)
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

func RandomLike() int32 {
	return int32(RandomInt(0, 1000))
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func DateNow() time.Time {
	now := time.Now().UTC()

	// Extract the year, month, and day
	year, month, day := now.Date()

	// Create a new time.Time object with only the date portion
	dateOnly := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return dateOnly
}

func DateYesterday() time.Time {
	now := time.Now().Add(24 * time.Hour).UTC()

	// Extract the year, month, and day
	year, month, day := now.Date()

	// Create a new time.Time object with only the date portion
	dateOnly := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return dateOnly
}

func DateFixed() time.Time {
	return time.Date(2000, time.May, 6, 0, 0, 0, 0, time.UTC)
}

func RandomOwner() string {
	return RandomString(6)
}
