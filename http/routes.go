package http

// routes register HTTP handlers
func (s *server) routes() {
	s.router.HandleFunc("/signup", s.InitiateRegistrationHandle())
}
