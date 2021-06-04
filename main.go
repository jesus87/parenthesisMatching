package main

import (
	"fmt"
)

func main() {
	codes := []string{
		//First Test
		`“{}”
	“{foo()}”
	{foo(bar[]) baz}”
	“{
	(blahblahblah)[][]
	Lorem{
	ipsum()
	}
	}”`,

		//Second Test
		`“{“
		“{[}]”
		“{foo(}”
		“foo[bar()]}”`}

	for _, code := range codes {
		fmt.Println(matchParenthesis(code))
	}
}

func matchParenthesis(code string) bool {
	opened := []rune{}
	for _, c := range code {
		if validate, k, _ := getValidation(c); validate {
			if len(opened) == 0 {
				opened = append(opened, k)
				continue
			}

			if lastMatch := opened[len(opened)-1]; k == lastMatch {
				opened = opened[:len(opened)-1]
				continue
			}

			opened = append(opened, k)
		}
	}

	return len(opened) == 0
}

func getValidation(r rune) (bool, rune, rune) {
	validator := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for k, v := range validator {
		if v == r || k == r {
			return true, k, r
		}
	}

	return false, '-', '-'
}
