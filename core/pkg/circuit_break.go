package pkg

import (
	"github.com/sony/gobreaker/v2"
	"log/slog"
	"time"
)

func NewCircuitBreak[T any](name string, countRequests uint32, failureRatioIn float64) *gobreaker.CircuitBreaker[T] {
	var st gobreaker.Settings
	st.Name = name
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= countRequests && failureRatio >= failureRatioIn // todo pode ser variável de ambiente
	}
	// When to flush counters int the Closed state
	st.Interval = 5 * time.Second
	// Time to switch from Open to Half-open
	st.Timeout = 7 * time.Second

	st.OnStateChange = func(_ string, from gobreaker.State, to gobreaker.State) {
		// Handler for every state change. We'll use for debugging purpose
		slog.Error("state changed from " + from.String() + " to " + to.String())
	}

	return gobreaker.NewCircuitBreaker[T](st)
}
