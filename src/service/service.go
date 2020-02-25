package service

import (
	"fmt"
	"sync"

	"github.com/annasobo/broad-mbta/src/model"
)

type Data struct {
	Route model.Route
	Err   error
}

func LoadData() (map[string]model.Route, error) {
	result := make(map[string]model.Route)
	routes, err := getRoutes()
	if err != nil {
		return nil, err
	}
	channel := make(chan *Data, len(routes))
	var wg sync.WaitGroup
	for _, route := range routes {
		wg.Add(1)
		println("Adding: " + route.ID)
		go loadStops(route.MbtaRouteToRoute(), channel, &wg)
	}
	wg.Wait()
	for i := 0; i < len(channel); i++ {
		res := <-channel
		if res.Err != nil {
			return nil, res.Err
		}
		result[res.Route.Id] = res.Route
	}
	return result, nil
}

func loadStops(route model.Route, channel chan *Data, wg *sync.WaitGroup) {
	defer wg.Done()

	mbtaStops, err := GetStopsByRoute(route.Id)
	if err != nil {
		channel <- &Data{Err: err}
		return
	}
	stops := make([]model.Stop, 0)
	for _, stop := range mbtaStops {
		stops = append(stops, stop.MbtaStopToStop())
	}
	route.Stops = stops
	channel <- &Data{Route: route, Err: nil}
	println(fmt.Sprintf("%v: %v", route.Name, len(stops)))
	return
}
