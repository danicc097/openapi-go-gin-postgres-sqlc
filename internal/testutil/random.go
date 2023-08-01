package testutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// alternative r = rand.New(rand.NewSource(time.Now().UnixNano()))
// is not concurrency safe. Would need something like:
// https://github.com/cilium/cilium/blob/32118ccaa8677a3b836e60542a52a2af94693533/pkg/rand/safe_rand.go#L23
// we don't really care since these are exclusively used for tests
// nolint: gochecknoinits
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt64 generates a random int64 between min and max.
func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomInt generates a random int between min and max.
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
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

	for i := 0; i < n; i++ {
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

// RandomEmail generates a random email.
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomNameIdentifier(3, "."))
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
