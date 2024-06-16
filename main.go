package main

import (
	"myapp/database"
	"myapp/routes"
)

func main() {
	database.InitDatabase()
	r := routes.SetupRouter()
	r.Run(":8080") // Menjalankan server pada port 8080
}
