package main

import (
	db "myGram/database"
	"myGram/routers"
	"os"
)

func main() {
	db.StartDatabase()
	var PORT = os.Getenv("PORT")
	app := routers.ServeApp()
	app.Run(":" + PORT)
}
