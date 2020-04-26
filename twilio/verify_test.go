package twilio

import (
	"errors"
	"testing"
)

func TestRecvVerificationSMS(t *testing.T) {

	type contact struct {
		phoneNumber string
		name        string
		recvSMS     func() error
		ok          bool
	}
	data := []contact{
		{
			phoneNumber: "+71234567890",
			name:        "RU",
			recvSMS: func() error {
				return errors.New("Fail to send validation code")
			},
			ok: false,
		},
		{
			phoneNumber: "+48564826882",
			name:        "PL",
			recvSMS: func() error {
				return nil
			},
			ok: true,
		},
	}

	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			if err := test.recvSMS(); err != nil && test.ok {
				t.Errorf("unable to send number verification code %v", err)
			}
		})
	}

}
