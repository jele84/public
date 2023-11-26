package main

import (
	"errors"
	"sync"
)

type Employee struct {
	Id   int
	Name string
}

type Storage interface {
	GetEmployee(id int) (Employee, error)
	Insert(e *Employee)
}

type MemoryStorage struct {
	data    map[int]Employee
	counter int
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Employee),
		// counter: 1,
	}
}

func (s *MemoryStorage) GetEmployee(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()
	employee, ok := s.data[id]
	if !ok {
		return employee, errors.New("user not found")
	}
	return employee, nil
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()
	e.Id = s.counter
	s.data[e.Id] = *e
	s.counter++
	s.Unlock()
}
