package service

import "sync"

type RatingStore interface {
	Add(laptopID string, score float64) (*Rating, error)
}

type Rating struct {
	Count uint32
	Sum   float64
}

type InMemoryRatingStore struct {
	mutex  sync.RWMutex
	rating map[string]*Rating
}

func NewInMemoryRatingStore() *InMemoryRatingStore {
	return &InMemoryRatingStore{
		rating: make(map[string]*Rating),
	}
}

func (s *InMemoryRatingStore) Add(laptopID string, score float64) (*Rating, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	rating, inStore := s.rating[laptopID]
	if inStore {
		rating.Sum += score
		rating.Count++
	} else {
		rating = &Rating{
			Sum:   score,
			Count: 1,
		}
	}

	s.rating[laptopID] = rating

	return rating, nil
}