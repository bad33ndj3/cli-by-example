package main

import (
	secret "clis-by-example/pkg/secret/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSafeCommand_Run(t *testing.T) {
	safe := &secret.Safe{}
	safe.On("GetSecretName").Return("secretName")
	safe.On("GetName").Return("getName")
	safe.On("SetName", mock.Anything)

	cmd := NewSafeCommand(safe)

	err := cmd.Run([]string{"getName"})
	if err != nil {
		t.Errorf("err should be nil, but is: %s", err.Error())
	}
}
