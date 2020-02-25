package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getRoutes() ([]MbtaRoute, error) {
	data, err := makeGetHttp(RoutesPath)
	if err != nil {
		return nil, err
	}
	routes := make([]MbtaRoute, 0)
	err = json.Unmarshal(data, &routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func getStopsByRoute(routeId string) ([]MbtaStop, error) {
	path := fmt.Sprintf(StopsPath, routeId)
	data, err := makeGetHttp(path)
	if err != nil {
		return nil, err
	}
	stops := make([]MbtaStop, 0)
	err = json.Unmarshal(data, &stops)
	if err != nil {
		return nil, err
	}
	return stops, nil
}

func makeGetHttp(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
