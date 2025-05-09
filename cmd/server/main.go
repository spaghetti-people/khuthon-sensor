package main

import "github.com/spaghetti-people/khuthon-sensor/api"

func main() {
	router := api.NewRouter()

	router.Run(":8082")
}
