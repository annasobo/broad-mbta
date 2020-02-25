package main

const RoutesPath = "https://api-v3.mbta.com/routes?filter[type]=0,1&fields[route]=id,long_name"
const StopsPath = "https://api-v3.mbta.com/stops?filter[route]=%v&include=route&fields[name]"
