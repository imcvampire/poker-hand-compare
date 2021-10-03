package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compareHand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		firstHand   string
		secondHand  string
		expected    string
		expectedErr error
	}{
		{firstHand: "AAAQQ", secondHand: "QQAAA", expected: "Tie"},
		{firstHand: "53QQ2", secondHand: "Q53Q2", expected: "Tie"},
		{firstHand: "53888", secondHand: "88385", expected: "Tie"},
		{firstHand: "QQAAA", secondHand: "AAAQQ", expected: "Tie"},
		{firstHand: "Q53Q2", secondHand: "53QQ2", expected: "Tie"},
		{firstHand: "88385", secondHand: "53888", expected: "Tie"},
		{firstHand: "AAAQQ", secondHand: "QQQAA", expected: "Hand 1"},
		{firstHand: "Q53Q4", secondHand: "53QQ2", expected: "Hand 1"},
		{firstHand: "53888", secondHand: "88375", expected: "Hand 1"},
		{firstHand: "33337", secondHand: "QQAAA", expected: "Hand 1"},
		{firstHand: "22333", secondHand: "AAA58", expected: "Hand 1"},
		{firstHand: "33389", secondHand: "AAKK4", expected: "Hand 1"},
		{firstHand: "44223", secondHand: "AA892", expected: "Hand 1"},
		{firstHand: "22456", secondHand: "AKQJT", expected: "Hand 1"},
		{firstHand: "99977", secondHand: "77799", expected: "Hand 1"},
		{firstHand: "99922", secondHand: "88866", expected: "Hand 1"},
		{firstHand: "9922A", secondHand: "9922K", expected: "Hand 1"},
		{firstHand: "99975", secondHand: "99965", expected: "Hand 1"},
		{firstHand: "99975", secondHand: "99974", expected: "Hand 1"},
		{firstHand: "99752", secondHand: "99652", expected: "Hand 1"},
		{firstHand: "99752", secondHand: "99742", expected: "Hand 1"},
		{firstHand: "99753", secondHand: "99752", expected: "Hand 1"},
		{firstHand: "QQQAA", secondHand: "AAAQQ", expected: "Hand 2"},
		{firstHand: "53QQ2", secondHand: "Q53Q4", expected: "Hand 2"},
		{firstHand: "88375", secondHand: "53888", expected: "Hand 2"},
		{firstHand: "QQAAA", secondHand: "33337", expected: "Hand 2"},
		{firstHand: "AAA58", secondHand: "22333", expected: "Hand 2"},
		{firstHand: "AAKK4", secondHand: "33389", expected: "Hand 2"},
		{firstHand: "AA892", secondHand: "44223", expected: "Hand 2"},
		{firstHand: "AKQJT", secondHand: "22456", expected: "Hand 2"},
		{firstHand: "77799", secondHand: "99977", expected: "Hand 2"},
		{firstHand: "88866", secondHand: "99922", expected: "Hand 2"},
		{firstHand: "9922K", secondHand: "9922A", expected: "Hand 2"},
		{firstHand: "99965", secondHand: "99975", expected: "Hand 2"},
		{firstHand: "99974", secondHand: "99975", expected: "Hand 2"},
		{firstHand: "99652", secondHand: "99752", expected: "Hand 2"},
		{firstHand: "99742", secondHand: "99752", expected: "Hand 2"},
		{firstHand: "99752", secondHand: "99753", expected: "Hand 2"},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%s %s = %s", test.firstHand, test.secondHand, test.expected), func(t *testing.T) {
			actual, actualErr := compareHand(test.firstHand, test.secondHand)

			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedErr, actualErr)
		})
	}
}
