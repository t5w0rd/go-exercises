package exercises

import (
	"bytes"
	"sync"
	"time"
)

const (
	flagAck = uint8(1 << iota)
)

type frame struct {
	data  []byte
	flags uint8
	seq   uint32
}

func seqBefore(seq1, seq2 uint32) bool {
	return int32(seq1-seq2) < 0
}

type device interface {
	send(f *frame)
}

type sender struct {
	win  int
	seq  uint32
	mss  int
	buf  *bytes.Buffer
	mtx  sync.Mutex
	rto  time.Duration
	quit chan struct{}
	dev  device
}

func newSender(dev device, win uint16, isn uint32) *sender {
	return &sender{
		win: int(win),
		seq: isn,
		mss: 1000,
		buf: &bytes.Buffer{},
		dev: dev,
	}
}

func (s *sender) write(p []byte) (n int, err error) {
	s.mtx.Lock()
	s.mtx.Unlock()
	n, err = s.buf.Write(p)
	if s.buf.Len() > int(s.mss) {
		s.sendUnlock()
	}
	return n, err
}

func (s *sender) start() {
	s.mtx.Lock()
	if s.quit == nil {
		s.quit = make(chan struct{})
	}
	s.mtx.Unlock()

	// TODO: 合并定时器
	t := time.NewTicker(time.Millisecond * 1000)
	defer t.Stop()

	for {
		select {
		case <-s.quit:
			return
		case <-t.C:
			// TODO: 发送
			s.send()
		}
	}
}

func (s *sender) send() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.send()
}

func (s *sender) sendUnlock() {
	l := s.buf.Len()
	if l == 0 {
		return
	}

	for off := 0; off < s.win && off < s.buf.Len(); {
		var frameSize int
		if s.win < s.buf.Len() {
			if left := s.win - off; left < s.mss {
				frameSize = left
			} else {
				frameSize = s.mss
			}
		} else {
			if left := s.buf.Len() - off; left < s.mss {
				frameSize = left
			} else {
				frameSize = s.mss
			}
		}
		s.sendFrame(s.buf.Bytes()[off:off+frameSize], s.seq+uint32(off))
		off += frameSize
	}
}

func (s *sender) sendFrame(data []byte, seq uint32) {
	p := &frame{
		data: data,
		seq:  seq,
	}
	s.dev.send(p)
	// TODO: 注册超时重传定时器
}

func (s *sender) stop() {
	s.mtx.Lock()
	q := s.quit
	s.mtx.Unlock()

	if q == nil {
		return
	}
	close(q)
}

type receiver struct {
	win uint16
	ack uint32
	buf *bytes.Buffer
}
