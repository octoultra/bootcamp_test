package store

import (
	"sync"
	"time"

	"entrytest/internal/message"
)

type MessageStore struct {
	mu       sync.RWMutex
	messages []message.Message
	nextID   int64
}

func New() *MessageStore {
	return &MessageStore{nextID: 1}
}

func (s *MessageStore) Create(text string) message.Message {
	s.mu.Lock()
	defer s.mu.Unlock()

	m := message.Message{
		ID:        s.nextID,
		Message:   text,
		CreatedAt: time.Now().UTC(),
	}
	s.nextID++
	s.messages = append(s.messages, m)
	return m
}

func (s *MessageStore) List() []message.Message {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]message.Message, len(s.messages))
	for i, m := range s.messages {
		out[len(s.messages)-1-i] = m
	}
	return out
}

func (s *MessageStore) Delete(id int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, m := range s.messages {
		if m.ID == id {
			out := make([]message.Message, 0, len(s.messages)-1)
			out = append(out, s.messages[:i]...)
			out = append(out, s.messages[i+1:]...)
			s.messages = out
			return true
		}
	}
	return false
}
