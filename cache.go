package gohelper

import (
	"sync"
	"sync/atomic"
)

type Cache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}

type MapCache struct {
	m map[string]interface{}
}

func NewMapCache() *MapCache {
	return &MapCache{m: make(map[string]interface{})}
}
func (M *MapCache) Get(key string) interface{} {
	if v, ok := M.m[key]; ok {
		return v
	}
	return nil
}

func (M *MapCache) Set(key string, value interface{}) {
	M.m[key] = value
}

type StructCache struct {
	lock sync.Mutex
	m    atomic.Value
}

func (s *StructCache) Get(key string) interface{} {
	return s.m.Load().(map[string]interface{})[key]
}

func (s *StructCache) Set(key string, value interface{}) {
	if s.m.Load() == nil {
		m := make(map[string]interface{})
		m[key] = value
		s.m.Store(m)
	} else {
		mm := s.m.Load().(map[string]interface{})
		nm := make(map[string]interface{}, len(mm)+1)
		for k, v := range mm {
			nm[k] = v
		}
		nm[key] = value
		s.m.Store(nm)
	}

}

type SyncMapCache struct {
	m sync.Map
}

func (s *SyncMapCache) Get(key string) interface{} {
	v, ok := s.m.Load(key)
	if ok {
		return v
	}
	return nil
}

func (s *SyncMapCache) Set(key string, value interface{}) {
	s.m.Store(key, value)
}
