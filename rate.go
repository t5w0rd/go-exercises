package exercises

import (
	"context"
	"sync/atomic"
	"time"
)

type RateLimiter struct {
	qps     int
	size    int
	tokens  int64
	stop    chan struct{}
	stopped int32
}

const RateLimiterFactor = 10

func NewRateLimiter(QPS int) *RateLimiter {
	size := QPS / RateLimiterFactor
	tb := &RateLimiter{
		qps:     QPS,
		size:    size,
		tokens:  int64(size),
		stop:    make(chan struct{}),
		stopped: 0,
	}

	go func() {
		tm := time.Now()
		ticker := time.NewTicker(time.Second / RateLimiterFactor)
		defer ticker.Stop()
		for {
			select {
			case <-tb.stop:
				return

			case <-ticker.C:
				now := time.Now()
				dur := now.Sub(tm)
				tm = now
				dt := int64(dur.Seconds() * float64(tb.qps))
				if tokens, size := tb.tokens+dt, int64(tb.size); tokens > size {
					atomic.StoreInt64(&tb.tokens, size)
				} else {
					atomic.AddInt64(&tb.tokens, dt)
				}
			}
		}
	}()
	return tb
}

func (tb *RateLimiter) Acquire() {
	for tb.acquireNoWait() == 0 {
		time.Sleep(time.Second / time.Duration(tb.qps))
	}
}

func (tb *RateLimiter) AcquireContext(ctx context.Context) {
	for tb.acquireNoWait() == 0 {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second / time.Duration(tb.qps))
		}
	}
}

func (tb *RateLimiter) AcquireNoWait() bool {
	return tb.acquireNoWait() == 1
}

func (tb *RateLimiter) acquireNoWait() int {
	//if atomic.LoadInt32(&tb.stopped) == 1 {
	if tb.stopped == 1 {
		return -1
	}

	//if atomic.LoadInt64(&tb.tokens) <= 0 {
	if tb.tokens <= 0 {
		return 0
	}
	atomic.AddInt64(&tb.tokens, -1)
	return 1
}

func (tb *RateLimiter) Stop() {
	//atomic.StoreInt32(&tb.stopped, 1)
	tb.stopped = 1
	close(tb.stop)
}
