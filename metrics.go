package metrics

// a function that creates, initializes, and returns a new stats collector.
func NewStats() *Stats {
	s := &Stats{}
	s.Reset()
	return s
}
