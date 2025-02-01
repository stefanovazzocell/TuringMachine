package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/store"
)

type GameResponse struct {
	Id        string   `json:"id"`
	Code      string   `json:"code"`
	Criterias []int    `json:"criterias"`
	Verifiers []string `json:"verifiers"`
	Laws      []int    `json:"laws"`
}

// Writes a game into a responsewriter
// If the game has no solution responds with http.StatusBadRequest
func writeGameResponse(w http.ResponseWriter, g game.Game) {
	code, ok := g.Solve()
	if !ok {
		// This game does not have a solution
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	criteriaCards, verificationCards, laws := g.GetCards()

	_ = json.NewEncoder(w).Encode(GameResponse{
		Id:        g.String(),
		Code:      code.String(),
		Criterias: criteriaCards,
		Verifiers: verificationCards,
		Laws:      laws,
	})
}

// Handles GET /api/game
func (a *api) handleGetGame(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if query.Has("id") {
		a.handleGetGameById(w, query.Get("id"))
		return
	}
	a.handleGetGameRandom(w, query)
}

// Returns the number of choices requested or -1 on error.
// Choices must be in the range [2,6]
func getChoicesCount(choices string) (c int) {
	if len(choices) != 1 {
		return -1
	}
	c = int(choices[0] - '0')
	if c < 2 || c > 6 {
		return -1
	}
	return c
}

// Handles GET /api/game?difficulty=1&choices=5
func (a *api) handleGetGameRandom(w http.ResponseWriter, query url.Values) {
	// Try to identify the criterias/choices range
	var choices int = 6
	if query.Has("choices") {
		choices = getChoicesCount(query.Get("choices"))
	} else if query.Has("criterias") {
		choices = getChoicesCount(query.Get("criterias"))
	}
	if choices == -1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	start, end := a.store.GameRangeByChoices(choices)
	if !query.Has("choices") && !query.Has("criterias") {
		start = 0
	}

	// Fetch a random game that matches the description
	var g game.Game
	var err error
	if query.Has("difficulty") {
		// Try to identify the difficulty
		difficulty := game.HardDifficulty
		switch strings.ToLower(query.Get("difficulty")) {
		case "0", "easy":
			difficulty = game.EasyDifficulty
		case "1", "medium", "standard":
			difficulty = game.StandardDifficulty
		}

		// If it's a hard problem try to look it up in the DB first as those are
		// the most likely to get a hit.
		if difficulty == game.HardDifficulty {
			g, err = a.store.GetRandomGameInRangeWithDifficulty(start, end, difficulty)
		}
		// If the difficulty is easy/medium or we hit the max number of retries
		// in searching for a hard game, generate a random one now.
		if difficulty != game.HardDifficulty || err == store.ErrMaxRetries {
			g, err = game.RandomSolvableGame(choices, difficulty)
		}
	} else {
		g, err = a.store.GetRandomGameInRange(start, end)
	}
	if err != nil {
		slog.Warn("failed to get random game", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeGameResponse(w, g)
}

// Handles GET /api/game?id=XXXXX
func (a *api) handleGetGameById(w http.ResponseWriter, id string) {
	game, err := game.GameFromString(id)
	if err != nil || !game.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Sort the game (more likely to be valid)
	game.Sort()
	// For a single game it's faster to compute if it's valid or not
	if err = game.ValidateStrict(); err != nil {
		w.Header().Set("TM-Invalid-Game-Reason", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Write response
	writeGameResponse(w, game)
}
