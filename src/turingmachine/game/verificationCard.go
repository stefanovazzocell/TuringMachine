package game

import (
	"math/rand"
	"strconv"
)

const (
	NumberOfVerificationCards = 95
)

var (
	verificationCardToLozenge  = [NumberOfVerificationCards]uint16{201, 206, 215, 220, 227, 233, 244, 253, 261, 267, 274, 280, 286, 293, 302, 309, 315, 322, 329, 334, 339, 346, 350, 356, 360, 370, 376, 381, 387, 392, 396, 403, 407, 413, 419, 429, 434, 440, 447, 455, 462, 470, 475, 481, 485, 491, 497, 503, 507, 516, 525, 532, 537, 546, 551, 560, 566, 572, 578, 582, 588, 592, 596, 604, 609, 614, 618, 628, 632, 636, 640, 646, 650, 654, 661, 665, 670, 681, 688, 695, 701, 709, 717, 723, 737, 741, 746, 751, 758, 766, 771, 778, 782, 787, 795}
	verificationCardToPound    = [NumberOfVerificationCards]uint16{798, 793, 786, 781, 776, 770, 765, 757, 750, 744, 740, 736, 729, 720, 715, 708, 699, 694, 687, 680, 669, 664, 658, 653, 649, 645, 639, 635, 631, 627, 617, 613, 608, 599, 595, 591, 587, 581, 577, 571, 564, 558, 550, 543, 536, 530, 523, 515, 506, 502, 496, 490, 484, 480, 474, 469, 461, 454, 445, 439, 433, 424, 418, 412, 406, 402, 395, 391, 386, 379, 374, 369, 359, 355, 349, 344, 338, 332, 327, 319, 314, 308, 299, 289, 279, 273, 266, 257, 252, 243, 232, 224, 219, 213, 205}
	verificationCardToSlash    = [NumberOfVerificationCards]uint16{204, 212, 217, 223, 231, 237, 251, 256, 264, 270, 278, 282, 288, 296, 304, 312, 317, 325, 331, 337, 341, 348, 353, 358, 365, 373, 378, 385, 390, 394, 401, 405, 410, 416, 423, 432, 437, 442, 453, 459, 464, 472, 479, 483, 487, 495, 499, 505, 514, 520, 528, 534, 541, 549, 557, 563, 568, 576, 580, 586, 590, 594, 598, 606, 611, 616, 625, 630, 634, 638, 643, 648, 652, 657, 663, 668, 677, 686, 691, 697, 706, 714, 719, 726, 739, 743, 749, 755, 763, 769, 775, 780, 785, 792, 797}
	verificationCardToCurrency = [NumberOfVerificationCards]uint16{796, 790, 783, 779, 773, 767, 759, 754, 747, 742, 738, 733, 725, 718, 710, 704, 696, 690, 684, 673, 667, 662, 656, 651, 647, 641, 637, 633, 629, 621, 615, 610, 605, 597, 593, 589, 585, 579, 573, 567, 562, 553, 547, 540, 533, 527, 518, 509, 504, 498, 492, 486, 482, 476, 471, 463, 458, 449, 441, 435, 430, 421, 414, 409, 404, 399, 393, 389, 382, 377, 372, 362, 357, 352, 347, 340, 335, 330, 324, 316, 311, 303, 294, 287, 277, 268, 263, 255, 247, 236, 228, 221, 216, 207, 202}
)

const (
	LozengeSymbol  = "◊"
	PoundSymbol    = "#"
	SlashSymbol    = "/"
	CurrencySymbol = "¤"
)

// A verification card
type VerificationCard int8

// Returns a string representation of each card with a random symbol picked
func getRandomVerificationSymbol(cards []VerificationCard) []string {
	res := make([]string, len(cards))
	r := rand.Intn(4)
	for i := range len(cards) {
		switch r {
		case 0:
			res[i] = cards[i].LozengeString()
		case 1:
			res[i] = cards[i].PoundString()
		case 2:
			res[i] = cards[i].SlashString()
		case 3:
			res[i] = cards[i].CurrencyString()
		}
	}
	return res
}

// Returns true if the verificationCard is valid
func (verificationCard VerificationCard) valid() bool {
	return 0 <= verificationCard && verificationCard < NumberOfVerificationCards
}

// Returns true if the symbol code provided matches any code for this card
func (VerificationCard VerificationCard) hasSymbol(symbol uint16) bool {
	return verificationCardToLozenge[VerificationCard] == symbol ||
		verificationCardToPound[VerificationCard] == symbol ||
		verificationCardToSlash[VerificationCard] == symbol ||
		verificationCardToCurrency[VerificationCard] == symbol
}

// Returns the Lozenge representation of this verification card as a string
func (verificationCard VerificationCard) LozengeString() string {
	if verificationCard == -1 {
		return ""
	}
	return LozengeSymbol +
		strconv.FormatUint(uint64(verificationCardToLozenge[verificationCard]), 10)
}

// Returns the Pound representation of this verification card as a string
func (verificationCard VerificationCard) PoundString() string {
	if verificationCard == -1 {
		return ""
	}
	return PoundSymbol +
		strconv.FormatUint(uint64(verificationCardToPound[verificationCard]), 10)
}

// Returns the Slash representation of this verification card as a string
func (verificationCard VerificationCard) SlashString() string {
	if verificationCard == -1 {
		return ""
	}
	return SlashSymbol +
		strconv.FormatUint(uint64(verificationCardToSlash[verificationCard]), 10)
}

// Returns the Currency representation of this verification card as a string
func (verificationCard VerificationCard) CurrencyString() string {
	if verificationCard == -1 {
		return ""
	}
	return CurrencySymbol +
		strconv.FormatUint(uint64(verificationCardToCurrency[verificationCard]), 10)
}

// Returns the Lozenge representation of this verification card
func (verificationCard VerificationCard) Lozenge() uint16 {
	if verificationCard == -1 {
		return 0
	}
	return verificationCardToLozenge[verificationCard]
}

// Returns the Pound representation of this verification card
func (verificationCard VerificationCard) Pound() uint16 {
	if verificationCard == -1 {
		return 0
	}
	return verificationCardToPound[verificationCard]
}

// Returns the Slash representation of this verification card
func (verificationCard VerificationCard) Slash() uint16 {
	if verificationCard == -1 {
		return 0
	}
	return verificationCardToSlash[verificationCard]
}

// Returns the Currency representation of this verification card
func (verificationCard VerificationCard) Currency() uint16 {
	if verificationCard == -1 {
		return 0
	}
	return verificationCardToCurrency[verificationCard]
}
