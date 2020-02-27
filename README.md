#Broad-mbta#

* Broad Institute MBTA challenge

* **category**    Application

## Description

The project is a golang standalone application. It uses MBTA public API without the API key to get datasets for subway Routes and their Stops.
Unfortunatelly because of time limitations (I had only a couple of evenings to solve the challenge) the application is solving Question 1 and 2 but not the Question 3. 
For Question3 the data input is ready, but the methods to build Routes graph and DFS algorithm to find the route path between 2 stops is in TODO state.

Features: 
- parallel http calls to get all MBTA stops by route
- http retry with delay in case of MBTA API throtteling
- minimal data model for MBTA

TODO:
- better exception handling. Errors should be more specific (ex. HTTPError, ParseError etc)
- Almost entire question 3 needs to be implemented 
- Test coverage should be better (closer to 100%)
- The console output is ugly 
- The list of stops in 3 columns isn't stright

## Requirements

To run the program you will need to install Golang v0.13 or Docker.

To install Golang just follow the instructions from here: https://golang.org/doc/install

If you don't want to install Golang, you can install docker as an alternative: https://docs.docker.com/install/

## Quick Start
1. Clone the project code from github: https://github.com/annasobo/broad-mbta.git
2a. If you have Golang installed:
- type ```make``` to run the application
- type ```make test``` to run unit tests with coverage

2b. If you have docker installed: 
- type ```docker-compose up``` - the command will execute golang docker container, as well as will test and run Broad-mbta application

3. To exit the application use command+C for MacOS (Ctrl+C for Linux and Windows)

## Sample output:

```Question 1 - retrieve all long names for Light and Heavy Rail.
Red Line, Green Line B, Mattapan Trolley, Blue Line, Green Line D, Green Line E, Orange Line, Green Line C

Question 2
Part 1 - subway route with the most stops
Route Green Line B has the most stops 24

Part 2 - subway route with the fewest stops
Route Mattapan Trolley has the fewest stops 8

Part 3 - List of stops that connect 2 or more subway routes
Downtown Crossing has multiple routes, the routes are: 
Red Line, Orange Line
Saint Paul Street has multiple routes, the routes are: 
Green Line C, Green Line B
Arlington has multiple routes, the routes are: 
Green Line C, Green Line B, Green Line D, Green Line E
Boylston has multiple routes, the routes are: 
Green Line C, Green Line B, Green Line D, Green Line E
Park Street has multiple routes, the routes are: 
Green Line C, Red Line, Green Line B, Green Line D, Green Line E
Hynes Convention Center has multiple routes, the routes are: 
Green Line C, Green Line B, Green Line D
North Station has multiple routes, the routes are: 
Green Line C, Green Line E, Orange Line
Ashmont has multiple routes, the routes are: 
Red Line, Mattapan Trolley
Copley has multiple routes, the routes are: 
Green Line C, Green Line B, Green Line D, Green Line E
State has multiple routes, the routes are: 
Blue Line, Orange Line
Kenmore has multiple routes, the routes are: 
Green Line C, Green Line B, Green Line D
Government Center has multiple routes, the routes are: 
Green Line C, Blue Line, Green Line D, Green Line E
Haymarket has multiple routes, the routes are: 
Green Line C, Green Line E, Orange Line
```



