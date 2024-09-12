package routes

import (
	"fleetloom/controllers"  // Importing the controllers package to use the handler functions
	"github.com/gorilla/mux" // Importing the Gorilla Mux package for routing HTTP requests
)

// RegisterUserRoutes registers the routes for the User entity.
func RegisterUserRoutes(router *mux.Router) {
	// Mapping the POST request to the CreateUser function.
	// When a POST request is made to /users, it calls CreateUser.
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	// Mapping the GET request to the GetUser function.
	// When a GET request is made to /users/{id}, it calls GetUser.
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")

	// Mapping the PUT request to the UpdateUser function.
	// When a PUT request is made to /users/{id}, it calls UpdateUser.
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")

	// Mapping the DELETE request to the DeleteUser function.
	// When a DELETE request is made to /users/{id}, it calls DeleteUser.
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// Mapping the GET request to the GetUserByEmail function.
	// When a GET request is made to /users/email/{email}, it calls GetUserByEmail.
	router.HandleFunc("/users/email/{email}", controllers.GetUserByEmail).Methods("GET")

	// Mapping the GET request to the GetUserByPhone function.
	// When a GET request is made to /users/phone/{phone}, it calls GetUserByPhone.
	router.HandleFunc("/users/phone/{phone}", controllers.GetUserByPhone).Methods("GET")

	// Mapping the GET request to the GetUserByName function.
	// When a GET request is made to /users/name/{name}, it calls GetUserByName.
	router.HandleFunc("/users/name/{name}", controllers.GetUserByName).Methods("GET")

}
