package rateLimiter

import (
	"container/list"
	"time"
)

type RateLimiter struct {
	limit    int
	interval time.Duration
	list     *list.List
}

func New(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:    limit,
		interval: interval,
		list:     list.New(),
	}
}

// Idea is to use sliding window [mostOldTimestamp...mostRecentTimestamp]
// check if (now - mostOldTimestamp) exceeds rate limiting interval
// if it is, then next request is allowed and mostOldTimestamp was made too long ago
// pop mostOldTimestamp from queue, and as we allow the request we push now to the end of queue
// if not, means we are pushing requests too fast, reject current until more time pasts

func (r RateLimiter) Take() bool {
	now := time.Now()
	// push back timestamps to queue until we reach limit
	if r.limit > r.list.Len() {
		r.list.PushBack(now)
		return true
	}

	front := r.list.Front()
	diff := now.Sub(front.Value.(time.Time))
	if diff >= r.interval {
		r.list.Remove(front)
		r.list.PushBack(now)
		return true
	}
	return false
}
