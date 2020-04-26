package twilio

// SMSVerifier is an interface that wrap the RecvSMS method
type SMSVerifier interface {
	RecvSMS() error
}
