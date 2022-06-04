package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	type test struct {
		InputID   string
		InputName string

		OutputID   string
		OutputName string

		Error error
	}

	tests := map[string]test{
		"success":         {InputID: "id", InputName: "name", OutputID: "id", OutputName: "name", Error: nil},
		"empty user id":   {InputID: "", InputName: "name", OutputID: "", OutputName: "name", Error: ErrEmptyUserID},
		"empty user name": {InputID: "id", InputName: "", OutputID: "id", OutputName: "", Error: ErrEmptyUserName},
	}

	for name, test := range tests {
		got, err := NewUser(test.InputID, test.InputName)
		if err != nil {
			assert.Equalf(t, err, test.Error, name)
			continue
		}
		assert.Equalf(t, got.ID, test.OutputID, name)
		assert.Equalf(t, got.Name, test.OutputName, name)
	}
}
