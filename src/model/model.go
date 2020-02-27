package model

// Internal model representing one MBTA stop
type Stop struct {
	Id       string
	Name     string
	RouteIds []string
}

// Internal model for MBTA route ex. Red Line, Blue Line
type Route struct {
	Id    string
	Name  string
	Stops []Stop
}

// Data structure for question 3. We can see routes as graph nodes connected with each other.
// Ex. Redline has connection with GreenLine (Park St) and OrangeLine (Downtown Crossing)
type RouteNode struct {
	Name        string
	RouteIds    []string
	Connections []*RouteNode
	Visited     bool
	Visiting    bool
}
