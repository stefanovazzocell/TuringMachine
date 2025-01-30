package game

// A state is a Game that is in progress with helpers to quickly process moves
type State struct {
	mask CodeMask
	Game Game
}

// Returns a new state from a given game
func StateFromGame(g Game) State {
	return State{
		Game: g,
		mask: g.GetMask(),
	}
}

// Returns a debug string
func (state State) Debug() string {
	return state.mask.GetCode().String() + "->" + state.Game.Debug()
}

// Returns true if the two states have the same CodeMask
func (state State) Similar(other State) bool {
	return state.mask.Equal(other.mask)
}

// Returns true if there is a unique solution for this game
func (state State) IsSolved() bool {
	return state.mask.Available() == 1
}

// Returns true if this state is invalid (i.e.: there are 0 possible solutions)
func (state State) IsInvalid() bool {
	return state.mask.HasNoSolution()
}

// Returns true if this state contains a redundant entry.
// It ignores games with a single choice
func (state State) HasRedundant() bool {
	choices := state.Game.NumberOfChoices()
	if choices <= 1 {
		return false
	}

	maskA, maskB, maskC, maskD, maskE, maskF :=
		state.Game[0].Mask(), state.Game[1].Mask(), state.Game[2].Mask(),
		state.Game[3].Mask(), state.Game[4].Mask(), state.Game[5].Mask()
	if state.mask.Equal(maskB.And(maskC).And(maskD).And(maskE).And(maskF)) {
		return true
	}
	if state.mask.Equal(maskA.And(maskC).And(maskD).And(maskE).And(maskF)) {
		return true
	}

	if choices <= 2 {
		return false
	}
	if state.mask.Equal(maskA.And(maskB).And(maskD).And(maskE).And(maskF)) {
		return true
	}

	if choices <= 3 {
		return false
	}
	if state.mask.Equal(maskA.And(maskB).And(maskC).And(maskE).And(maskF)) {
		return true
	}

	if choices <= 4 {
		return false
	}
	if state.mask.Equal(maskA.And(maskB).And(maskC).And(maskD).And(maskF)) {
		return true
	}

	if choices <= 5 {
		return false
	}
	return state.mask.Equal(maskA.And(maskB).And(maskC).And(maskD).And(maskE))
}

// Advances the last choice to the next law until one of the following
// conditions is met:
// 1. There are no further choices to make (returns false)
// 2. We find a valid choice (returns true)
//
// As a special case, if no choice was ever made, it returns false
func (state State) NextValidChoice(baseMask CodeMask) (State, bool) {
	idx := state.Game.NumberOfChoices() - 1
	if idx == -1 {
		return state, false
	}
	for {
		state.Game[idx] = state.Game[idx].NextLaw()
		if state.Game[idx] == BlankChoice {
			return state, false
		}
		state.mask = baseMask.And(state.Game[idx].Mask())
		if state.mask.Equal(baseMask) || state.mask.HasNoSolution() {
			continue
		}
		return state, true
	}
}

// Adds the next valid choice to the game. Returns false if none available.
func (state State) AddValidChoice() (State, bool) {
	idx := state.Game.NumberOfChoices()
	if idx == MaxNumberOfChoicesPerGame {
		return state, false
	}
	baseMask := state.mask
	if idx != 0 {
		// If any choice was already made, take the last as a starting point
		state.Game[idx] = state.Game[idx-1]
	}
	state.Game[idx] = state.Game[idx].NextCriteria()
	for {
		if state.Game[idx] == BlankChoice {
			return state, false
		}
		state.mask = baseMask.And(state.Game[idx].Mask())
		if state.mask.Equal(baseMask) || state.mask.HasNoSolution() {
			state.Game[idx] = state.Game[idx].NextLaw()
			continue
		}
		return state, true
	}
}
