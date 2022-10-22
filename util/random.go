package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomSKu() string {
	toReturn := []string{"character", "integer"}

	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < 10; i++ {
		returning := toReturn[rand.Intn(len(toReturn))]

		if returning == "character" {
			c := alphabet[rand.Intn(k)]
			sb.WriteByte(c)
		} else {
			sb.WriteString(strconv.Itoa(RandomInt(0, 10)))
		}
	}

	return sb.String()

}
func RandomCountry() string {
	countries := []string{"ke", "ci", "eg", "ma"}

	return countries[rand.Intn(len(countries))]
}
func RandomProductName() string {
	return RandomString(5) + " " + RandomString(6)
}
