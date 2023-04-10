package log

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SyncLog struct {
	dataChan  chan []byte
	maxPer    uint32
	maxBuffer uint32
	lock      sync.Mutex
}

func NewSyncLog(buffer uint32) *SyncLog {
	l := new(SyncLog)
	l.maxPer = 0
	l.dataChan = make(chan []byte, buffer)
	return l
}

func (s *SyncLog) storageLog(message []byte) {
	defer s.lock.Unlock()
	s.lock.Lock()
	atomic.AddUint32(&s.maxPer, 1)
	s.dataChan <- message
	if s.maxPer >= s.maxBuffer/2 {
		go func() {
			s.sendLog()
			atomic.StoreUint32(&s.maxPer, s.maxBuffer-s.maxBuffer/2)
		}()
	}
}

func (s *SyncLog) sendLog() {
	for {
		select {
		case data, ok := <-s.dataChan:
			if !ok {
				return
			}
			fmt.Println("DATA:", string(data))
		}
	}
}

func (s *SyncLog) Write(p []byte) (n int, err error) {
	s.storageLog(p)
	return len(p), nil
}
