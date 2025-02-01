package game_test

import (
	"math"
	"slices"
	"testing"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
)

const (
	firstCriteria  = game.BlankChoice + 1
	secondCriteria = firstCriteria + 2
)

func TestNextLaw(t *testing.T) {
	t.Parallel()

	for i := range game.MaxChoice + 1 {
		next := game.Choice(i).NextLaw()
		if i >= game.MaxChoice {
			if next != game.BlankChoice {
				t.Errorf("MaxChoice should have been BlankChoice, instead got %d",
					next)
			}
			continue
		}
		if next != game.Choice(i+1) {
			t.Errorf("Expected next law for %d to be %d, got %d",
				i, i+1, next)
		}
	}
}

func TestChoices(t *testing.T) {
	t.Parallel()

	var choice game.Choice
	for choice = range game.Choice(game.MaxChoice + 1) {
		// Special case: check if blank choice
		if choice == game.BlankChoice {
			if !choice.IsValid() {
				t.Fatal("BlankChoice incorrectly marked as invalid")
			}
			if choice.Criteria() != nil || choice.Law() != nil {
				t.Fatal("BlankChoice has non-nil criteria or law")
			}
			if choice.NextCriteria() != game.Choice(1) {
				t.Fatalf("BlankChoice reported incorrect NextCriteria: %d",
					choice.NextCriteria())
			}
			if choice.Mask() != game.BaseMask {
				t.Fatalf("BlankChoice reported incorrect Mask: %v",
					choice.Mask())
			}
			continue
		}

		if !choice.IsValid() {
			t.Fatalf("Choice(%d) incorrectly marked as invalid", choice)
		}
		criteria := choice.Criteria()
		if criteria == nil {
			t.Fatalf("Choice(%d) has nil criteria", choice)
		}
		law := choice.Law()
		if law == nil {
			t.Fatalf("Choice(%d) has nil law", choice)
		}
		if !slices.ContainsFunc(criteria.Laws, func(l *game.Law) bool {
			return l == law
		}) {
			t.Fatalf("Choice(%d) has invalid match of criteria %d + law %d",
				choice, criteria.Id, law.Id)
		}
		if choice.Mask() != choice.Law().Mask {
			t.Fatalf("Choice(%d) mask is %v but law has mask %v",
				choice, choice.Mask(), choice.Law().Mask)
		}
		if criteria.Difficulty() != choice.Difficulty() {
			t.Fatalf("Choice(%d) has difficulty %d but its criteria %d has difficulty %d",
				choice, choice.Difficulty(), criteria.Id, criteria.Difficulty())
		}

		found := false
		for _, law := range choice.Criteria().Laws {
			if choice.Law() == law {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Choice(%d) with criteria %d has orphaned law %d",
				choice, choice.Criteria().Id, choice.Law().Id)
		}

		next := choice.NextCriteria()
		if choice.Criteria().Id == game.NumberOfCriterias {
			if next != game.BlankChoice {
				t.Fatalf("Choice(%d).NextCriteria() returned %d instead of Blank",
					choice, next)
			}
		} else {
			if !next.IsValid() {
				t.Fatalf("Choice(%d).NextCriteria() returned the invalid choice %d",
					choice, next)
			}
			if next.Criteria().Id != choice.Criteria().Id+1 {
				t.Fatalf("Choice(%d).NextCriteria() returned Choice(%d) with criteria %d instead of expected %d",
					choice, next, next.Criteria().Id, choice.Criteria().Id+1)
			}
		}
	}
	// Overflow check
	for choice = game.Choice(game.MaxChoice + 1); choice <= math.MaxInt8; choice++ {
		if choice.IsValid() {
			t.Fatalf("[overflow] Choice(%d) incorrectly marked as valid", choice)
		}
		if choice.Criteria() != nil || choice.Law() != nil {
			t.Fatalf("[overflow] Choice(%d) has non-nil criteria or law", choice)
		}
		if choice.NextCriteria() != game.Choice(1) {
			t.Fatalf("[overflow] Choice(%d) reported incorrect NextCriteria: %d",
				choice, choice.NextCriteria())
		}
		if choice.Mask() != game.BaseMask {
			t.Fatalf("[overflow] Choice(%d) reported incorrect Mask: %v",
				choice, choice.Mask())
		}
		continue
	}
}

func TestNextCriteria(t *testing.T) {
	t.Parallel()

	testCases := map[game.Choice]game.Choice{
		game.BlankChoice:  firstCriteria,
		firstCriteria:     secondCriteria,
		firstCriteria + 1: secondCriteria,
	}
	for test, expected := range testCases {
		actual := test.NextCriteria()
		if actual != expected {
			t.Errorf("Expected %d.NextCriteria()==%d, instead got %d",
				test, expected, actual)
		}
	}
}
