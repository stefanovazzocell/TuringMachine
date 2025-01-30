package game

import (
	"math/bits"
)

const (
	maskHi = 0b0001111111111111111111111111111111111111111111111111111111111111
	maskLo = 0b1111111111111111111111111111111111111111111111111111111111111111
)

var (
	// A mask with all possible codes marked as available
	BaseMask = CodeMask{
		hi: maskHi,
		lo: maskLo,
	}
)

// A code mask represents one or more Code(s)
type CodeMask struct {
	hi uint64
	lo uint64
}

// Returns true if a code is green in this mask
func (cm CodeMask) Check(code Code) bool {
	idx := code.GetIndex()
	if idx <= 64 {
		return (cm.lo>>idx)&0b1 == 0b1
	}
	return (cm.hi>>(idx-64))&0b1 == 0b1
}

// Returns the count of all available codes.
func (cm CodeMask) Available() uint8 {
	return uint8(bits.OnesCount64(cm.hi) + bits.OnesCount64(cm.lo))
}

// Returns all the possible codes for this mask
func (cm CodeMask) GetAllCodes() []Code {
	codes := make([]Code, 0, cm.Available())
	if cm.lo != 0 {
		for i := uint8(0); i < 64; i++ {
			if cm.lo&0b1 == 0b1 {
				codes = append(codes, CodeFromIndex(i))
			}
			cm.lo = cm.lo >> 1
		}
	}
	if cm.hi != 0 {
		for i := uint8(0); i < 64; i++ {
			if cm.hi&0b1 == 0b1 {
				codes = append(codes, CodeFromIndex(i+64))
			}
			cm.hi = cm.hi >> 1
		}
	}
	return codes
}

// Returns the first available Code found.
// If no solution is found it returns an invalid code.
func (cm CodeMask) GetCode() Code {
	if cm.lo != 0 {
		for i := uint8(0); i < 64; i++ {
			if cm.lo&0b1 == 0b1 {
				return CodeFromIndex(i)
			}
			cm.lo = cm.lo >> 1
		}
	}
	if cm.hi != 0 {
		for i := uint8(0); i < 64; i++ {
			if cm.hi&0b1 == 0b1 {
				return CodeFromIndex(i + 64)
			}
			cm.hi = cm.hi >> 1
		}
	}
	return Code(0)
}

// Returns CodeMask after applying a function to all the numbers
func (s CodeMask) applyFn(fn func(c Code) bool) CodeMask {
	// lo
	for idx := uint8(0); idx < 64; idx++ {
		if (s.lo>>idx)&0b1 == 0b0 {
			// Already zero, skip
			continue
		}
		if fn(CodeFromIndex(idx)) {
			// It's also a match, continue
			continue
		}
		s.lo = s.lo & ^(0b1 << idx)
	}
	// hi
	for idx := uint8(0); idx < numberOfCodes-64; idx++ {
		if (s.hi>>idx)&0b1 == 0b0 {
			// Already zero, skip
			continue
		}
		if fn(CodeFromIndex(idx + 64)) {
			// It's also a match, continue
			continue
		}
		s.hi = s.hi & ^(0b1 << idx)
	}
	return s
}

// Returns true if no codes are available
func (s CodeMask) HasNoSolution() bool {
	return s.hi|s.lo == 0
}

// Returns the bitwise AND of cm And m (cm&m).
func (cm CodeMask) And(m CodeMask) CodeMask {
	return CodeMask{cm.hi & m.hi, cm.lo & m.lo}
}

// Returns true if cm and m match.
func (cm CodeMask) Equal(m CodeMask) bool {
	return cm.hi == m.hi && cm.lo == m.lo
}
