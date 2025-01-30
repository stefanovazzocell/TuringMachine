package game

import "errors"

const (
	// A mask for a single digit
	digitMask uint8 = 0b111
	// The maximum value a single digit can have (5)
	upperDigit uint8 = 0b101
	// The minimum value a single digit can have (1)
	lowerDigit uint8 = 0b001
	// The swift for the triangle number
	triangleSwift = 6
	// The swift for the square number
	squareSwift = 3
	// Maximum value for Code
	MaxCode Code = 5<<triangleSwift + 5<<squareSwift + 5
	// Minimum value for Code
	MinCode Code = 1<<triangleSwift + 1<<squareSwift + 1
	// The number of possible codes
	numberOfCodes = 5 * 5 * 5
)

var (
	// Error returned when the code string is not 3-bytes long
	ErrBadCodeLength = errors.New("the code must have exactly 3 digits")
	// Error returned when the code string has an invalid character/digit
	ErrBadCodeDigit = errors.New("the code must only have digits between 1 and 5 (included)")
)

// A code represents a 3-digit number where each digit is in the range [1,5].
// It's represented as a 3 pairs of 3-bits (i.e. 9 bits) each representing one
// of the numbers.
type Code uint16

// Returns a code given a set of 3 digits.
// Note that this panics if the digits are outside of the allowed range [1,5].
func CodeFromNumbers(triangle, square, circle uint8) Code {
	if (triangle < lowerDigit || upperDigit < triangle) ||
		(square < lowerDigit || upperDigit < square) ||
		(circle < lowerDigit || upperDigit < circle) {
		panic("digit outside of allowed range")
	}
	return Code(uint16(triangle&digitMask)<<triangleSwift +
		uint16(square&digitMask)<<squareSwift +
		uint16(circle&digitMask))
}

// From a code index [0, 125) returns the code
func CodeFromIndex(idx uint8) Code {
	return CodeFromNumbers(((idx/25)%5 + 1), ((idx/5)%5 + 1), (idx%5 + 1))
}

// Returns the index of this code [0, 125)
func (code Code) GetIndex() uint8 {
	return (code.Triangle()-1)*25 + (code.Square()-1)*5 + (code.Circle() - 1)
}

// Parse a code from a string (i.e. "152" = CodeFromNumbers(1,5,2))
func CodeFromString(codeStr string) (code Code, err error) {
	if len(codeStr) != 3 {
		err = ErrBadCodeLength
		return
	}
	if (codeStr[0] < '1' || '5' < codeStr[0]) ||
		(codeStr[1] < '1' || '5' < codeStr[1]) ||
		(codeStr[2] < '1' || '5' < codeStr[2]) {
		err = ErrBadCodeDigit
		return
	}
	code = CodeFromNumbers(uint8(codeStr[0]-'0'),
		uint8(codeStr[1]-'0'),
		uint8(codeStr[2]-'0'))
	return
}

// Returns a code as a string
func (code Code) String() string {
	codeStr := make([]byte, 3)
	codeStr[2] = '0' + code.Circle()
	codeStr[1] = '0' + code.Square()
	codeStr[0] = '0' + code.Triangle()
	return string(codeStr)
}

// Returns the code as an integer
func (code Code) Int() int {
	return int(code.Triangle())*100 + int(code.Square())*10 + int(code.Circle())
}

// Increments code by 1 if it's not already Maxcode (in which case it returns
// it as is).
func (code Code) Incr() Code {
	if code == MaxCode {
		return MaxCode
	}
	// First digit (circle)
	if code.Circle() < 5 {
		code += 1
		return code
	}
	code -= 4
	// Second digit (square)
	if code.Square() < 5 {
		code += 1 << squareSwift
		return code
	}
	code -= 4 << squareSwift
	// Third digit (triangle)
	code += 1 << triangleSwift
	return code
}

// Returns the value for the triangle (left-most) number in the range [1,5]
func (code Code) Triangle() uint8 {
	return uint8(code>>triangleSwift) & digitMask
}

// Returns the value for the square (middle) number in the range [1,5]
func (code Code) Square() uint8 {
	return uint8(code>>squareSwift) & digitMask
}

// Returns the value for the circle (right-most) number in the range [1,5]
func (code Code) Circle() uint8 {
	return uint8(code) & digitMask
}

// Returns the sum of all the numbers in this code
func (code Code) sum() uint8 {
	return code.Circle() + code.Square() + code.Triangle()
}
