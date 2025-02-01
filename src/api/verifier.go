package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
)

type VerifyResponse struct {
	Check bool `json:"check"`
}

// Handles GET /api/verify?law=12&proposal=345
func (a *api) handleVerify(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if !query.Has("law") || !query.Has("proposal") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	law, err := strconv.ParseUint(query.Get("law"), 10, 8)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	code, err := game.CodeFromString(query.Get("proposal"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	valid, ok := game.CheckCode(uint8(law), code)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(VerifyResponse{
		Check: valid,
	})
}
