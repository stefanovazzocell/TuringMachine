package game

var laws = map[uint8]*Law{
	1: newLaw(1, 45, "△ = 1", func(c Code) bool {
		return c.Triangle() == 1
	}),
	3: newLaw(3, 33, "△ = 3", func(c Code) bool {
		return c.Triangle() == 3
	}),
	4: newLaw(4, 67, "△ = 4", func(c Code) bool {
		return c.Triangle() == 4
	}),
	// 5: newLaw(5, 81, "△ = 5", func(c Code) bool {
	// 	return c.Triangle() == 5
	// }),
	6: newLaw(6, 60, "□ = 1", func(c Code) bool {
		return c.Square() == 1
	}),
	8: newLaw(8, 40, "□ = 3", func(c Code) bool {
		return c.Square() == 3
	}),
	9: newLaw(9, 38, "□ = 4", func(c Code) bool {
		return c.Square() == 4
	}),
	// 10: newLaw(10, 78, "□ = 5", func(c Code) bool {
	// 	return c.Square() == 5
	// }),
	11: newLaw(11, 6, "○ = 1", func(c Code) bool {
		return c.Circle() == 1
	}),
	13: newLaw(13, 53, "○ = 3", func(c Code) bool {
		return c.Circle() == 3
	}),
	14: newLaw(14, 0, "○ = 4", func(c Code) bool {
		return c.Circle() == 4
	}),
	// 15: newLaw(15, 82, "○ = 5", func(c Code) bool {
	// 	return c.Circle() == 5
	// }),

	16: newLaw(16, 72, "△ > 1", func(c Code) bool {
		return c.Triangle() > 1
	}),
	18: newLaw(18, 36, "△ > 3", func(c Code) bool {
		return c.Triangle() > 3
	}),
	19: newLaw(19, 34, "□ > 1", func(c Code) bool {
		return c.Square() > 1
	}),
	21: newLaw(21, 21, "□ > 3", func(c Code) bool {
		return c.Square() > 3
	}),
	22: newLaw(22, 35, "○ > 1", func(c Code) bool {
		return c.Circle() > 1
	}),
	24: newLaw(24, 41, "○ > 3", func(c Code) bool {
		return c.Circle() > 3
	}),

	25: newLaw(25, 91, "△ < 3", func(c Code) bool {
		return c.Triangle() < 3
	}),
	26: newLaw(26, 37, "△ < 4", func(c Code) bool {
		return c.Triangle() < 4
	}),
	28: newLaw(28, 28, "□ < 3", func(c Code) bool {
		return c.Square() < 3
	}),
	29: newLaw(29, 76, "□ < 4", func(c Code) bool {
		return c.Square() < 4
	}),
	31: newLaw(31, 30, "○ < 3", func(c Code) bool {
		return c.Circle() < 3
	}),
	32: newLaw(32, 93, "○ < 4", func(c Code) bool {
		return c.Circle() < 4
	}),

	34: newLaw(34, 88, "△ is even", func(c Code) bool {
		return fnEven(c.Triangle())
	}),
	35: newLaw(35, 43, "□ is even", func(c Code) bool {
		return fnEven(c.Square())
	}),
	36: newLaw(36, 31, "○ is even", func(c Code) bool {
		return fnEven(c.Circle())
	}),
	37: newLaw(37, 58, "△ is odd", func(c Code) bool {
		return fnOdd(c.Triangle())
	}),
	38: newLaw(38, 51, "□ is odd", func(c Code) bool {
		return fnOdd(c.Square())
	}),
	39: newLaw(39, 73, "○ is odd", func(c Code) bool {
		return fnOdd(c.Circle())
	}),

	40: newLaw(40, 1, "no 1s", func(c Code) bool {
		return fnCountDigits(c, 1) == 0
	}),
	41: newLaw(41, 23, "one 1", func(c Code) bool {
		return fnCountDigits(c, 1) == 1
	}),
	42: newLaw(42, 39, "two 1s", func(c Code) bool {
		return fnCountDigits(c, 1) == 2
	}),
	46: newLaw(46, 71, "no 3s", func(c Code) bool {
		return fnCountDigits(c, 3) == 0
	}),
	47: newLaw(47, 92, "one 3", func(c Code) bool {
		return fnCountDigits(c, 3) == 1
	}),
	48: newLaw(48, 2, "two 3s", func(c Code) bool {
		return fnCountDigits(c, 3) == 2
	}),
	49: newLaw(49, 26, "no 4s", func(c Code) bool {
		return fnCountDigits(c, 4) == 0
	}),
	50: newLaw(50, 27, "one 4", func(c Code) bool {
		return fnCountDigits(c, 4) == 1
	}),
	51: newLaw(51, 55, "two 4s", func(c Code) bool {
		return fnCountDigits(c, 4) == 2
	}),

	55: newLaw(55, 3, "△ + □ + ○ = even", func(c Code) bool {
		return fnEven(c.sum())
	}),
	56: newLaw(56, 61, "△ + □ + ○ = odd", func(c Code) bool {
		return fnOdd(c.sum())
	}),

	57: newLaw(57, 59, "△ + □ + ○ = 3x", func(c Code) bool {
		return fnMultiple(c.sum(), 3)
	}),
	58: newLaw(58, 12, "△ + □ + ○ = 4x", func(c Code) bool {
		return fnMultiple(c.sum(), 4)
	}),
	59: newLaw(59, 54, "△ + □ + ○ = 5x", func(c Code) bool {
		return fnMultiple(c.sum(), 5)
	}),

	60: newLaw(60, 14, "△ + □ + ○ = 6", func(c Code) bool {
		return c.sum() == 6
	}),
	67: newLaw(67, 77, "△ + □ + ○ > 6", func(c Code) bool {
		return c.sum() > 6
	}),
	74: newLaw(74, 25, "△ + □ + ○ < 6", func(c Code) bool {
		return c.sum() < 6
	}),

	81: newLaw(81, 46, "no pairs", func(c Code) bool {
		return !fnHasPair(c)
	}),
	82: newLaw(82, 20, "pairs", fnHasPair),

	83: newLaw(83, 22, "no numbers in ascending order", func(c Code) bool {
		return fnSequenceAscending(c) == 0
	}),
	84: newLaw(84, 10, "2 numbers in ascending order", func(c Code) bool {
		return fnSequenceAscending(c) == 1
	}),

	85: newLaw(85, 50, "no even two three", func(c Code) bool {
		return fnCountEven(c) == 0
	}),
	86: newLaw(86, 52, "one even two odd", func(c Code) bool {
		return fnCountEven(c) == 1
	}),
	87: newLaw(87, 66, "two even one odd", func(c Code) bool {
		return fnCountEven(c) == 2
	}),
	88: newLaw(88, 24, "three even no odd", func(c Code) bool {
		return fnCountEven(c) == 3
	}),

	89: newLaw(89, 84, "△ = □", func(c Code) bool {
		return c.Triangle() == c.Square()
	}),
	90: newLaw(90, 7, "△ = ○", func(c Code) bool {
		return c.Triangle() == c.Circle()
	}),
	91: newLaw(91, 90, "□ = ○", func(c Code) bool {
		return c.Square() == c.Circle()
	}),
	92: newLaw(92, 83, "△ > □", func(c Code) bool {
		return c.Triangle() > c.Square()
	}),
	93: newLaw(93, 79, "△ > ○", func(c Code) bool {
		return c.Triangle() > c.Circle()
	}),
	94: newLaw(94, 56, "□ > △", func(c Code) bool {
		return c.Square() > c.Triangle()
	}),
	95: newLaw(95, 68, "□ > ○", func(c Code) bool {
		return c.Square() > c.Circle()
	}),
	// 96: newLaw(96, 65, "○ > △", func(c Code) bool {
	// 	return c.Circle() > c.Triangle()
	// }),
	// 97: newLaw(97, 94, "○ > □", func(c Code) bool {
	// 	return c.Circle() > c.Square()
	// }),

	98: newLaw(98, 80, "△ + □ = 4", func(c Code) bool {
		return c.Triangle()+c.Square() == 4
	}),
	100: newLaw(100, 87, "△ + □ = 6", func(c Code) bool {
		return c.Triangle()+c.Square() == 6
	}),
	103: newLaw(103, 4, "△ + ○ = 4", func(c Code) bool {
		return c.Triangle()+c.Circle() == 4
	}),
	105: newLaw(105, 63, "△ + ○ = 6", func(c Code) bool {
		return c.Triangle()+c.Circle() == 6
	}),
	108: newLaw(108, 29, "□ + ○ = 4", func(c Code) bool {
		return c.Square()+c.Circle() == 4
	}),
	110: newLaw(110, 9, "□ + ○ = 6", func(c Code) bool {
		return c.Square()+c.Circle() == 6
	}),

	113: newLaw(113, 74, "△ > *", func(c Code) bool {
		return c.Triangle() > max(c.Square(), c.Circle())
	}),
	114: newLaw(114, 18, "□ > *", func(c Code) bool {
		return c.Square() > max(c.Triangle(), c.Circle())
	}),
	115: newLaw(115, 64, "○ > *", func(c Code) bool {
		return c.Circle() > max(c.Triangle(), c.Square())
	}),
	116: newLaw(116, 62, "△ < *", func(c Code) bool {
		return c.Triangle() < min(c.Square(), c.Circle())
	}),
	117: newLaw(117, 13, "□ < *", func(c Code) bool {
		return c.Square() < min(c.Triangle(), c.Circle())
	}),
	118: newLaw(118, 17, "○ < *", func(c Code) bool {
		return c.Circle() < min(c.Triangle(), c.Square())
	}),

	119: newLaw(119, 8, "a triple number", func(c Code) bool {
		square := c.Square()
		return c.Triangle() == square && square == c.Circle()
	}),
	120: newLaw(120, 85, "a double number", fnHasPair),
	121: newLaw(121, 49, "no repetition", fnNoRepeats),

	122: newLaw(122, 15, "no sequence of numbers in ascending or descending order", func(c Code) bool {
		return fnSequenceAscending(c)+fnSequenceDescending(c) == 0
	}),
	123: newLaw(123, 86, "2 numbers in ascending or descending order", func(c Code) bool {
		return fnSequenceAscending(c) == 1 || fnSequenceDescending(c) == 1
	}),
	124: newLaw(124, 48, "3 numbers in ascending or descending order", func(c Code) bool {
		return fnSequenceAscending(c) == 2 || fnSequenceDescending(c) == 2
	}),

	125: newLaw(125, 75, "△ >= *", func(c Code) bool {
		return c.Triangle() >= max(c.Square(), c.Circle())
	}),
	126: newLaw(126, 89, "□ >= *", func(c Code) bool {
		return c.Square() >= max(c.Triangle(), c.Circle())
	}),
	127: newLaw(127, 19, "○ >= *", func(c Code) bool {
		return c.Circle() >= max(c.Triangle(), c.Square())
	}),
	128: newLaw(128, 69, "△ <= *", func(c Code) bool {
		return c.Triangle() <= min(c.Square(), c.Circle())
	}),
	129: newLaw(129, 57, "□ <= *", func(c Code) bool {
		return c.Square() <= min(c.Triangle(), c.Circle())
	}),
	130: newLaw(130, 42, "○ <= *", func(c Code) bool {
		return c.Circle() <= min(c.Triangle(), c.Square())
	}),

	131: newLaw(131, 70, "even > odd", func(c Code) bool {
		return fnCountEven(c) >= 2
	}),
	132: newLaw(132, 47, "even < odd", func(c Code) bool {
		return fnCountEven(c) <= 1
	}),

	133: newLaw(133, 32, "ascending order", func(c Code) bool {
		square := c.Square()
		return c.Triangle() < square && square < c.Circle()
	}),
	134: newLaw(134, 11, "descending order", func(c Code) bool {
		square := c.Square()
		return c.Triangle() > square && square > c.Circle()
	}),
	135: newLaw(135, 44, "no order", fnNoOrder),

	136: newLaw(136, 16, "△ + □ > 6", func(c Code) bool {
		return c.Triangle()+c.Square() > 6
	}),
	137: newLaw(137, 5, "△ + □ < 6", func(c Code) bool {
		return c.Triangle()+c.Square() < 6
	}),

	138: newLaw(138, 78, "□ > 4", func(c Code) bool {
		return c.Square() > 4
	}),

	139: newLaw(139, 56, "△ < □", func(c Code) bool {
		return c.Triangle() < c.Square()
	}),
	140: newLaw(140, 65, "△ < ○", func(c Code) bool {
		return c.Triangle() < c.Circle()
	}),
	141: newLaw(141, 94, "□ < ○", func(c Code) bool {
		return c.Square() < c.Circle()
	}),

	142: newLaw(142, 81, "△ > 4", func(c Code) bool {
		return c.Triangle() > 4
	}),
	143: newLaw(143, 82, "○ > 4", func(c Code) bool {
		return c.Circle() > 4
	}),

	144: newLaw(144, 83, "□ < △", func(c Code) bool {
		return c.Square() < c.Triangle()
	}),
}

