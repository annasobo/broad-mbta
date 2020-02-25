package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/annasobo/broad-mbta/src/model"
)

func getRoutes() ([]*model.MbtaRoute, error) {
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

func GetStopsByRoute(routeId string) ([]*model.MbtaStop, error) {
	path := fmt.Sprintf(model.StopsPath, routeId)
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

func makeGetHttp(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode%100 > 3 {
		return nil, errors.New("Http response different than expected: " + string(resp.StatusCode))
	}
	return ioutil.ReadAll(resp.Body)
}
