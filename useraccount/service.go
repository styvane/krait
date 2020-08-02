package useraccount

// Userservice defines the interface type for the user
type UserService interface {
	Save() error
	Update() error
	Delete() error
}
