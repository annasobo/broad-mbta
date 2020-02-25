package main

import (
	"encoding/json"

	"github.com/annasobo/broad-mbta/src/service"
)

func main() {
	data, err := service.LoadData()
	if err != nil {
		println(err.Error())
	}
	empData, err := json.Marshal(data)
	if err != nil {
		println(err.Error())
		return
	}
	println(string(empData))
}