// Helper that returns true if a given number is a multiple of another
func fnMultiple(u uint8, target uint8) bool {
	return u%target == 0
}

// Helper to check if a uint8 is even
func fnEven(u uint8) bool {
	return fnMultiple(u, 2)
}

// Helper to check if a uint8 is odd
func fnOdd(u uint8) bool {
	return u%2 == 1
}

// Helper that counts the number of even digits
func fnCountEven(c Code) (count int) {
	if fnEven(c.Triangle()) {
		count += 1
	}
	if fnEven(c.Square()) {
		count += 1
	}
	if fnEven(c.Circle()) {
		count += 1
	}
	return
}

// Helper that returns true if the digits are in neither ascending nor
// descending order
func fnNoOrder(c Code) bool {
	triangle, square, circle := c.Triangle(), c.Square(), c.Circle()
	return !((triangle < square && square < circle) || (triangle > square && square > circle))
}

// Helper that returns true if exactly two digits are equal
func fnHasPair(c Code) bool {
	triangle, square, circle := c.Triangle(), c.Square(), c.Circle()
	if triangle == square {
		return square != circle
	}
	return triangle == circle || square == circle
}

// Helper that returns true if all digits are different
func fnNoRepeats(c Code) bool {
	triangle, square, circle := c.Triangle(), c.Square(), c.Circle()
	return triangle != square && square != circle && triangle != circle
}

// Helper that counts the digits of a number matching a target
func fnCountDigits(c Code, target uint8) (count int) {
	if c.Triangle() == target {
		count += 1
	}
	if c.Square() == target {
		count += 1
	}
	if c.Circle() == target {
		count += 1
	}
	return
}

// Helper that returns true if a and b are in ascending order
func fnIsAscending(a, b uint8) bool {
	return a+1 == b
}

// Helper that returns true if a and b are in descending order
func fnIsDescending(a, b uint8) bool {
	return a == b+1
}

// Helper that returns the number of digits that form a sequence of ascending
// numbers
func fnSequenceAscending(c Code) (count int) {
	square := c.Square()
	if fnIsAscending(c.Triangle(), square) {
		count += 1
	}
	if fnIsAscending(square, c.Circle()) {
		count += 1
	}
	return
}

// Helper that returns the number of digits that form a sequence of descending
// numbers
func fnSequenceDescending(c Code) (count int) {
	square := c.Square()
	if fnIsDescending(c.Triangle(), square) {
		count += 1
	}
	if fnIsDescending(square, c.Circle()) {
		count += 1
	}
	return
}
