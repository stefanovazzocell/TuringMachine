package game_test

import (
	"strconv"
	"testing"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
)

func TestCode(t *testing.T) {
	t.Parallel()

	code := game.MinCode
	counter := 0
	for {
		counter++

		codeStr := code.String()
		codeFromStr, err := game.CodeFromString(codeStr)
		if err != nil || codeFromStr != code {
			t.Errorf("game.CodeFromString(%q) returned (%d, %v), but expected %d",
				codeStr, codeFromStr, err, code)
		}

		codeInt := code.Int()
		triangle, square, circle := uint8(codeInt/100%10), uint8(codeInt/10%10), uint8(codeInt%10)
		codeFromInt := game.CodeFromNumbers(triangle, square, circle)
		if err != nil || codeFromInt != code {
			t.Errorf("game.CodeFromNumbers(%d, %d, %d) returned (%d), but expected %d",
				triangle, square, circle, codeFromInt, code)
		}

		codeIntStr := strconv.Itoa(codeInt)
		if codeStr != codeIntStr {
			t.Errorf("expected codeStr (%q) and codeIntStr (%q) to be equal to each other",
				codeStr, codeIntStr)
		}

		codeIdx := code.GetIndex()
		if int(codeIdx) != (counter - 1) {
			t.Errorf("%d.GetIndex() = %d but expected %d",
				code, codeIdx, counter-1)
		}

		if code == game.MaxCode {
			code = code.Incr()
			if code != game.MaxCode {
				t.Errorf("expected Incr() to return the same value once max is reached, instead got %d",
					code)
			}
			break
		}
		code = code.Incr()
	}
	// There should be 125 codes
	if counter != 125 {
		t.Errorf("expected 125 codes, instead got %d", counter)
	}
}
