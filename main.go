package main

import (
	"finalProject1/database"
	routers "finalProject1/routers"
	"fmt"
)

func main() {
	fmt.Println("Final Project 1")
	database.StartDB()
	routers.Router()
}
