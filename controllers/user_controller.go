package controllers

import (
	"database/sql"  // Importing the database/sql package to interact with the SQL database
	"encoding/json" // Importing the encoding/json package to handle JSON data
	"log"           // Importing the log package to log errors or information
	"net/http"      // Importing the net/http package to handle HTTP requests and responses
	"strconv"       // Importing the strconv package to convert strings to other types like int
    "strings"       // Importing the string package to manipulate and inspect strings
	"fleetloom/database" // Importing the database package to access the database connection
	"fleetloom/models"   // Importing the models package to use the User struct
	"github.com/gorilla/mux" // Importing the Gorilla Mux package for routing HTTP requests
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body models.User true "User Info"
// @Success 200 {object} models.User
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	sqlStatement := `
    INSERT INTO users (first_name, last_name, current_address, permanent_address, pincode, city, state, contact_number, email, password, emergency_first_name, emergency_last_name, aadhar_card_number, driving_license_number, pan_card_number, branch, user_role)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
    RETURNING id`

	err := database.DB.QueryRow(sqlStatement, user.FirstName, user.LastName, user.CurrentAddress, user.PermanentAddress, user.Pincode, user.City, user.State, user.ContactNumber, user.Email, user.Password, user.EmergencyFirstName, user.EmergencyLastName, user.AadharCardNumber, user.DrivingLicenseNumber, user.PanCardNumber, user.Branch, user.UserRole).Scan(&user.ID)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(user)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Retrieve details of a specific user by their ID
// @Tags User
// @Param id path int true "User ID"
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Retrieving the `id` parameter from the URL path using Gorilla Mux.
	params := mux.Vars(r)

	// Converting the `id` parameter from a string to an integer.
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest) // Return a 400 Bad Request if the ID is not valid
		return
	}

	var user models.User // Declaring a variable `user` of type `models.User` to hold the retrieved user data

	// SQL statement to select a user by ID from the `users` table.
	sqlStatement := `SELECT * FROM users WHERE id=$1`

	// Executing the SQL statement to retrieve the user.
	// The `DB.QueryRow` function executes a query that is expected to return at most one row.
	// It scans the returned row into the fields of the `user` struct.
	row := database.DB.QueryRow(sqlStatement, id)
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CurrentAddress, &user.PermanentAddress, &user.Pincode, &user.City, &user.State, &user.ContactNumber, &user.Email, &user.Password, &user.EmergencyFirstName, &user.EmergencyLastName, &user.AadharCardNumber, &user.DrivingLicenseNumber, &user.PanCardNumber, &user.Branch, &user.UserRole ,&user.CreatedAt)

	// Checking if there was an error while scanning the row.
	if err != nil {
		// If no rows were found, return a 404 Not Found error.
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			log.Fatal(err) // For any other error, log it and stop the program.
		}
		return
	}

	// Encode the retrieved `user` struct into JSON and write it to the response.
	json.NewEncoder(w).Encode(user)
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update an existing user's details by their ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "Updated User Info"
// @Success 200 {string} string "User updated successfully"
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Retrieving the `id` parameter from the URL path using Gorilla Mux.
	params := mux.Vars(r)

	// Converting the `id` parameter from a string to an integer.
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest) // Return a 400 Bad Request if the ID is not valid
		return
	}

	var user models.User // Declaring a variable `user` of type `models.User` to hold the incoming updated user data

	// Decoding the JSON request body into the `user` struct.
	json.NewDecoder(r.Body).Decode(&user)

	// SQL statement to update the user's details in the `users` table.
	sqlStatement := `
    UPDATE users 
    SET first_name=$1, last_name=$2, current_address=$3, permanent_address=$4, pincode=$5, city=$6, state=$7, contact_number=$8, email=$9, password=$10, emergency_first_name=$11, emergency_last_name=$12, aadhar_card_number=$13, driving_license_number=$14, pan_card_number=$15, branch=$16, user_role=$17 
    WHERE id=$18`

	// Executing the SQL statement with the data from the `user` struct.
	// `DB.Exec` is used for queries that don't return rows, such as an UPDATE.
	_, err = database.DB.Exec(sqlStatement, user.FirstName, user.LastName, user.CurrentAddress, user.PermanentAddress, user.Pincode, user.City, user.State, user.ContactNumber, user.Email, user.Password, user.EmergencyFirstName, user.EmergencyLastName, user.AadharCardNumber, user.DrivingLicenseNumber, user.PanCardNumber, user.Branch, id)

	// If there is an error executing the query, log it and stop the program.
	if err != nil {
		log.Fatal(err)
	}

	// Write a success message to the response, indicating the user was updated.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User updated successfully")
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Remove a user from the system by their ID
// @Tags User
// @Param id path int true "User ID"
// @Produce  json
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Retrieving the `id` parameter from the URL path using Gorilla Mux.
	params := mux.Vars(r)

	// Converting the `id` parameter from a string to an integer.
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest) // Return a 400 Bad Request if the ID is not valid
		return
	}

	// SQL statement to delete a user by ID from the `users` table.
	sqlStatement := `DELETE FROM users WHERE id=$1`

	// Executing the SQL statement to delete the user.
	// `DB.Exec` is used for queries that don't return rows, such as DELETE.
	_, err = database.DB.Exec(sqlStatement, id)

	// If there is an error executing the query, log it and stop the program.
	if err != nil {
		log.Fatal(err)
	}

	// Write a success message to the response, indicating the user was deleted.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User deleted successfully")
}

