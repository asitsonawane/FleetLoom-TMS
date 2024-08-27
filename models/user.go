package models

import "time" // Importing the time package to handle time-related functions

// User struct defines the structure of the user entity that will be stored in the database.
// Each field in the struct corresponds to a column in the `users` table.
type User struct {
	ID                   int       `json:"id"`                     // Unique identifier for the user (Primary Key)
	FirstName            string    `json:"first_name"`             // User's first name
	LastName             string    `json:"last_name"`              // User's last name
	CurrentAddress       string    `json:"current_address"`        // User's current address
	PermanentAddress     string    `json:"permanent_address"`      // User's permanent address
	Pincode              string    `json:"pincode"`                // Pincode for the user's address
	City                 string    `json:"city"`                   // City where the user resides
	State                string    `json:"state"`                  // State where the user resides
	ContactNumber        string    `json:"contact_number"`         // User's contact number
	Email                string    `json:"email"`                  // User's email address
	Password             string    `json:"password"`               // User's password (should be hashed in a real application)
	EmergencyFirstName   string    `json:"emergency_first_name"`   // First name of the user's emergency contact
	EmergencyLastName    string    `json:"emergency_last_name"`    // Last name of the user's emergency contact
	AadharCardNumber     string    `json:"aadhar_card_number"`     // User's Aadhar card number (Indian ID)
	DrivingLicenseNumber string    `json:"driving_license_number"` // User's driving license number
	PanCardNumber        string    `json:"pan_card_number"`        // User's PAN card number (Indian tax ID)
	Branch               string    `json:"branch"`                 // The branch where the user is associated
	CreatedAt            time.Time `json:"created_at"`             // Timestamp when the user was created
}
