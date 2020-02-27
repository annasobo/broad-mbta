package model

// MBTA models used to parse json string from MBTA API
type MbtaRoutes struct {
	Data []*MbtaRoute `json:"data"`
}

type MbtaRoute struct {
	Attributes *MbtaRouteAttributes `json:"attributes,omitempty"`
	ID         string               `json:"id,omitempty"`
}

type MbtaRouteAttributes struct {
	LongName string `json:"long_name,omitempty"`
}

type MbtaStops struct {
	Data []*MbtaStop `json:"data"`
}

type MbtaStop struct {
	Attributes *MbtaStopAttributes `json:"attributes,omitempty"`
	ID         string              `json:"id,omitempty"`
}

type MbtaStopAttributes struct {
	Name string `json:"name,omitempty"`
}

// Converter function to confert Mbta Stop model into internal Stop model
func (m MbtaStop) MbtaStopToStop() Stop {
	return Stop{
		Id:   m.ID,
		Name: m.Attributes.Name,
	}
}

// Converter function to confert Mbta Route model into internal Route model
func (m MbtaRoute) MbtaRouteToRoute() Route {
	return Route{
		Id:    m.ID,
		Name:  m.Attributes.LongName,
		Stops: nil,
	}
}
