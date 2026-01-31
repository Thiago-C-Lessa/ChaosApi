package chaos

import (
	"sync"

	"github.com/google/uuid"
)

type InMemoryStore struct {
	sync.RWMutex
	data map[string]*Config
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]*Config),
	}
}

func (s *InMemoryStore) Find(path, method string) (*Config, bool) {
	s.RLock()
	defer s.RUnlock()
	for _, cfg := range s.data {
		if cfg.Path == path && cfg.Method == method {
			return cfg, true
		}
	}
	return nil, false
}

func (s *InMemoryStore) Create(cfg *Config) string {
	s.Lock()
	defer s.Unlock()
	id := uuid.NewString()
	cfg.ID = id
	s.data[id] = cfg
	return id
}

func (s *InMemoryStore) List() []*Config {
	s.RLock()
	defer s.RUnlock()
	result := make([]*Config, 0, len(s.data))
	for _, cfg := range s.data {
		result = append(result, cfg)
	}
	return result
}

func (s *InMemoryStore) Delete(id string) bool {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.data[id]; ok {
		delete(s.data, id)
		return true
	}
	return false
}
