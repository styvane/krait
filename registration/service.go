package registration

// Userservice defines the interface type for the user
type UserService interface {
	Update(p *Profile) error
	Delete(p *Profile) error
}
