package job

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Store struct {
	mu   sync.RWMutex
	jobs map[string]*Job
}

func NewStore() *Store {
	return &Store{
		jobs: make(map[string]*Job),
	}
}

func (s *Store) Create() *Job {
	job := &Job{
		ID:        uuid.NewString(),
		Status:    StatusQueued,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mu.Lock()
	s.jobs[job.ID] = job
	s.mu.Unlock()

	return job
}

func (s *Store) Update(job *Job) {
	job.UpdatedAt = time.Now()
}

func (s *Store) Get(id string) (*Job, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.jobs[id]
	return job, ok
}
