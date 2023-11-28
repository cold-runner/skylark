package cache

import (
	"time"
)

type memory struct {
	m map[string]interface{}
}

func (m *memory) Set(s string, s2 string) error {
	//TODO implement me
	panic("implement me")
}

func (m *memory) Get(s string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *memory) SetExpiration(s string, s2 string, duration time.Duration) error {
	//TODO implement me
	panic("implement me")
}
