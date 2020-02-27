package model

// TODO: What is missing is the implementation of differnet configs for various envioronments.
// Everything that can be configurable shoud go here

// RoutesPath is filtered by routeType, because we need only Light and Heavy Train (type 0 and 1) to do the exercise.
// To limit the ammount of data returned by the MBTA API I also specify what kind of fields should be returned (id, long_name)
const RoutesPath = "https://api-v3.mbta.com/routes?filter[type]=0,1&fields[route]=id,long_name"

// Bacause I didn't find any way to get routes and their stops in the same call, I use the MBTA stops endpoint to get all stops by route.
// The filter by route in case of stops endpoint is required by MBTA API.
const StopsPath = "https://api-v3.mbta.com/stops?filter[route]=%v&include=route&fields[name]"
