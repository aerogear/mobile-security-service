package models

// User is the model struct for users
// swagger:model User
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
