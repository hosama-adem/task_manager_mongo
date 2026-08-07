package main

import (
	"Task_manager/router"

	"Task_manager/data"
)

func main() {

	data.StartMongo()
	r := router.Router()

	r.Run()
}
