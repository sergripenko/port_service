package port

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sergripenko/port_service/internal/domain"
	"github.com/sergripenko/port_service/internal/repository"
)

var mockErr = errors.New("mock error")

func TestService_AddPorts(t *testing.T) {

	testCases := []struct {
		title         string
		input         []byte
		mockRepoCalls func(provider *MockRepositoryProvider)
		err           error
	}{
		{
			title: "not JSON object format",
			input: []byte(`[]`),
			err:   mockErr,
		},
		{
			title: "empty file",
			input: []byte(``),
			err:   io.EOF,
		},
		{
			title: "error repo GetPort",
			input: []byte(`{
			"AEAJM": {
			"name": "Ajman",
			"city": "Ajman",
			"country": "United Arab Emirates",
			"alias": [],
			"regions": [],
			"coordinates": [
				55.5136433,
				25.4052165
			],
			"province": "Ajman",
			"timezone": "Asia/Dubai",
			"unlocs": [
			"AEAJM"
			],
			"code": "52000"
			}}`),
			mockRepoCalls: func(provider *MockRepositoryProvider) {
				provider.On("GetPort", mock.Anything, mock.Anything).Return(
					func(ctx context.Context, id string) *domain.Port {
						return nil
					},
					func(ctx context.Context, id string) error {
						return mockErr
					}).Once()
			},
			err: mockErr,
		},
		{
			title: "add new 1 port",
			input: []byte(`{
			"AEAJM": {
			"name": "Ajman",
			"city": "Ajman",
			"country": "United Arab Emirates",
			"alias": [],
			"regions": [],
			"coordinates": [
				55.5136433,
				25.4052165
			],
			"province": "Ajman",
			"timezone": "Asia/Dubai",
			"unlocs": [
			"AEAJM"
			],
			"code": "52000"
			}}`),
			mockRepoCalls: func(provider *MockRepositoryProvider) {
				provider.On("GetPort", mock.Anything, mock.Anything).Return(
					func(ctx context.Context, id string) *domain.Port {
						return nil
					},
					func(ctx context.Context, id string) error {
						return repository.ErrRecordNotFount
					}).Once()
				provider.On("AddPort", mock.Anything, mock.Anything).Return(
					func(ctx context.Context, port *domain.Port) *domain.Port {
						return nil
					},
					func(ctx context.Context, port *domain.Port) error {
						return mockErr
					}).Once()
			},
			err: nil,
		},
		{
			title: "add new 2 ports",
			input: []byte(`{
			"AEAJM": {
    		"name": "Ajman",
    		"city": "Ajman",
    		"country": "United Arab Emirates",
    		"alias": [],
    		"regions": [],
    		"coordinates": [
      			55.5136433,
      			25.4052165
    		],
    		"province": "Ajman",
    		"timezone": "Asia/Dubai",
    		"unlocs": [
      		"AEAJM"
    		],
    		"code": "52000"
  			},
			"AEAUH": {
			"name": "Abu Dhabi",
			"coordinates": [
			  54.37,
			  24.47
			],
			"city": "Abu Dhabi",
			"province": "Abu Z¸aby [Abu Dhabi]",
			"country": "United Arab Emirates",
			"alias": [],
			"regions": [],
			"timezone": "Asia/Dubai",
			"unlocs": [
			  "AEAUH"
			],
			"code": "52001"
		  }}`),
			mockRepoCalls: func(provider *MockRepositoryProvider) {
				provider.On("GetPort", mock.Anything, mock.Anything).Return(
					func(ctx context.Context, id string) *domain.Port {
						return nil
					},
					func(ctx context.Context, id string) error {
						return repository.ErrRecordNotFount
					}).Twice()
				provider.On("AddPort", mock.Anything, mock.Anything).Return(
					func(ctx context.Context, port *domain.Port) *domain.Port {
						return &domain.Port{}
					},
					func(ctx context.Context, port *domain.Port) error {
						return nil
					}).Twice()
			},
			err: nil,
		},
		{
			title: "error repo Update port",
			input: []byte(`{
			"AEAUH": {
			"name": "Ajman",
			"city": "Ajman",
			"country": "United Arab Emirates",
			"alias": [],
			"regions": [],
			"coordinates": [
				55.5136433,
				25.4052165
			],
			"province": "Ajman",
			"timezone": "Asia/Dubai",
			"unlocs": [
			"AEAJM"
			],
			"code": "52000"
			}}`),
			mockRepoCalls: func(provider *MockRepositoryProvider) {
				provider.On("GetPort", mock.Anything, mock.Anything).Return(
					func(ctx context.Context, id string) *domain.Port {
						return &domain.Port{}
					},
					func(ctx context.Context, id string) error {
						return nil
					}).Once()
				provider.On("UpdatePort", mock.Anything, mock.Anything).Return(
					func(ctx context.Context, port *domain.Port) *domain.Port {
						return nil
					},
					func(ctx context.Context, port *domain.Port) error {
						return mockErr
					}).Once()
			},
			err: mockErr,
		},
	}

	for _, tc := range testCases {
		repoMock := &MockRepositoryProvider{}
		tc := tc

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()
			portService := NewService(repoMock)
			if tc.mockRepoCalls != nil {
				tc.mockRepoCalls(repoMock)
			}
			err := portService.AddPorts(context.TODO(), bytes.NewBuffer(tc.input))
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
