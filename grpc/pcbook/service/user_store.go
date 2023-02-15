package service

import (
	"errors"
	"sync"
)

var ErrorUserAlreadyExists = errors.New("user already exists")
var ErrorUserNotFound = errors.New("user not found")

type UserStore interface {
	Save(user *User) error
	Find(username string) (*User, error)
}

type InMemoryUserStore struct {
	mutex sync.Mutex
	users map[string]*User
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*User),
	}
}

func (s *InMemoryUserStore) Save(user *User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, inStore := s.users[user.Username]; inStore {
		return ErrorUserAlreadyExists
	}

	s.users[user.Username] = user.Clone()

	return nil
}

func (s *InMemoryUserStore) Find(username string) (*User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	user, inStore := s.users[username]
	if !inStore {
		return nil, ErrorUserNotFound
	}

	return user, nil
}
