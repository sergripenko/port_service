package mem

import (
	"context"

	"github.com/sergripenko/port_service/internal/domain"
	"github.com/sergripenko/port_service/internal/service/port/repository"
)

func (r *Repository) AddPort(ctx context.Context, port *domain.Port) (*domain.Port, error) {
	r.mut.Lock()
	defer r.mut.Unlock()

	r.ports[port.ID] = port
	return port, nil
}

func (r *Repository) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	r.mut.RLock()
	defer r.mut.RUnlock()

	port, exist := r.ports[id]
	if !exist {
		return nil, repository.ErrRecordNotFount
	}
	return port, nil
}

func (r *Repository) UpdatePort(ctx context.Context, port *domain.Port) (*domain.Port, error) {
	r.mut.Lock()
	defer r.mut.Unlock()

	r.ports[port.ID] = port
	return port, nil
}
