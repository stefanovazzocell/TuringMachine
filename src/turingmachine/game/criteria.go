package game

// A criteria represents a criteria card
type Criteria struct {
	Id          uint8
	Description string
	Laws        []*Law
}

// Creates a new criteria given an id, a description and a set of laws
func newCriteria(id uint8, description string, laws []*Law) *Criteria {
	for _, law := range laws {
		if law == nil {
			panic("nil law")
		}
	}
	return &Criteria{
		Id:          id,
		Description: description,
		Laws:        laws,
	}
}

// Returns the difficulty of this criteria
func (criteria *Criteria) Difficulty() Difficulty {
	if criteria.Id <= difficultyEasyMaxCriteriaId {
		return EasyDifficulty
	}
	if criteria.Id <= difficultyStandardMaxCriteriaId {
		return StandardDifficulty
	}
	return HardDifficulty
}
