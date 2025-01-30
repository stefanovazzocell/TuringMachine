package game

import (
	"fmt"
	"math"
)

const (
	// A special case of choice indicating no choice at all
	BlankChoice Choice = 0
	// The highest valid choice (also the number of valid choices)
	MaxChoice = 179
	// Number of easy choices (ignoring BlankChoice)
	numberOfEasyChoice = 48
	// Number of standard choices (ignoring BlankChoice)
	numberOfStandardChoice = 61
)

var (
	// Map Choice to *Criteria
	choiceToCriteria = [math.MaxUint8 + 1]*Criteria{
		nil, Criterias[0], Criterias[0], Criterias[1], Criterias[1], Criterias[1], Criterias[2], Criterias[2], Criterias[2], Criterias[3], Criterias[3], Criterias[3], Criterias[4], Criterias[4], Criterias[5], Criterias[5], Criterias[6], Criterias[6], Criterias[7], Criterias[7], Criterias[7], Criterias[8], Criterias[8], Criterias[8], Criterias[9], Criterias[9], Criterias[9], Criterias[10], Criterias[10], Criterias[10], Criterias[11], Criterias[11], Criterias[11], Criterias[12], Criterias[12], Criterias[12], Criterias[13], Criterias[13], Criterias[13], Criterias[14], Criterias[14], Criterias[14], Criterias[15], Criterias[15], Criterias[16], Criterias[16], Criterias[16], Criterias[16], Criterias[17], Criterias[17], Criterias[18], Criterias[18], Criterias[18], Criterias[19], Criterias[19], Criterias[19], Criterias[20], Criterias[20], Criterias[21], Criterias[21], Criterias[21], Criterias[22], Criterias[22], Criterias[22], Criterias[23], Criterias[23], Criterias[24], Criterias[24], Criterias[24], Criterias[25], Criterias[25], Criterias[25], Criterias[26], Criterias[26], Criterias[26], Criterias[27], Criterias[27], Criterias[27], Criterias[28], Criterias[28], Criterias[28], Criterias[29], Criterias[29], Criterias[29], Criterias[30], Criterias[30], Criterias[30], Criterias[31], Criterias[31], Criterias[31], Criterias[32], Criterias[32], Criterias[32], Criterias[32], Criterias[32], Criterias[32], Criterias[33], Criterias[33], Criterias[33], Criterias[34], Criterias[34], Criterias[34], Criterias[35], Criterias[35], Criterias[35], Criterias[36], Criterias[36], Criterias[36], Criterias[37], Criterias[37], Criterias[37], Criterias[38], Criterias[38], Criterias[38], Criterias[38], Criterias[38], Criterias[38], Criterias[39], Criterias[39], Criterias[39], Criterias[39], Criterias[39], Criterias[39], Criterias[39], Criterias[39], Criterias[39], Criterias[40], Criterias[40], Criterias[40], Criterias[40], Criterias[40], Criterias[40], Criterias[40], Criterias[40], Criterias[40], Criterias[41], Criterias[41], Criterias[41], Criterias[41], Criterias[41], Criterias[41], Criterias[42], Criterias[42], Criterias[42], Criterias[42], Criterias[42], Criterias[42], Criterias[43], Criterias[43], Criterias[43], Criterias[43], Criterias[43], Criterias[43], Criterias[44], Criterias[44], Criterias[44], Criterias[44], Criterias[44], Criterias[44], Criterias[45], Criterias[45], Criterias[45], Criterias[45], Criterias[45], Criterias[45], Criterias[46], Criterias[46], Criterias[46], Criterias[46], Criterias[46], Criterias[46], Criterias[47], Criterias[47], Criterias[47], Criterias[47], Criterias[47], Criterias[47], Criterias[47], Criterias[47], Criterias[47],
	}
	// Map from choice to law
	choiceToLaw = [math.MaxUint8 + 1]*Law{
		nil, laws[1], laws[16], laws[25], laws[3], laws[18], laws[28], laws[8], laws[21], laws[29], laws[9], laws[138], laws[34], laws[37], laws[35], laws[38], laws[36], laws[39], laws[40], laws[41], laws[42], laws[46], laws[47], laws[48], laws[49], laws[50], laws[51], laws[139], laws[89], laws[92], laws[140], laws[90], laws[93], laws[141], laws[91], laws[95], laws[116], laws[117], laws[118], laws[113], laws[114], laws[115], laws[131], laws[132], laws[85], laws[86], laws[87], laws[88], laws[55], laws[56], laws[137], laws[100], laws[136], laws[119], laws[120], laws[121], laws[81], laws[82], laws[133], laws[134], laws[135], laws[74], laws[60], laws[67], laws[83], laws[84], laws[122], laws[123], laws[124], laws[25], laws[28], laws[31], laws[26], laws[29], laws[32], laws[1], laws[6], laws[11], laws[3], laws[8], laws[13], laws[4], laws[9], laws[14], laws[16], laws[19], laws[22], laws[18], laws[21], laws[24], laws[34], laws[35], laws[36], laws[37], laws[38], laws[39], laws[128], laws[129], laws[130], laws[125], laws[126], laws[127], laws[57], laws[58], laws[59], laws[98], laws[103], laws[108], laws[100], laws[105], laws[110], laws[1], laws[6], laws[11], laws[16], laws[19], laws[22], laws[25], laws[28], laws[31], laws[3], laws[8], laws[13], laws[18], laws[21], laws[24], laws[26], laws[29], laws[32], laws[4], laws[9], laws[14], laws[138], laws[142], laws[143], laws[116], laws[117], laws[118], laws[113], laws[114], laws[115], laws[139], laws[89], laws[92], laws[140], laws[90], laws[93], laws[144], laws[89], laws[94], laws[141], laws[91], laws[95], laws[40], laws[41], laws[42], laws[46], laws[47], laws[48], laws[46], laws[47], laws[48], laws[49], laws[50], laws[51], laws[40], laws[41], laws[42], laws[49], laws[50], laws[51], laws[139], laws[140], laws[141], laws[89], laws[90], laws[91], laws[92], laws[93], laws[95],
	}
	// Map from choice to mask
	choiceToMask = [math.MaxUint8 + 1]CodeMask{
		BaseMask, laws[1].Mask, laws[16].Mask, laws[25].Mask, laws[3].Mask, laws[18].Mask, laws[28].Mask, laws[8].Mask, laws[21].Mask, laws[29].Mask, laws[9].Mask, laws[138].Mask, laws[34].Mask, laws[37].Mask, laws[35].Mask, laws[38].Mask, laws[36].Mask, laws[39].Mask, laws[40].Mask, laws[41].Mask, laws[42].Mask, laws[46].Mask, laws[47].Mask, laws[48].Mask, laws[49].Mask, laws[50].Mask, laws[51].Mask, laws[139].Mask, laws[89].Mask, laws[92].Mask, laws[140].Mask, laws[90].Mask, laws[93].Mask, laws[141].Mask, laws[91].Mask, laws[95].Mask, laws[116].Mask, laws[117].Mask, laws[118].Mask, laws[113].Mask, laws[114].Mask, laws[115].Mask, laws[131].Mask, laws[132].Mask, laws[85].Mask, laws[86].Mask, laws[87].Mask, laws[88].Mask, laws[55].Mask, laws[56].Mask, laws[137].Mask, laws[100].Mask, laws[136].Mask, laws[119].Mask, laws[120].Mask, laws[121].Mask, laws[81].Mask, laws[82].Mask, laws[133].Mask, laws[134].Mask, laws[135].Mask, laws[74].Mask, laws[60].Mask, laws[67].Mask, laws[83].Mask, laws[84].Mask, laws[122].Mask, laws[123].Mask, laws[124].Mask, laws[25].Mask, laws[28].Mask, laws[31].Mask, laws[26].Mask, laws[29].Mask, laws[32].Mask, laws[1].Mask, laws[6].Mask, laws[11].Mask, laws[3].Mask, laws[8].Mask, laws[13].Mask, laws[4].Mask, laws[9].Mask, laws[14].Mask, laws[16].Mask, laws[19].Mask, laws[22].Mask, laws[18].Mask, laws[21].Mask, laws[24].Mask, laws[34].Mask, laws[35].Mask, laws[36].Mask, laws[37].Mask, laws[38].Mask, laws[39].Mask, laws[128].Mask, laws[129].Mask, laws[130].Mask, laws[125].Mask, laws[126].Mask, laws[127].Mask, laws[57].Mask, laws[58].Mask, laws[59].Mask, laws[98].Mask, laws[103].Mask, laws[108].Mask, laws[100].Mask, laws[105].Mask, laws[110].Mask, laws[1].Mask, laws[6].Mask, laws[11].Mask, laws[16].Mask, laws[19].Mask, laws[22].Mask, laws[25].Mask, laws[28].Mask, laws[31].Mask, laws[3].Mask, laws[8].Mask, laws[13].Mask, laws[18].Mask, laws[21].Mask, laws[24].Mask, laws[26].Mask, laws[29].Mask, laws[32].Mask, laws[4].Mask, laws[9].Mask, laws[14].Mask, laws[138].Mask, laws[142].Mask, laws[143].Mask, laws[116].Mask, laws[117].Mask, laws[118].Mask, laws[113].Mask, laws[114].Mask, laws[115].Mask, laws[139].Mask, laws[89].Mask, laws[92].Mask, laws[140].Mask, laws[90].Mask, laws[93].Mask, laws[144].Mask, laws[89].Mask, laws[94].Mask, laws[141].Mask, laws[91].Mask, laws[95].Mask, laws[40].Mask, laws[41].Mask, laws[42].Mask, laws[46].Mask, laws[47].Mask, laws[48].Mask, laws[46].Mask, laws[47].Mask, laws[48].Mask, laws[49].Mask, laws[50].Mask, laws[51].Mask, laws[40].Mask, laws[41].Mask, laws[42].Mask, laws[49].Mask, laws[50].Mask, laws[51].Mask, laws[139].Mask, laws[140].Mask, laws[141].Mask, laws[89].Mask, laws[90].Mask, laws[91].Mask, laws[92].Mask, laws[93].Mask, laws[95].Mask,
	}
	// Map from (CriteriaId-1) to first Choice
	nextCriteria = [MaxChoice + 1]Choice{
		Choice(1), Choice(3), Choice(3), Choice(6), Choice(6), Choice(6), Choice(9), Choice(9), Choice(9), Choice(12), Choice(12), Choice(12), Choice(14), Choice(14), Choice(16), Choice(16), Choice(18), Choice(18), Choice(21), Choice(21), Choice(21), Choice(24), Choice(24), Choice(24), Choice(27), Choice(27), Choice(27), Choice(30), Choice(30), Choice(30), Choice(33), Choice(33), Choice(33), Choice(36), Choice(36), Choice(36), Choice(39), Choice(39), Choice(39), Choice(42), Choice(42), Choice(42), Choice(44), Choice(44), Choice(48), Choice(48), Choice(48), Choice(48), Choice(50), Choice(50), Choice(53), Choice(53), Choice(53), Choice(56), Choice(56), Choice(56), Choice(58), Choice(58), Choice(61), Choice(61), Choice(61), Choice(64), Choice(64), Choice(64), Choice(66), Choice(66), Choice(69), Choice(69), Choice(69), Choice(72), Choice(72), Choice(72), Choice(75), Choice(75), Choice(75), Choice(78), Choice(78), Choice(78), Choice(81), Choice(81), Choice(81), Choice(84), Choice(84), Choice(84), Choice(87), Choice(87), Choice(87), Choice(90), Choice(90), Choice(90), Choice(96), Choice(96), Choice(96), Choice(96), Choice(96), Choice(96), Choice(99), Choice(99), Choice(99), Choice(102), Choice(102), Choice(102), Choice(105), Choice(105), Choice(105), Choice(108), Choice(108), Choice(108), Choice(111), Choice(111), Choice(111), Choice(117), Choice(117), Choice(117), Choice(117), Choice(117), Choice(117), Choice(126), Choice(126), Choice(126), Choice(126), Choice(126), Choice(126), Choice(126), Choice(126), Choice(126), Choice(135), Choice(135), Choice(135), Choice(135), Choice(135), Choice(135), Choice(135), Choice(135), Choice(135), Choice(141), Choice(141), Choice(141), Choice(141), Choice(141), Choice(141), Choice(147), Choice(147), Choice(147), Choice(147), Choice(147), Choice(147), Choice(153), Choice(153), Choice(153), Choice(153), Choice(153), Choice(153), Choice(159), Choice(159), Choice(159), Choice(159), Choice(159), Choice(159), Choice(165), Choice(165), Choice(165), Choice(165), Choice(165), Choice(165), Choice(171), Choice(171), Choice(171), Choice(171), Choice(171), Choice(171),
	}
	// Map from choice to criteria mask
	choiceToCriteriaIdMask = [MaxChoice + 1]uint64{
		0, 1 << 0, 1 << 0, 1 << 1, 1 << 1, 1 << 1, 1 << 2, 1 << 2, 1 << 2, 1 << 3, 1 << 3, 1 << 3, 1 << 4, 1 << 4, 1 << 5, 1 << 5, 1 << 6, 1 << 6, 1 << 7, 1 << 7, 1 << 7, 1 << 8, 1 << 8, 1 << 8, 1 << 9, 1 << 9, 1 << 9, 1 << 10, 1 << 10, 1 << 10, 1 << 11, 1 << 11, 1 << 11, 1 << 12, 1 << 12, 1 << 12, 1 << 13, 1 << 13, 1 << 13, 1 << 14, 1 << 14, 1 << 14, 1 << 15, 1 << 15, 1 << 16, 1 << 16, 1 << 16, 1 << 16, 1 << 17, 1 << 17, 1 << 18, 1 << 18, 1 << 18, 1 << 19, 1 << 19, 1 << 19, 1 << 20, 1 << 20, 1 << 21, 1 << 21, 1 << 21, 1 << 22, 1 << 22, 1 << 22, 1 << 23, 1 << 23, 1 << 24, 1 << 24, 1 << 24, 1 << 25, 1 << 25, 1 << 25, 1 << 26, 1 << 26, 1 << 26, 1 << 27, 1 << 27, 1 << 27, 1 << 28, 1 << 28, 1 << 28, 1 << 29, 1 << 29, 1 << 29, 1 << 30, 1 << 30, 1 << 30, 1 << 31, 1 << 31, 1 << 31, 1 << 32, 1 << 32, 1 << 32, 1 << 32, 1 << 32, 1 << 32, 1 << 33, 1 << 33, 1 << 33, 1 << 34, 1 << 34, 1 << 34, 1 << 35, 1 << 35, 1 << 35, 1 << 36, 1 << 36, 1 << 36, 1 << 37, 1 << 37, 1 << 37, 1 << 38, 1 << 38, 1 << 38, 1 << 38, 1 << 38, 1 << 38, 1 << 39, 1 << 39, 1 << 39, 1 << 39, 1 << 39, 1 << 39, 1 << 39, 1 << 39, 1 << 39, 1 << 40, 1 << 40, 1 << 40, 1 << 40, 1 << 40, 1 << 40, 1 << 40, 1 << 40, 1 << 40, 1 << 41, 1 << 41, 1 << 41, 1 << 41, 1 << 41, 1 << 41, 1 << 42, 1 << 42, 1 << 42, 1 << 42, 1 << 42, 1 << 42, 1 << 43, 1 << 43, 1 << 43, 1 << 43, 1 << 43, 1 << 43, 1 << 44, 1 << 44, 1 << 44, 1 << 44, 1 << 44, 1 << 44, 1 << 45, 1 << 45, 1 << 45, 1 << 45, 1 << 45, 1 << 45, 1 << 46, 1 << 46, 1 << 46, 1 << 46, 1 << 46, 1 << 46, 1 << 47, 1 << 47, 1 << 47, 1 << 47, 1 << 47, 1 << 47, 1 << 47, 1 << 47, 1 << 47,
	}
)

