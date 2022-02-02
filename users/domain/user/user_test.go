package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	type test struct {
		inputID   string
		inputName string

		outputID   string
		outputName string

		fails bool
		err   error
	}

	tests := []test{
		{inputID: "id", inputName: "name", outputID: "id", outputName: "name", fails: false, err: nil},
		{inputID: "", inputName: "name", outputID: "", outputName: "name", fails: true, err: ErrInvalidUserID},
		{inputID: "id", inputName: "", outputID: "id", outputName: "", fails: true, err: ErrInvalidUserName},
	}

	for _, test := range tests {
		got, err := NewUser(test.inputID, test.inputName)

		if test.fails {
			assert.Equal(t, err, test.err)
			continue
		}

		assert.NoError(t, err)
		assert.Equal(t, got.ID, test.outputID)
		assert.Equal(t, got.Name, test.outputName)
	}
}
