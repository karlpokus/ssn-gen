package ssn

import (
	"math/rand"
	"time"
	"strconv"
	"strings"
)

// randInts generates a random number (>= min && < max) and returns them one by one
func randInts(min, max int) (int, int, int) {
	num := min + rand.Intn(max - min)
	if num > 99 {
		return num / 100 % 10, num / 10 % 10, num % 10
	}
	if num < 10 {
		return 0, num, 0
	}
	return num / 10 % 10, num % 10, 0
}

// partial returns a partial personnummer - omitting the last digit - in a slice
func partial() []int {
	y, yy, _ := randInts(0, 100)
	m, mm, _ := randInts(1, 13)
	d, dd, _ := randInts(1, 29)
	//x, xx, xxx := randInts(111, 1000)
	x, xx, xxx := randInts(980, 1000) // add safe range

	return []int{y, yy, m, mm, d, dd, x, xx, xxx}
}

// complete computes the last digit of a partial and returns a complete personnummer
func complete(ints []int) []int {
	var sum int
	for i, v := range ints {
		v = v * (2 - i % 2)
		if v >= 10 {
			v -= 9
		}
		sum += v
	}
	lastDigit := (100 - sum) % 10
	return append(ints, lastDigit)
}

// toString takes a slice of ints and returns a personnummer string in the
// format yymmdd-xxxx
func toString(a []int) string {
  b := make([]string, len(a)+1) // +1 for the separator
	for i, v := range a {
		if i == 6 {
			b = append(b, "-")
		}
		b = append(b, strconv.Itoa(v))
  }
  return strings.Join(b, "")
}

// GenN returns n safe ssns separated by line breaks
func GenN(n int) string {
	rand.Seed(time.Now().UnixNano())
	var out []string

	for i := 0; i < n; i++ {
		out = append(out, toString(complete(partial())))
	}
	return strings.Join(out, "\n")
}
