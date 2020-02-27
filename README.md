# ~#PROJECT#~

* Broad Institute MBTA challenge *

* **category**    Application

## Description

The project a golang standalone application. It uses MBTA public API without the API key to get datasets for subway Routes and their Stops.
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

To run the program you will need to install golang or Docker.

To install golang just follow the instructions from here: https://golang.org/doc/install

If you don't want to install golang, you can install docker as an alternative: https://docs.docker.com/install/

## Quick Start
- Clone the project code from github: https://github.com/annasobo/broad-mbta.git
- If you have Golang installed:
-- type ```make``` to run the application
-- type ```make test``` to run unit tests with coverage

- If you have docker installed: 
-- type ```docker-compose up``` - the command will execute golang docker container, as well as will test and run Broad-mbta application





