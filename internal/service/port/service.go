package port

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/pkg/errors"

	"github.com/sergripenko/port_service/internal/domain"
	"github.com/sergripenko/port_service/internal/repository"
)

type Service struct {
	wg   sync.WaitGroup
	repo RepositoryProvider
}

func NewService(repo RepositoryProvider) *Service {
	return &Service{
		wg:   sync.WaitGroup{},
		repo: repo,
	}
}

func (s *Service) AddPorts(ctx context.Context, reader io.Reader) error {
	// Increment wait group for graceful shutdown.
	s.wg.Add(1)
	defer s.wg.Done()

	r := bufio.NewReader(reader)
	decoder := json.NewDecoder(r)

	// Expect start of object as the first token.
	t, err := decoder.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("expected {, got %v", t)
	}

	type portData struct {
		Name        string    `json:"name"`
		City        string    `json:"city"`
		Country     string    `json:"country"`
		Alias       []string  `json:"alias"`
		Regions     []string  `json:"regions"`
		Coordinates []float64 `json:"coordinates"`
		Province    string    `json:"province"`
		Timezone    string    `json:"timezone"`
		Unlocs      []string  `json:"unlocs"`
		Code        string    `json:"code"`
	}
	var i int

	for decoder.More() {
		// Read the key.
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		portID, ok := t.(string) // type assert token to string.
		if !ok {
			continue
		}

		var elm portData
		err = decoder.Decode(&elm)
		if err != nil {
			return err
		}

		// Form new port object.
		port := &domain.Port{
			ID:          portID,
			Name:        elm.Name,
			City:        elm.City,
			Country:     elm.Country,
			Alias:       elm.Alias,
			Regions:     elm.Regions,
			Coordinates: elm.Coordinates,
			Province:    elm.Province,
			Timezone:    elm.Timezone,
			Unlocs:      elm.Unlocs,
			Code:        elm.Code,
		}
		i++
		_, err = s.repo.GetPort(ctx, portID)
		if err != nil {
			if errors.Is(err, repository.ErrRecordNotFount) {
				// Add new port to repo.
				if _, err = s.repo.AddPort(ctx, port); err != nil {
					return err
				}
				continue
			} else {
				return err
			}
		}
		if _, err = s.repo.UpdatePort(ctx, port); err != nil {
			return err
		}
	}
	if _, err = decoder.Token(); err != nil {
		return err
	}

	log.Printf("new %d ports processsed", i)
	return nil
}

func (s *Service) Shutdown() {
	s.wg.Wait()
}
