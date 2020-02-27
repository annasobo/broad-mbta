package model

import "testing"

import "encoding/json"

func TestJsonRoute(t *testing.T) {
	tests := []struct {
		jsonStr string
		err     error
	}{
		{
			jsonStr: RoutesResponse,
			err:     nil,
		},
	}
	var routes MbtaRoutes
	for _, test := range tests {
		err := json.Unmarshal([]byte(test.jsonStr), &routes)
		if err != test.err {
			t.Error(err.Error())
		}
	}
}

func TestJsonStops(t *testing.T) {
	tests := []struct {
		jsonStr string
		err     error
	}{
		{
			jsonStr: StopsResponse,
			err:     nil,
		},
	}
	var stops MbtaStop
	for _, test := range tests {
		err := json.Unmarshal([]byte(test.jsonStr), &stops)
		if err != test.err {
			t.Error(err.Error())
		}
	}
}
