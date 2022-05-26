package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var kaprekar_constant = map[int]int{4: 6174, 3: 495}
var kaprekar_max_iterations = map[int]int{4: 7, 3: 6}

func main() {

	for _, v := range []string{"3215", "450", "1111", "21"} {
		parts, err := kaprekar(v)
		out := output{Error: err, Parts: parts}
		fmt.Println(v, out)
	}

}

type output struct {
	Error error   `json:"error"`
	Parts [][]int `json:"parts"`
}

func kaprekar(in string) ([][]int, error) {
	if err := checkDigitsDontMatch(in); err != nil {
		return nil, err
	}

	l := len(in)
	if _, ok := kaprekar_max_iterations[l]; !ok {
		return nil, fmt.Errorf("constant for %d digit number unknown", l)
	}

	num, err := strconv.Atoi(in)
	if err != nil {
		return nil, fmt.Errorf("input was not a number")
	}
	parts := [][]int{}
	max := kaprekar_max_iterations[l]
	k := kaprekar_constant[l]
	for {
		if num == k {
			return parts, nil
		}

		if len(parts) > max {
			return nil, fmt.Errorf("could not get to constant within %d iterations", max)
		}

		small, big := sortUsingStringSlices(num, l)
		num = big - small

		parts = append(parts, []int{big, small, num})
	}
}

func checkDigitsDontMatch(in string) error {
	split := strings.Split(in, "")

	m := map[string]bool{}
	for _, v := range split {
		m[v] = true
	}

	if len(m) < 2 {
		return fmt.Errorf("all digits must not repeat")
	}

	return nil
}

func sortUsingStringSlices(in int, digits int) (int, int) {
	padded := fmt.Sprintf("%0*d", digits, in)
	s := strings.Split(padded, "")

	sort.Sort(sort.StringSlice(s))
	small, _ := strconv.Atoi(strings.Join(s, ""))

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	big, _ := strconv.Atoi(strings.Join(s, ""))

	return small, big
}
