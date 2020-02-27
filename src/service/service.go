package service

import (
	"errors"
	"strconv"
	"sync"

	"github.com/annasobo/broad-mbta/src/model"
)

// Data represents the structure used by channel to process requests in parallel.
// Channel is a very similar concept to Queue used in distributed systems.
// Every worker at the end will report its result to the channel. The result can be a Route or error
type Data struct {
	Route model.Route
	Err   error
}

// HTTPFunctions is for testing  purposes
type HTTPFunctions struct {
	GetHTTPRoutes func() ([]*model.MbtaRoute, error)
	GetHTTPStops  func(routeId string) ([]*model.MbtaStop, error)
}

// LoadData method connects to MBTA API, takes all MBTA subway routes together with their stops
func (f *HTTPFunctions) LoadData() (map[string]model.Route, error) {
	result := make(map[string]model.Route)
	// get routes from MBTA API
	routes, err := GetRoutes()
	if err != nil {
		return nil, err
	}
	// We want to make n different http calls in parallel, where n is a number of routes returned by the MBTA API
	channel := make(chan *Data, len(routes))
	// Close channel when work is done
	defer close(channel)
	// The WaitGroup will tell if all the workers are done with work, so we can process the result.
	// I could use a bool channel as well, but WaitGroup is a little faster
	var wg sync.WaitGroup
	for _, route := range routes {
		// The worker ready to go
		wg.Add(1)
		// The worker starts processing
		go f.loadStops(route.MbtaRouteToRoute(), channel, &wg)
	}
	// Wait until all workers are done
	wg.Wait()

	for i := 0; i < len(routes); i++ {
		// Load the first result into res variable
		res := <-channel
		// if the result has error, return
		if res.Err != nil {
			return nil, res.Err
		}
		// Load the stops by route into Route map
		result[res.Route.Id] = res.Route
	}
	return result, nil
}

// loadStops is a function executed as go routine. It's running in parallel for every route
func (*HTTPFunctions) loadStops(route model.Route, channel chan *Data, wg *sync.WaitGroup) {
	// Notify WaitGroup when done
	defer wg.Done()

	mbtaStops, err := GetStopsByRoute(route.Id)
	if err != nil {
		// Load error to the channel and return
		channel <- &Data{Err: err}
		return
	}
	// convert []MbtaStop result to internal []Stop model
	stops := make([]model.Stop, 0)
	for _, stop := range mbtaStops {
		stops = append(stops, stop.MbtaStopToStop())
	}
	route.Stops = stops
	channel <- &Data{Route: route, Err: nil}
	return
}

// LoadStopMap converts the Routes map data map[string]model.Route to stop map map[string]model.Stop
func LoadStopMap(data map[string]model.Route) map[string]model.Stop {
	stopMap := make(map[string]model.Stop)
	for keyR, valR := range data {
		for _, stop := range valR.Stops {
			if val, ok := stopMap[stop.Name]; ok {
				stopMap[stop.Name] = model.Stop{Name: stop.Name, RouteIds: append(val.RouteIds, keyR)}
			} else {
				stopMap[stop.Name] = model.Stop{Name: stop.Name, RouteIds: []string{keyR}}
			}
		}
	}
	return stopMap
}

// ValidStation validates the std input for question3
func ValidStation(input string, stations []string, graph map[string]model.Stop) (string, error) {
	result := ""
	if from, err := strconv.Atoi(input); err == nil && from >= 0 && from < len(stations) {
		result = stations[from]
	} else if _, ok := graph[input]; ok {
		result = input
	} else {
		return "", errors.New("Input is not correct")
	}
	return result, nil
}

// Maximum method Find maximum number of stops per Route
func Maximum(data map[string]model.Route) (routeName string, stopsCount int) {
	max := 0
	route := ""
	for _, val := range data {
		if len(val.Stops) > max {
			max = len(val.Stops)
			route = val.Name
		}
	}
	return route, max
}

// Minimum method Find minimum number of stops by route
func Minimum(data map[string]model.Route) (routeName string, stopsCount int) {
	min := int((^uint(0)) >> 1)
	route := ""
	for _, val := range data {
		if len(val.Stops) < min {
			min = len(val.Stops)
			route = val.Name
		}
	}
	return route, min
}

// StopWithMultipleRoutes Find all stops that have more than one route
func StopWithMultipleRoutes(data map[string]model.Route) map[string][]string {
	graph := LoadStopMap(data)
	result := make(map[string][]string)
	for _, val := range graph {
		if len(val.RouteIds) > 1 {
			result[val.Name] = val.RouteIds
		}
	}
	return result
}

// TODO: Because I'm running out of time, I didin't implement this method.
// The basic idea is to see the MBTA routes as a graph where 2 nodes are connected if there is a stop that connects them.
// Because we need the values of the nodes in path to answer question 3, we need an algorithm with backtracking.
// The perfect algorithm would be DFS
// DFS goes here:
func FindPath(from string, to string, graph *map[string]model.RouteNode) ([]string, error) {
	return nil, nil
}

// TODO: LoadGraph builds Route graph where route is a node and 2 routes are connected if there is stop that connects them.
func LoadGraph(graph map[string]model.Stop) map[string]model.RouteNode {
	return nil
}