// A choice is a combination of a criteria card and an associated law.
// Choice maps are automatically populated on init.
type Choice uint8

// Returns a choice from a given criteria + verifier. If not found returns false
func ChoiceFromCriteriaVerifier(criteria uint8, verifier uint16) (choice Choice, ok bool) {
	choice = nextCriteria[BlankChoice]
	// Find the first Choice with the given criteria
	for choice <= MaxChoice && choice.Criteria().Id != criteria {
		choice = nextCriteria[choice]
	}
	if choice > MaxChoice {
		return
	}
	// Find the matching law
	for {
		if choice > MaxChoice || choice.Criteria().Id != criteria {
			return
		}

		if choice.Law().VerificationCard.hasSymbol(verifier) {
			ok = true // Ot matches a symbol
		}
		if verifier <= math.MaxUint8 && choice.Law().Id == uint8(verifier) {
			ok = true // It matches a law id
		}
		if ok {
			return
		}

		choice++
	}
}

// Returns a debug string
func (choice Choice) Debug() string {
	if choice == BlankChoice {
		return "[blank]"
	}
	if !choice.IsValid() {
		return "[invalid]"
	}
	return fmt.Sprintf("[%d:%s]",
		choice.Criteria().Id, choice.Law().Description)
}

// Returns true if a choice is valid (or blank)
func (choice Choice) IsValid() bool {
	return choice <= MaxChoice
}

