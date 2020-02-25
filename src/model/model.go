package model

import "sync"

type Stop struct {
	Id      string
	Name    string
	Visited bool // Metadata
}

type Route struct {
	Id    string
	Name  string
	Stops []Stop
	Mux   sync.Mutex
}

type MbtaRoute struct {
	Attributes *MbtaRouteAttributes `json:"attributes,omitempty"`
	ID         string               `json:"id,omitempty"`
}

type MbtaRouteAttributes struct {
	LongName string `json:"long_name,omitempty"`
}

type MbtaStop struct {
	Attributes *MbtaStopAttributes `json:"attributes,omitempty"`
	ID         string              `json:"id,omitempty"`
}

type MbtaStopAttributes struct {
	Name string `json:"name,omitempty"`
}

func (m MbtaStop) mbtaStopToStop() Stop {
	return Stop{
		Id:   m.ID,
		Name: m.Attributes.Name,
	}
}

func (m MbtaRoute) mbtaRouteToRoute(mbtaStops []MbtaStop) Route {
	var stops []Stop
	if mbtaStops != nil {

	}
	return Route{
		Id:    m.ID,
		Name:  m.Attributes.LongName,
		Stops: stops,
	}
}
