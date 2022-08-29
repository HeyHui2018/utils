package utils

import (
	"fmt"
	"sync"
	"time"
)

type ReadBlockMap struct {
	m map[string]*entry
	l sync.RWMutex
}

type entry struct {
	c       chan struct{}
	value   interface{}
	isExist bool
}

func (s ReadBlockMap) write(key string, value interface{}) {
	s.l.Lock()
	defer s.l.Unlock()

	e, ok := s.m[key]
	if !ok {
		s.m[key] = &entry{
			value:   value,
			isExist: true,
		}
		return
	}
	e.value = value
	if !e.isExist {
		if e.c != nil {
			close(e.c)
			e.c = nil
		}
	}
}

func (s ReadBlockMap) read(key string, timeout time.Duration) interface{} {
	s.l.RLock()

	e, ok := s.m[key]
	if ok {
		if e.isExist {
			s.l.RUnlock()
			return e.value
		}
		s.l.RUnlock()
		select {
		case <-time.After(timeout):
			return fmt.Errorf("timeout")
		case <-e.c:
			return e.value
		}
	}
	s.l.RUnlock()
	s.l.Lock()
	s.m[key] = &entry{
		c:       make(chan struct{}),
		isExist: false,
	}
	s.l.Unlock()
	select {
	case <-time.After(timeout):
		return fmt.Errorf("timeout")
	case <-e.c:
		return e.value
	}
}
