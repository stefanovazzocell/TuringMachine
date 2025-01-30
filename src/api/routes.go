package api

import (
	"net/http"
)

// Registers all the server routes
func (a api) registerRoutes() {
	// GET /api/game?difficulty=hard&choices=5
	// GET /api/game?id=XXXXX
	a.mux.HandleFunc("GET /api/game", a.corsWrapper("GET", a.handleGetGame))
	// POST /api/solve {criterias: [...], verifiers: [...]}
	a.mux.HandleFunc("POST /api/solve", a.corsWrapper("POST", a.handleSolveGame))
	// GET /api/verify?law=12&proposal=345
	a.mux.HandleFunc("GET /api/verify", a.corsWrapper("GET", a.handleVerify))

	// Default handler
	a.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}
