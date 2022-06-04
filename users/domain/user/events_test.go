package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserRegisteredEvent(t *testing.T) {
	tests := map[string]struct {
		ID            string
		Name          string
		ExpectedError error
	}{
		"success": {
			ID:            "00000000-0000-0000-0000-000000000000",
			Name:          "dummy",
			ExpectedError: nil,
		},
		"empty id": {
			ID:            "",
			Name:          "dummy",
			ExpectedError: errors.New("invalid aggregate id"),
		},
		"empty name": {
			ID:            "00000000-0000-0000-0000-000000000000",
			Name:          "",
			ExpectedError: errors.New("invalid user name"),
		},
	}

	for name, test := range tests {
		_, err := UserEventsFactory{}.NewUserRegisteredEvent(test.ID, test.Name)
		assert.Equalf(t, test.ExpectedError, err, name)
	}
}
