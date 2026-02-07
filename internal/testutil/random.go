// nolint: gosec
package testutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt64 generates a random int64 between min and max.
func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomInt generates a random int between min and max.
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomHEXColor generates a random color in hexadecimal format (#RRGGBB).
func RandomHEXColor() string {
	return fmt.Sprintf("#%02X%02X%02X", rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// RandomBool generates a random boolean.
func RandomBool() bool {
	return []bool{true, false}[rand.Intn(2)]
}

// RandomDate generates a random UTC date.
func RandomDate() time.Time {
	return time.Date(
		RandomInt(1971, 2022),
		time.Month(RandomInt64(1, 12)),
		RandomInt(1, 28),
		0, 0, 0, 0,
		time.UTC,
	)
}

// RandomDate generates a random local date.
func RandomLocalDate() time.Time {
	return time.Date(
		RandomInt(1971, 2022),
		time.Month(RandomInt64(1, 12)),
		RandomInt(1, 28),
		0, 0, 0, 0,
		time.Local, // pgx decodes as local. Also must .Truncate(time.Microsecond) to compare pgx time.Time
	)
}

// RandomString generates a random string of length n.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for range n {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates a random name.
func RandomName() string {
	return RandomNameIdentifier(1, "") + " " + RandomString(int(RandomInt64(10, 15)))
}

// RandomMoney generates a random amount of money.
func RandomMoney() int64 {
	return RandomInt64(0, 1000)
}

// RandomFirstName generates a random first name.
func RandomFirstName() string {
	return firstNames[rand.Intn(len(firstNames))]
}

// RandomLastName generates a random last name.
func RandomLastName() string {
	return lastNames[rand.Intn(len(lastNames))]
}

// RandomFrom selects a random item from a list. Assumes the list is not empty.
func RandomFrom[T any](items []T) T {
	index := rand.Intn(len(items))
	return items[index]
}

// RandomEmail generates a random email.
func RandomEmail() string {
	return RandomNameIdentifier(3, ".") + "@email.com"
}

// RandomNameIdentifier generates a random name identifier,
// such as eminently-sincere-mollusk-aksticpemgicjrtb.
// Prefix count is configurable via n.
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
	ss = append(ss, RandomString(16))

	return strings.Join(ss, sep)
}
