package api

import (
	"time"
)

const (
	DefaultStoreForceCreate = false
	DefaultShutdownTimeout  = 5 * time.Second
)

type apiConfig struct {
	// The file name for the store
	StoreFileName string
	// Set this option to force-recreate the store at init
	StoreForceCreate bool

	// The allowed origin(s) for CORS.
	// "*" allows all
	CorsOrigins string

	// Timeout for http server shutdown
	ShutdownTimeout time.Duration
}

// Returns an apiConfig with the default values
func NewAPIConfig(storeFileName string, corsOrigin string) apiConfig {
	return apiConfig{
		StoreFileName:    storeFileName,
		StoreForceCreate: DefaultStoreForceCreate,

		CorsOrigins: corsOrigin,

		ShutdownTimeout: DefaultShutdownTimeout,
	}
}
