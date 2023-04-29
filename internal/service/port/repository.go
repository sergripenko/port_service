package port

import (
	"context"

	"github.com/sergripenko/port_service/internal/domain"
)

type RepositoryProvider interface {
	AddPort(ctx context.Context, port *domain.Port) (*domain.Port, error)
	GetPort(ctx context.Context, id string) (*domain.Port, error)
	UpdatePort(ctx context.Context, port *domain.Port) (*domain.Port, error)
}
