package sessions

import (
	"fmt"
	"sync"
	"time"
)

type Storage interface {
	Read(s *Session) error
	Write(s *Session) error
	Remove(s *Session) error
}

// MemoryStorage 將session儲存到內存中
type MemoryStorage struct {
	// tm 超時管理器
	tm map[string]*time.Timer
	//RW鎖防止修改map時發生併發問題
	rw      sync.RWMutex
	store   map[string]*Session
	gcTruck chan string
}

func NewMemoryStorage() *MemoryStorage {
	m := &MemoryStorage{
		tm:      make(map[string]*time.Timer, 1024),
		rw:      sync.RWMutex{},
		store:   make(map[string]*Session),
		gcTruck: make(chan string, 1024),
	}
	go m.gc()
	return m
}

func (m *MemoryStorage) Read(s *Session) (err error) {
	m.rw.RLock()
	defer m.rw.RUnlock()
	if session, ok := m.store[s.id]; ok {
		s.Values = session.Values
		s.CreateAt = session.CreateAt
		s.ExpireAt = session.ExpireAt
		return nil
	}
	return fmt.Errorf("read Memory session error %v", err)
}

func (m *MemoryStorage) Write(s *Session) (err error) {
	m.rw.Lock()
	defer m.rw.Unlock()
	m.store[s.id] = s
	if m.store[s.id] == nil {
		go func() {
			m.tm[s.id] = time.NewTimer(time.Until(s.ExpireAt))
			<-m.tm[s.id].C
			m.gcTruck <- s.id
			m.tm[s.id].Stop()
		}()
	}
	return nil
}

func (m *MemoryStorage) Remove(s *Session) (err error) {
	m.rw.Lock()
	defer m.rw.Unlock()
	delete(m.store, s.id)
	delete(m.tm, s.id)
	return nil
}

func (m *MemoryStorage) gc() {
	for {
		select {
		case sid := <-m.gcTruck:
			m.rw.Lock()
			delete(m.store, sid)
			m.rw.RUnlock()
		default:
			fmt.Println("Go running....")
		}
	}
}
