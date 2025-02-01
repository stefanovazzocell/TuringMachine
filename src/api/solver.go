package api

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
)

type SolverResponse struct {
	Id        string   `json:"id"`
	Solutions []string `json:"solutions"`
	Criterias []int    `json:"criterias"`
	Verifiers []string `json:"verifiers"`
	Laws      []int    `json:"laws"`
}

// Writes a SolverResponse into a responsewriter
func writeSolverResponse(w http.ResponseWriter, g game.Game) {
	codes := g.GetMask().GetAllCodes()
	codesStr := make([]string, len(codes))
	for i := range len(codes) {
		codesStr[i] = codes[i].String()
	}

	criteriaCards, verificationCards, laws := g.GetCards()

	_ = json.NewEncoder(w).Encode(SolverResponse{
		Id:        g.String(),
		Solutions: codesStr,
		Criterias: criteriaCards,
		Verifiers: verificationCards,
		Laws:      laws,
	})
}

type SolverRequest struct {
	Criterias []int `json:"criterias"`
	Verifiers []int `json:"verifiers"`
}

func (sr SolverRequest) GetCriteriasVerifiers() (criterias []uint8, verifiers []uint16, ok bool) {
	// Request validation
	n := len(sr.Criterias)
	if n <= 0 || n > game.MaxNumberOfChoicesPerGame || n != len(sr.Verifiers) {
		return
	}
	for i := range n {
		if sr.Criterias[i] > math.MaxUint8 || sr.Verifiers[i] > math.MaxUint16 {
			return
		}
	}

	criterias = make([]uint8, n)
	verifiers = make([]uint16, n)
	for i := range n {
		criterias[i] = uint8(sr.Criterias[i])
		verifiers[i] = uint16(sr.Verifiers[i])
	}
	ok = true
	return
}

// Handles POST /api/solve {criterias: [...], verifiers: [...]}
func (a *api) handleSolveGame(w http.ResponseWriter, r *http.Request) {
	request := SolverRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	criterias, verifiers, ok := request.GetCriteriasVerifiers()
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	g, ok := game.GameFromCards(criterias, verifiers)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	writeSolverResponse(w, g)
}
