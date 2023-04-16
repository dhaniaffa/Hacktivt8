package main

import (
	db "myGram/database"
	"myGram/routers"
)

func main() {
	db.StartDatabase()
	app := routers.ServeApp()
	app.Run(":8080")
}
