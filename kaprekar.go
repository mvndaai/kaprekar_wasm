package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"syscall/js"
)

var kaprekar_constant = map[int]int{4: 6174, 3: 495}
var kaprekar_max_iterations = map[int]int{4: 7, 3: 6}

func main() {

	jsFuncName := "kaprekar"
	js.Global().Set(jsFuncName, kaprekarWrapper())
	fmt.Printf("Function '%s' loaded from Go\n", jsFuncName)

	// Keep the program alive so functions can be run over and over
	<-make(chan bool)
}

type output struct {
	Error interface{} `json:"error"`
	Parts [][]int     `json:"parts"`
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

func kaprekarWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		input := args[0].String()
		parts, err := kaprekar(input)
		out := output{Parts: parts}
		if err != nil {
			out.Error = err.Error()
		}

		pretty, err := json.Marshal(out)
		if err != nil {
			return "Could not marshal json"
		}
		return string(pretty)
	})
}
