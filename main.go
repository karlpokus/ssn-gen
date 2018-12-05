/*
	version 0.3?
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

func partialBuild(n int) <-chan []int {
	out := make(chan []int)

	go func() {
		for i := 0; i < n; i++ {
			y, yy, _ := randInts(0, 100)
			m, mm, _ := randInts(1, 13)
			d, dd, _ := randInts(1, 29)
			x, xx, xxx := randInts(111, 1000)

			out <- []int{y, yy, m, mm, d, dd, x, xx, xxx}
		}
		close(out)
	}()

	return out
}

func completeBuild(in <-chan []int) <-chan []int {
	out := make(chan []int)

	go func() {
		for ints := range in {
			var sum int

			for i, v := range ints {
				v = v * (2 - i % 2)

				if v >= 10 {
					v -= 9
				}
				sum += v
			}

			lastDigit := (100 - sum) % 10

			out <- append(ints, lastDigit)
		}
		close(out)
	}()

	return out
}

func sliceToString(in <-chan []int) <-chan string {
	out := make(chan string)

	go func() {
		for ints := range in {
			temp := make([]string, len(ints))

			for i, v := range ints {
		  	temp[i] = strconv.Itoa(v)
		  }

		  out <- strings.Join(temp, "")
		}
		close(out)
	}()

	return out
}

func ssnGen(n int) string {
	var out []string

	for s := range sliceToString(completeBuild(partialBuild(n))) {
		out = append(out, s)
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
