package metrics

import (
	"fmt"
	"sync"
	"time"
)

// A writer interface borrowed directly from io.
type Writer interface {
	Write([]byte) (int, error)
}

// A utility structure to collect metrics concurrently.
type Stats struct {
	mu     sync.RWMutex
	start  time.Time
	keys   []string
	values []int
}

// Creates or updates a stored metric and returns its value.
func (s *Stats) Add(k string, v int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.keys {
		if s.keys[i] == k {
			s.values[i] += v
			return s.values[i]
		}
	}
	s.keys = append(s.keys, k)
	s.values = append(s.values, v)
	return v
}

// Initializes metric storage and start time, clearing previous values.
func (s *Stats) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keys = []string{}
	s.values = []int{}
	s.start = time.Now()
}

// Returns the duration since Reset was called.
func (s *Stats) Duration() time.Duration {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return time.Since(s.start)
}

// Accepts a writer to print metrics and execution time.
//
// If no metrics exist, then no output will be written.
func (s *Stats) Print(w Writer) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i := range s.keys {
		fmt.Fprintf(w, "%s: %d\n", s.keys[i], s.values[i])
	}
	if len(s.keys) > 0 {
		fmt.Fprintf(w, "%s\n", s.Duration())
	}
}
