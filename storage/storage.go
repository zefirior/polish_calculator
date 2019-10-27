package storage

import (
	"fmt"
	"strconv"
	"strings"
)

type Storage struct {
	data map[string]int
}

func NewStorage() *Storage {
	return &Storage{make(map[string]int)}
}

func (s *Storage) Set(key string, value int) {
	s.data[key] = value
}

func (s *Storage) Get(key string) int {
	if i, ok := s.data[key]; ok {
		return i
	}
	panic(fmt.Sprintf("Missing key %s in storage", key))
}

func FromString(variables string) *Storage {
	s := NewStorage()
	if len(variables) > 0 {
		for i, raw := range strings.Split(variables, " ") {
			key := string(int('a') + i)

			if value, err := strconv.Atoi(raw); err != nil {
				panic(fmt.Sprintf("Fail variable %s for storage", raw))
			} else {
				s.Set(key, value)
			}
		}
	}
	return s
}
