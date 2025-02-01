package game_test

import (
	"slices"
	"testing"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
)

func TestCriterias(t *testing.T) {
	t.Parallel()

	for i, criteria := range game.Criterias {
		if criteria == nil {
			t.Errorf("Criteria %d is nil", i)
			continue
		}
		// Expect criterias in order
		if criteria.Id != uint8(i+1) {
			t.Errorf("Criteria %d has unexpected id %d", i, criteria.Id)
			continue
		}
		// Expect between 2 and 9 laws per criteria
		if len(criteria.Laws) < 2 || len(criteria.Laws) > 9 {
			t.Errorf("Criteria %d has %d law(s)", i, len(criteria.Laws))
			continue
		}
		// Expect all laws to be valid
		lawSeen := make(map[uint8]bool, len(criteria.Laws))
		for li, law := range criteria.Laws {
			if law == nil {
				t.Errorf("Criteria %d's law %d is nil", i, li)
				continue
			}
			if lawSeen[law.Id] {
				t.Errorf("Criteria %d has repeating laws: %+v", i, criteria.Laws)
				break
			}
			lawSeen[law.Id] = true
			if law.Mask.HasNoSolution() {
				t.Errorf("Criteria %d's law %d has no solution", i, li)
			}
		}
		// Expect no duplicate laws
		seen := make([]game.CodeMask, 0, len(criteria.Laws))
		for _, law := range criteria.Laws {
			if slices.Contains(seen, law.Mask) {
				t.Fatalf("Criteria %d has duplicate law",
					criteria.Id)
			}
			seen = append(seen, law.Mask)
		}
	}
}
