package port

import (
	"context"
	"io"
)

type ServiceProvider interface {
	AddPorts(ctx context.Context, reader io.Reader) error
	Shutdown()
}
