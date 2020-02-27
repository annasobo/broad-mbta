package service

import (
	"encoding/json"
	"testing"

	"github.com/annasobo/broad-mbta/src/model"
	"gotest.tools/assert"
)

func GetMockRoutes() ([]*model.MbtaRoute, error) {
	var routes *model.MbtaRoutes
	err := json.Unmarshal([]byte(model.RoutesResponse), &routes)
	if err != nil {
		return nil, err
	}

	return routes.Data, nil
}

func GetMockStops(routeId string) ([]*model.MbtaStop, error) {
	var stops *model.MbtaStops
	err := json.Unmarshal([]byte(model.StopsResponse), &stops)
	if err != nil {
		return nil, err
	}
	return stops.Data, nil
}

// Test should be parametrized and check all edge cases
func TestService(t *testing.T) {
	mockHTTP := HTTPFunctions{GetHTTPRoutes: GetMockRoutes, GetHTTPStops: GetMockStops}
	data, err := mockHTTP.LoadData()
	if err != nil {
		t.Error(err.Error())
	}
	r,s := Maximum(data)
	assert.Equal(t, r, "Green Line B")
	assert.Equal(t, s, 24)
	r,s = Minimum(data)
	assert.Equal(t, r, "Mattapan Trolley")
	assert.Equal(t, s, 8)
	stopMap := StopWithMultipleRoutes(data)
	assert.Equal(t, len(stopMap), 13)
}
