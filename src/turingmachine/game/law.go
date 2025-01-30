package game

// A law represents a function in this game such as "all digits are odd"
type Law struct {
	Description      string
	Mask             CodeMask
	VerificationCard VerificationCard
	Id               uint8
}

// Given a law and a code, it returns true if the code passes the law, false if
// it doesn't - and it reports ok as false if no such law exists
func CheckCode(law_id uint8, code Code) (valid bool, ok bool) {
	law, ok := laws[law_id]
	if !ok {
		return
	}
	valid = law.Mask.Check(code)
	return
}

// Creates a new criteria given an id, a verification card, a description and a
// function
func newLaw(id uint8, verificationCard VerificationCard, description string, fn func(c Code) bool) *Law {
	if !verificationCard.valid() {
		panic("invalid verification card")
	}
	return &Law{
		Id:               id,
		VerificationCard: verificationCard,
		Description:      description,
		Mask:             BaseMask.applyFn(fn),
	}
}
