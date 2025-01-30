package game_test

import (
	"strconv"
	"strings"
	"testing"

	"stefanovazzoler.com/turingmachine/src/turingmachine/game"
)

func numberMatchesString(n uint16, str string) bool {
	return strings.HasSuffix(str, strconv.FormatUint(uint64(n), 10))
}

func TestVerificationCard(t *testing.T) {
	t.Parallel()

	for vc := range game.VerificationCard(game.NumberOfVerificationCards) {
		// Lozenge
		n := vc.Lozenge()
		str := vc.LozengeString()
		if !numberMatchesString(n, str) {
			t.Errorf("[vc:%d] Lozenge mismatch found: %d != %q", vc, n, str)
		}
		// Currency
		n = vc.Currency()
		str = vc.CurrencyString()
		if !numberMatchesString(n, str) {
			t.Errorf("[vc:%d] Currency mismatch found: %d != %q", vc, n, str)
		}
		// Pound
		n = vc.Pound()
		str = vc.PoundString()
		if !numberMatchesString(n, str) {
			t.Errorf("[vc:%d] Pound mismatch found: %d != %q", vc, n, str)
		}
		// Slash
		n = vc.Slash()
		str = vc.SlashString()
		if !numberMatchesString(n, str) {
			t.Errorf("[vc:%d] Slash mismatch found: %d != %q", vc, n, str)
		}
	}
}
