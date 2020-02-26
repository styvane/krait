package registration

import (
	"testing"
)

func TestValidRequest(t *testing.T) {
	r := Request{"+71234567890", "Ru"}
	err := r.Validate()
	if err != nil {
		t.Error("Invalid request")
	}
}
