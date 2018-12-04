/*
	version 0.1
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
	"flag"
)

// >= min && < max
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

func partialBuild() []int {
	y, yy, _ := randInts(0, 100)
	m, mm, _ := randInts(1, 13)
	d, dd, _ := randInts(1, 29)
	x, xx, xxx := randInts(111, 1000)

	return []int{y, yy, m, mm, d, dd, x, xx, xxx}
}

func completeBuild(ints []int) []int {
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

func sliceToString(a []int) string {
  b := make([]string, len(a))

	for i, v := range a {
  	b[i] = strconv.Itoa(v)
  }

  return strings.Join(b, "")
}

func ssnGen(n int) string {
	var out []string

	for i := 0; i < n; i++ {
		out = append(out, sliceToString(completeBuild(partialBuild())))
	}
	return strings.Join(out, "\n")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var ssns = flag.Int("ssns", 3, "number of ssns to create")
	flag.Parse()

	data := ssnGen(*ssns);

	fmt.Print(data)
}
