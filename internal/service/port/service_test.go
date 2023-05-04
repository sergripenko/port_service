package port

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/sergripenko/port_service/internal/repository/mem"
)

func TestService_AddPorts(t *testing.T) {
	repo := mem.NewRepository()
	portService := NewService(repo)

	testCases := []struct {
		title string
		input []byte
		err   error
	}{
		{
			title: "empty file",
			input: []byte(``),
			err:   io.EOF,
		},
		{
			title: "correct case",
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
			err: nil,
		},
		{
			title: "update port",
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
  			},
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
    		"code": "52001"
  			}}`),
			err: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()
			err := portService.AddPorts(context.TODO(), bytes.NewBuffer(tc.input))
			if err != tc.err {
				t.Fatal(err)
			}
		})
	}
}
