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

func (r *RateLimiter) Acquire() {
	for r.acquireNoWait() == 0 {
		time.Sleep(time.Second / time.Duration(r.qps))
	}
}

func (r *RateLimiter) AcquireContext(ctx context.Context) {
	for r.acquireNoWait() == 0 {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second / time.Duration(r.qps))
		}
	}
}

func (r *RateLimiter) AcquireNoWait() bool {
	return r.acquireNoWait() == 1
}

func (r *RateLimiter) acquireNoWait() int {
	//if atomic.LoadInt32(&r.stopped) == 1 {
	if r.stopped == 1 {
		return -1
	}

	//if atomic.LoadInt64(&r.tokens) <= 0 {
	if r.tokens <= 0 {
		return 0
	}
	atomic.AddInt64(&r.tokens, -1)
	return 1
}

func (r *RateLimiter) Stop() {
	//atomic.StoreInt32(&r.stopped, 1)
	r.stopped = 1
	close(r.stop)
}

type RateLimiter2 struct {
	q         *Queue
	cur       int
	sum       int
	pos       int
	subWin    int
	subWinNum int
	limit     int
}

func NewRateLimiter2(QPS int) *RateLimiter2 {
	subWin := 1
	subWinNum := 5
	q := NewQueue(subWinNum + 1)

	return &RateLimiter2{
		q:         q,
		cur:       0,
		sum:       0,
		subWin:    subWin,
		subWinNum: subWinNum,
		limit:     QPS * subWin * subWinNum,
	}
}

func (r *RateLimiter2) acquireNoWait() int {
	pos := int(time.Now().Unix()) / r.subWin
	if pos > r.pos {
		r.q.EnQueue([2]int{r.pos, r.cur})
		r.sum += r.cur
		r.pos = pos
		r.cur = 0

		leftPos := pos - r.subWinNum
		for head, ok := r.q.Head(); ok; head, ok = r.q.Head() {
			v := head.([2]int)
			p, sub := v[0], v[1]
			if p >= leftPos {
				break
			}
			r.q.DeQueue()
			r.sum -= sub
		}
	}
	//println(r.limit)
	if r.sum+r.cur > r.limit {
		//println(r.sum+r.cur, r.limit)
		return 0
	}
	r.cur++
	return 1
}

func (r *RateLimiter2) Acquire() {
	for r.acquireNoWait() == 0 {
		time.Sleep(time.Second / time.Duration(r.limit/(r.subWin*r.subWinNum)))
	}
}

func (r *RateLimiter2) AcquireNoWait() bool {
	return r.acquireNoWait() == 1
}
