package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/annasobo/broad-mbta/src/model"
	"github.com/annasobo/broad-mbta/src/service"
)

func main() {
	data, err := service.LoadData()
	if err != nil {
		println(err.Error())
	}
	PrintQuestion1(data)
	PrintQuestion2(data)
	PrintQuestion3(data)
}

func PrintQuestion1(data map[string]model.Route) {
	println("Question 1 - retrieve all long names for Light and Heavy Rail.")
	idx := 0
	length := len(data)
	for _, val := range data {
		idx += 1
		print(val.Name)
		// don't print coma at the end
		if idx <= length-1 {
			print(", ")
		}
	}
	println()
	println()
}

func PrintQuestion2(data map[string]model.Route) {
	println("Question 2")
	println("Part 1 - subway route with the most stops")
	r, c := service.Maximum(data)
	println(fmt.Sprintf("Route %v has the most stops %v", r, c))
	println()
	println("Part 2 - subway route with the fewest stops")
	r, c = service.Minimum(data)
	println(fmt.Sprintf("Route %v has the fewest stops %v", r, c))
	println()
	println("Part 3 - List of stops that connect 2 or more subway routes")
	routes := service.StopWithMultipleRoutes(data)
	for key, val := range routes {
		println(fmt.Sprintf("%v has multiple routes, the routes are: ", key))
		for idx, routeId := range val {
			print(data[routeId].Name)
			if idx < len(val)-1 {
				print(", ")
			}
		}
		println()
	}
	println()
}

func PrintQuestion3(data map[string]model.Route) {
	graph := service.LoadStopMap(data)
	stations := make([]string, 0)
	for idx, _ := range graph {
		stations = append(stations, idx)
	}
	sort.Strings(stations)
	println("From the list below choose 2 stops (type its number or name)")
	for idx, val := range stations {
		print(fmt.Sprintf("%v - %v\t\t", idx, val))
		if (idx+1)%3 == 0 {
			println()
		}
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number or name of the first stop: ")
	fromIn, _ := reader.ReadString('\n')
	from := ""
	to := ""
	var err error
	for from, err = service.ValidStation(strings.Trim(fromIn, "\n "), stations, graph); err != nil; {
		println("The station name or number is incorrect, enter valid value: ")
		fromIn, _ = reader.ReadString('\n')
	}
	println("Enter the number or name of the second stop: ")
	toIn, _ := reader.ReadString('\n')
	for to, err = service.ValidStation(strings.Trim(toIn, "\n "), stations, graph); err != nil; {
		fmt.Println("The station name or number is incorrect, enter valid value: ")
		toIn, _ = reader.ReadString('\n')
	}
	fmt.Printf("You choose %v and %v", from, to)
	fmt.Println()
	//calculateRoute(from, to, &graph)
}
