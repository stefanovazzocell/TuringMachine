package api

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"stefanovazzoler.com/turingmachine/src/turingmachine/store"
)

// An API for turingmachine
type api struct {
	store *store.Store

	server *http.Server
	mux    *http.ServeMux

	config apiConfig
}

// Sets up a new API and registers itself as the server handler.
// The provided server should not be modified by other entities going forward.
func NewApi(server *http.Server, config apiConfig) (a *api, err error) {
	// Make sure the server is non-nil
	if server == nil {
		panic("the server cannot be nil")
	}

	a = &api{
		server: server,
		mux:    http.NewServeMux(),

		config: config,
	}

	// Check if we need to create the store
	createStore := config.StoreForceCreate
	if !createStore {
		_, err = os.Stat(config.StoreFileName)
		createStore = errors.Is(err, os.ErrNotExist)
	}
	// Create or open the store
	if createStore {
		a.store, err = store.CreateStore(config.StoreFileName)
	} else {
		a.store, err = store.OpenStore(config.StoreFileName)
	}
	if err != nil {
		return
	}

	// Register routes and set http handler
	a.registerRoutes()
	server.Handler = a.mux
	return
}

// Start listening for incoming connections
func (a api) ListenAndServe() {
	go func() {
		err := a.server.ListenAndServe()
		if err != http.ErrServerClosed {
			slog.Error("got error from ListenAndServe",
				"err", err)
			a.Close()
			os.Exit(-1)
		}
	}()
	slog.Info("listening",
		"addr", a.server.Addr)
}

// Start listening for incoming TLS connections
func (a api) ListenAndServeTLS(certFile, keyFile string) {
	go func() {
		err := a.server.ListenAndServeTLS(certFile, keyFile)
		if err != http.ErrServerClosed {
			slog.Error("got error from ListenAndServeTLS",
				"err", err)
		}
	}()
	slog.Info("listening on tls",
		"addr", a.server.Addr)
}

// A helper that waits for a valid interrupt
func (a api) AwaitInterrupt() {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan
	slog.Info("Interrupt received",
		"interrupt", (<-signalChan).String())
}

// Shuts down the http server and closes the store
func (a api) Close() {
	var err error
	// Shutdown http server
	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
	{
		defer cancelShutdown()
		err = a.server.Shutdown(gracefullCtx)
	}
	if err != nil {
		slog.Error("got error during server shutdown",
			"err", err)
	}
	// Closes store
	err = a.store.Close()
	if err != nil {
		slog.Error("got error while closing api",
			"err", err)
	}
}
