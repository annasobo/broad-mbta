package model

type Stop struct {
	Id      string
	Name    string
	Visited bool // Metadata
}

type Route struct {
	Id    string
	Name  string
	Stops []Stop
}

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

func (m MbtaStop) MbtaStopToStop() Stop {
	return Stop{
		Id:   m.ID,
		Name: m.Attributes.Name,
	}
}

func (m MbtaRoute) MbtaRouteToRoute() Route {
	return Route{
		Id:    m.ID,
		Name:  m.Attributes.LongName,
		Stops: nil,
	}
}
