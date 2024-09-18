package formulaParser

import (
	"fmt"
	"reflect"
	"testing"
)

type TestPair struct {
	input  string
	output []string
}

var testingData = []TestPair{
	TestPair{"", []string{"", "", ""}},
	TestPair{"x", []string{"", "", "x"}},
	TestPair{"1", []string{"", "1", ""}},
	TestPair{"-x", []string{"-", "", "x"}},
	TestPair{"-6x", []string{"-6", "", "x"}},
	TestPair{"6x", []string{"6", "", "x"}},
	TestPair{"+6x", []string{"+6", "", "x"}},
	TestPair{"-7", []string{"", "-7", ""}},
	TestPair{"x + 1", []string{"", "+1", "x"}},
	TestPair{"x - 1", []string{"", "-1", "x"}},
	TestPair{"6x + 1", []string{"6", "+1", "x"}},
	TestPair{"6x - 1", []string{"6", "-1", "x"}},
	TestPair{"-6x + 1", []string{"-6", "+1", "x"}},
	TestPair{"-6x - 1", []string{"-6", "-1", "x"}},
	TestPair{"xy + 1", []string{"", "+1", "xy"}},
}


type digitOpTD struct {
	input  rune
	output bool
}

var digitOperatorTD = []digitOpTD{
	digitOpTD{'0', true},
	digitOpTD{'1', true},
	digitOpTD{'2', true},
	digitOpTD{'3', true},
	digitOpTD{'4', true},
	digitOpTD{'5', true},
	digitOpTD{'6', true},
	digitOpTD{'7', true},
	digitOpTD{'8', true},
	digitOpTD{'9', true},
	digitOpTD{'+', true},
	digitOpTD{'-', true},
	digitOpTD{' ', false},
	digitOpTD{'x', false},
}


func TestSeparateBrackets(t *testing.T) {
	correct_res := []string{"x", "x + 1"}
	res := SeparateBrackets("(x)(x + 1)")

	if !reflect.DeepEqual(res, correct_res) {
		errString := fmt.Sprintf("%v != %v", res, correct_res)
		t.Error(errString)
	}

	correct_res = []string{"x + 1", "x"}
	res = SeparateBrackets("(x + 1)(x)")

	if !reflect.DeepEqual(res, correct_res) {
		errString := fmt.Sprintf("%v != %v", res, correct_res)
		t.Error(errString)
	}
}

func TestSeparatePait(t *testing.T) {
	for _, testpair := range testingData {
		output := SeparatePair(testpair.input)
		if !reflect.DeepEqual(output, testpair.output) {
			errString := fmt.Sprintf("%v != %v", output, testpair.output)
			t.Error(errString)
		}
	}
}

func TestDigitOperator(t *testing.T) {
	for _, testpair := range digitOperatorTD {
		output := digitOrOperator(testpair.input)
		if output != testpair.output {
			errString := fmt.Sprintf("%v != %v", output, testpair.output)
			t.Error(errString)
		}
	}
}
