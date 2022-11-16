package testutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// nolint: gochecknoinits
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates a random name.
func RandomName() string {
	return RandomNameIdentifier(1, "") + RandomString(int(RandomInt(10, 25)))
}

// RandomMoney generates a random amount of money.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomFirstName generates a random first name.
func RandomFirstName() string {
	return firstNames[rand.Intn(len(firstNames))]
}

// RandomLastName generates a random last name.
func RandomLastName() string {
	return lastNames[rand.Intn(len(lastNames))]
}

// RandomEmail generates a random email.
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomNameIdentifier(3, ".")+RandomString(int(RandomInt(5, 10))))
}

// RandomNameIdentifier generates a random name identifier,
// such as eminently-sincere-mollusk.
func RandomNameIdentifier(n int, sep string) string {
	adv := adverbs[rand.Intn(len(adverbs))]
	adj := adjectives[rand.Intn(len(adjectives))]
	nam := names[rand.Intn(len(names))]

	var ss []string
	switch n {
	case 1:
		ss = append(ss, nam)
	case 2:
		ss = append(ss, adj, nam)
	default:
		ss = append(ss, adv, adj, nam)
	}
	return strings.Join(ss, sep)
}