// Returns a mask for this choice's criteria id. Useful to count the number of
// unique choices made
func (choice Choice) CriteriaIdMask() uint64 {
	return choiceToCriteriaIdMask[choice]
}

// Returns the law associated with this choice if any
func (choice Choice) Law() *Law {
	return choiceToLaw[choice]
}

// Returns the criteria associated with this choice if any
func (choice Choice) Criteria() *Criteria {
	return choiceToCriteria[choice]
}

// Returns the difficulty of the criteria associated with this choice.
// If no such criteria, returns HardDifficulty (or EasyDifficulty for blank).
func (choice Choice) Difficulty() Difficulty {
	if choice < numberOfEasyChoice {
		return EasyDifficulty
	}
	if choice < numberOfStandardChoice {
		return StandardDifficulty
	}
	return HardDifficulty
}

// Advances to the next valid law if any is available, otherwise returns
// BlankChoice.
func (choice Choice) NextLaw() Choice {
	if choice >= MaxChoice {
		return BlankChoice
	}
	return choice + 1
}

// Returns to the next valid criteria if any is available, otherwise returns
// BlankChoice.
// As a special case, BlankChoice is mapped to the first valid criteria
func (choice Choice) NextCriteria() Choice {
	if choice >= MaxChoice {
		return BlankChoice
	}
	return nextCriteria[choice]
}

// If the choice is valid returns the mask for this choice's law;
// otherwise returns NewCodeMask()
func (choice Choice) Mask() CodeMask {
	if choice > MaxChoice {
		return BaseMask
	}
	return choiceToMask[choice]
}
