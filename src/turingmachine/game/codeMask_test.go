package game_test

import (
	"slices"
	"testing"

	"stefanovazzoler.com/turingmachine/src/turingmachine/game"
)

func TestCodeMask(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		law      *game.Law
		code     game.Code
		expected bool
	}{
		{game.Criterias[0].Laws[0], game.CodeFromNumbers(1, 1, 1), true},  // △ = 1
		{game.Criterias[1].Laws[0], game.CodeFromNumbers(5, 1, 1), false}, // △ < 3
		{game.Criterias[9].Laws[1], game.CodeFromNumbers(5, 4, 5), true},  // one 4
	}

	for i, testCase := range testCases {
		mask := testCase.law.Mask

		codes := mask.GetAllCodes()
		if slices.Contains(codes, testCase.code) != testCase.expected {
			t.Errorf("[%d] Law(%d).Mask.GetAllCodes() contains code(%s) expected to be %v but got opposite",
				i, testCase.law.Id, testCase.code.String(), testCase.expected)
		}

		for _, code := range codes {
			if !mask.Check(code) {
				t.Errorf("[%d] Law(%d).Mask.GetAllCodes() included code(%s) but doesn't pass the Check()",
					i, testCase.law.Id, code)
			}
			valid, ok := game.CheckCode(testCase.law.Id, code)
			if !valid || !ok {
				t.Errorf("[%d] Law(%d).Mask.GetAllCodes() included code(%s) but CheckCode() returned (%v, %v)",
					i, testCase.law.Id, code, valid, ok)
			}
		}
	}
}
