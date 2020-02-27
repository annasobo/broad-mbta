package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/annasobo/broad-mbta/src/model"
)

// GetRoutes returns Routes from MBTA API
func GetRoutes() ([]*model.MbtaRoute, error) {
	data, err := makeGetHttp(model.RoutesPath)
	if err != nil {
		return nil, err
	}
	var routes model.MbtaRoutes
	err = json.Unmarshal(data, &routes)
	if err != nil {
		return nil, err
	}
	return routes.Data, nil
}

// GetStopsByRoute returns all stops of the given routeId in correct order
func GetStopsByRoute(routeID string) ([]*model.MbtaStop, error) {
	path := fmt.Sprintf(model.StopsPath, routeID)
	data, err := makeGetHttp(path)
	if err != nil {
		return nil, err
	}
	var stops model.MbtaStops
	err = json.Unmarshal(data, &stops)
	if err != nil {
		return nil, err
	}
	return stops.Data, nil
}

// TODO: Number of retries should go to config file
// makeGetHttp is a simple function to make http GET call to the API. It takes path of the API as an argument
// makeGetHttp returns a body from the respnse or error in case of failed http call
// Because there can be throtteling like in the MBTA case, where we are limited to 20 requests per sec,
// the method is waiting one second and retrying the call in case of errors
func makeGetHttp(path string) ([]byte, error) {
	retry := 2
	resp, err := http.Get(path)
	if err != nil && retry > 0 {
		retry -= 1
		time.Sleep(1000)
		resp, err = http.Get(path)
	}
	if err != nil && retry <= 0 {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode%100 > 3 {
		return nil, errors.New("Http response different than expected: " + string(resp.StatusCode))
	}
	return ioutil.ReadAll(resp.Body)
}