// GetUserByEmail godoc
// @Summary Get a user by Email
// @Description Retrieve details of a specific user by their Email
// @Tags User
// @Param email path string true "User Email"
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {string} string "Invalid email"
// @Failure 404 {string} string "User not found"
// @Router /users/email/{email} [get]
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	// Retrieving the `email` parameter from the URL path using Gorilla Mux.
	params := mux.Vars(r)
	email := params["email"]

	var user models.User // Declaring a variable `user` of type `models.User` to hold the retrieved user data

	// SQL statement to select a user by email from the `users` table.
	sqlStatement := `SELECT * FROM users WHERE email=$1`

	// Executing the SQL statement to retrieve the user.
	row := database.DB.QueryRow(sqlStatement, email)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CurrentAddress, &user.PermanentAddress, &user.Pincode, &user.City, &user.State, &user.ContactNumber, &user.Email, &user.Password, &user.EmergencyFirstName, &user.EmergencyLastName, &user.AadharCardNumber, &user.DrivingLicenseNumber, &user.PanCardNumber, &user.Branch, &user.UserRole, &user.CreatedAt)

	// Checking if there was an error while scanning the row.
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			log.Fatal(err)
		}
		return
	}

	// Encode the retrieved `user` struct into JSON and write it to the response.
	json.NewEncoder(w).Encode(user)
}


// GetUserByPhone godoc
// @Summary Get a user by Phone Number
// @Description Retrieve details of a specific user by their Phone Number
// @Tags User
// @Param phone path string true "User Phone Number"
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {string} string "Invalid phone number"
// @Failure 404 {string} string "User not found"
// @Router /users/phone/{phone} [get]
func GetUserByPhone(w http.ResponseWriter, r *http.Request) {
	// Retrieving the `phone` parameter from the URL path using Gorilla Mux.
	params := mux.Vars(r)
	phone := params["phone"]

	var user models.User // Declaring a variable `user` of type `models.User` to hold the retrieved user data

	// SQL statement to select a user by phone number from the `users` table.
	sqlStatement := `SELECT * FROM users WHERE contact_number=$1`

	// Executing the SQL statement to retrieve the user.
	row := database.DB.QueryRow(sqlStatement, phone)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CurrentAddress, &user.PermanentAddress, &user.Pincode, &user.City, &user.State, &user.ContactNumber, &user.Email, &user.Password, &user.EmergencyFirstName, &user.EmergencyLastName, &user.AadharCardNumber, &user.DrivingLicenseNumber, &user.PanCardNumber, &user.Branch, &user.UserRole, &user.CreatedAt)

	// Checking if there was an error while scanning the row.
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			log.Fatal(err)
		}
		return
	}

	// Encode the retrieved `user` struct into JSON and write it to the response.
	json.NewEncoder(w).Encode(user)
}

// GetUserByName godoc
// @Summary Get a user by First Name and Last Name
// @Description Retrieve details of a specific user by their First Name and Last Name
// @Tags User
// @Param name path string true "User First Name and Last Name"
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {string} string "Invalid user name"
// @Failure 404 {string} string "User not found"
// @Router /users/name/{name} [get]
func GetUserByName(w http.ResponseWriter, r *http.Request) {
	// Retrieving the `name` parameter from the URL path using Gorilla Mux.
	params := mux.Vars(r)
	fullName := params["name"]
	// Splitting the full name into First Name and Last Name.
	names := strings.SplitN(fullName, " ", 2)

	// If the name doesn't split into two parts, return an error.
	if len(names) != 2 {
		http.Error(w, "Invalid user name", http.StatusBadRequest) // Return a 400 Bad Request if the name is not valid
		return
	}

	firstName := names[0]
	lastName := names[1]

	var user models.User // Declaring a variable `user` of type `models.User` to hold the retrieved user data

	// SQL statement to select a user by First Name and Last Name from the `users` table.
	sqlStatement := `SELECT * FROM users WHERE first_name=$1 AND last_name=$2`

	// Executing the SQL statement to retrieve the user.
	row := database.DB.QueryRow(sqlStatement, firstName, lastName)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CurrentAddress, &user.PermanentAddress, &user.Pincode, &user.City, &user.State, &user.ContactNumber, &user.Email, &user.Password, &user.EmergencyFirstName, &user.EmergencyLastName, &user.AadharCardNumber, &user.DrivingLicenseNumber, &user.PanCardNumber, &user.Branch, &user.UserRole, &user.CreatedAt)

	// Checking if there was an error while scanning the row.
	if err != nil {
		// If no rows were found, return a 404 Not Found error.
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			log.Fatal(err) // For any other error, log it and stop the program.
		}
		return
	}

	// Encode the retrieved `user` struct into JSON and write it to the response.
	json.NewEncoder(w).Encode(user)
}
