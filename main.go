package main

import (
	"fleetloom/database"
	"fleetloom/routes"
	"log"
	"net/http"

	_ "fleetloom/docs" // Import generated docs (Swag will create this package)
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger" // Import Swagger middleware
)

// @title fleetloom API
// @version 1.0
// @description This is a fleetloom server.
// @host localhost:8080
// @BasePath /

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	
	connectionString := "host=172.17.0.1 user=postgres password=R£Q7D0iz0hz5 dbname=fleetloom_db port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	//connectionString := "host=localhost user=postgres password=rohan dbname=fleetloom_db port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	database.InitDB(connectionString)

	router := mux.NewRouter()

	// Registering Swagger handler
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	routes.RegisterUserRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
