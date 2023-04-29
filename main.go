package main

import (
	"DTSGolang/Kelas2/Assignment8/database"
	"DTSGolang/Kelas2/Assignment8/routers"
)

var PORT = ":8000"

func main() {
	database.StartDB()
	routers.StartServer().Run(PORT)
}
