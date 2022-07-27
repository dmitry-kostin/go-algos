package main

import (
	"testing"
	"time"
)

func TestRateLimiter_InLimits(t *testing.T) {
	type conf struct {
		limit    int
		interval time.Duration
	}
	tests := []struct {
		name            string
		rateLimiterConf conf
		dataFlowConf    conf
		want            bool
	}{
		{
			name:            "1 event per 200ms",
			rateLimiterConf: conf{1, 200 * time.Millisecond},
			dataFlowConf:    conf{10, 201 * time.Millisecond},
			want:            true,
		},
		{
			name:            "5 events per 200ms",
			rateLimiterConf: conf{5, 200 * time.Millisecond},
			dataFlowConf:    conf{20, 50 * time.Millisecond},
			want:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := New(tt.rateLimiterConf.limit, tt.rateLimiterConf.interval)
			ticker := generateDataFlow(tt.dataFlowConf.limit, tt.dataFlowConf.interval)
			counter := 1
			for range ticker {
				now := time.Now()
				got := rl.Take()
				if got != tt.want {
					t.Errorf("take was rejected counter: %v, got: %v, time: %v", counter, got, now)
				}
				//t.Logf("counter: %v, got: %v, time: %v", counter, got, now)
				counter++
			}
		})
	}
}

// generateDataFlow generate n events in total, emit every d (seconds)
func generateDataFlow(n int, d time.Duration) chan interface{} {
	out := make(chan interface{})
	limiter := time.Tick(d)
	go func() {
		for i := 0; i < n; i++ {
			out <- struct{}{}
			<-limiter
		}
		close(out)
	}()
	return out
}
