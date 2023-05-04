package mem

import (
	"sync"

	"github.com/sergripenko/port_service/internal/domain"
)

type Repository struct {
	mut   sync.RWMutex // mutex lock
	ports map[string]*domain.Port
}

func NewRepository() *Repository {
	return &Repository{
		mut:   sync.RWMutex{},
		ports: map[string]*domain.Port{},
	}
}
