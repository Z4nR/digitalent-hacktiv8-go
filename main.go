package main

import (
	"mygram/database"
	"mygram/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":3000")
}
