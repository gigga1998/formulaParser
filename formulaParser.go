package formulaParser

import (
	"strings"
	"unicode"
)

func SeparateBrackets(formula string) []string {
	/*
		(ax + b)(cx + d) -> acx2 + (ad + bc)x + bd
	*/
	var output = []string{}
	var statement = ""

	for _, ch := range formula {
		if ch == '(' {
			statement = ""
		} else if ch == ')' {
			output = append(output, statement)
		} else {
			statement += string(ch)
		}

	}
	return output
}

func SeparatePair(pair string) []string {
	/*
		"ax + b" -> {a, b, x}
	*/
	trimString := strings.ReplaceAll(pair, " ", "")
	firstCoef := ""
	secondCoef := ""
	variable := ""
	i := 0

	for _, ch := range trimString {
		if digitOrOperator(ch) {
			firstCoef += string(ch)
			i++
		} else {
			break
		}
	}

	trimString = trimString[i:]
	i = 0
	for _, ch := range trimString {
		if !digitOrOperator(ch) {
			variable += string(ch)
			i++
		} else {
			break
		}
	}

	trimString = trimString[i:]
	for _, ch := range trimString {
		if digitOrOperator(ch) {
			secondCoef += string(ch)
			i++
		} else {
			break
		}
	}

	if variable == "" {
		return []string{"", firstCoef, ""}
	}
	return []string{firstCoef, secondCoef, variable}
}

func digitOrOperator(ch rune) bool {
	if unicode.IsDigit(ch) || ch == '-' || ch == '+' {
		return true
	}
	return false
}
