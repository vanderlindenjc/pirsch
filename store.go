package pirsch

import (
	"time"
)

// Store is the database storage interface.
type Store interface {
	// SaveHits saves new hits.
	SaveHits([]Hit) error

	// Session returns the last path, time, and session timestamp for given client, fingerprint, and maximum age.
	Session(int64, string, time.Time) (string, time.Time, time.Time, error)

	// Count returns the number of results for given query.
	Count(string, ...interface{}) (int, error)

	// Get returns a single result for given query.
	// The result must be a pointer.
	Get(interface{}, string, ...interface{}) error

	// Select returns the results for given query.
	// The results must be a pointer to a slice.
	Select(interface{}, string, ...interface{}) error
}
