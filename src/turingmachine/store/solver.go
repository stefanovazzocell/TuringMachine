package store

import (
	"errors"
	"log/slog"
	"os"
	"runtime/pprof"
	"slices"
	"sync"
	"time"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
)

const (
	bufferMultiplier                 = 1 << 11
	writeBufferSize                  = game.MaxNumberOfChoicesPerGame * bufferMultiplier
	approximateNumberOfExpectedGames = 20881150
	defaultSolutionSize              = approximateNumberOfExpectedGames >> 2
)

// solves for all the possible games and writes the solutions to a file
func solve(filename string) (err error) {
	f, err := os.OpenFile("./cpu.prof", os.O_CREATE|os.O_RDWR|os.O_EXCL, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Delete any existing file (if present)
	err = os.Remove(filename)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		slog.Error("failed to delete existing file",
			"filename", filename,
			"err", err)
		return
	}

	// Start a thread for each initial choice/move. This way we can scale with
	// the number of available CPUs while allocating only a constant amount of
	// RAM.
	solvers := make([]*solver, game.MaxChoice)
	wg := sync.WaitGroup{}
	wg.Add(game.MaxChoice)
	start := time.Now()
	for i := range game.MaxChoice {
		go func(i int) {
			// Init the solver
			solvers[i] = newSolver()
			// Setup the first game state
			state := game.StateFromGame(game.Game{game.Choice(i + 1)})
			// Solve all games from this starting point
			solvers[i].solveNext(state)
			// Sort the solutions in this solver
			solvers[i].sortSolutions()

			wg.Done()
		}(i)
	}
	wg.Wait()
	slog.Info("discovered solutions",
		"solutions", countSolutions(solvers),
		"duration", time.Since(start).String())

	// Merge the sorted solutions
	start = time.Now()
	result := solutions(solvers)
	slog.Info("sorted solutions",
		"solutions", len(result),
		"duration", time.Since(start).String())

	// Store solutions
	start = time.Now()
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	buf := make([]byte, writeBufferSize)
	bufIdx := 0
	for i := range len(result) {
		// err = game.Write(file)
		result[i].WriteTo(buf, bufIdx*game.MaxNumberOfChoicesPerGame)
		bufIdx += 1
		if bufIdx == bufferMultiplier {
			// Write buffer
			_, err = file.Write(buf)
			if err != nil {
				slog.Error("got error while writing game to file",
					"err", err,
					"duration", time.Since(start).String())
				return err
			}
			bufIdx = 0
		}
	}
	// Write remaining data
	_, err = file.Write(buf[:bufIdx*game.MaxNumberOfChoicesPerGame])
	if err != nil {
		slog.Error("got error while writing game to file",
			"err", err,
			"duration", time.Since(start).String())
		return err
	}
	if err := file.Sync(); err != nil {
		slog.Error("got error while syncing file", "err", err)
		return err
	}
	if err := file.Close(); err != nil {
		slog.Error("got error while closing file", "err", err)
		return err
	}
	slog.Info("wrote all games",
		"filePath", filename,
		"duration", time.Since(start).String())
	return
}

// A set of solutions to the game
type solution []game.Game

// A game solver keeps track of solutions to the game
type solver struct {
	solution solution
}

// Returns all solutions (sorted) from a slice of non-nil solvers
func solutions(solvers []*solver) solution {
	solutions := make([]solution, len(solvers))
	for i := range solvers {
		solutions[i] = solvers[i].solution
	}
	return mergeSorted(solutions)
}

// Given a slice of non-nil solvers, returns the total number of games
func countSolutions(solvers []*solver) (count int) {
	for _, solver := range solvers {
		count += len(solver.solution)
	}
	return
}

// Returns a new solver
func newSolver() *solver {
	return &solver{
		solution: make(solution, 0, defaultSolutionSize),
	}
}

// Sorts this solution set
func (s *solver) sortSolutions() {
	slices.SortFunc(s.solution, func(a, b game.Game) int {
		return a.Value() - b.Value()
	})
}

// Sorts this solution set
func (s *solver) addSolution(g game.Game) {
	s.solution = append(s.solution, g)
}

func (s *solver) solveNext(state game.State) {
	// Base cases
	if state.IsInvalid() || state.HasRedundant() {
		return
	}
	if state.IsSolved() {
		s.addSolution(state.Game)
		return
	}
	if state.Game.NumberOfChoices() == game.MaxNumberOfChoicesPerGame {
		return
	}
	// Recursive case
	nextState, ok := state.AddValidChoice()

	for ok {
		s.solveNext(nextState)
		nextState, ok = nextState.NextValidChoice(state.Game.GetMask())
	}
}
