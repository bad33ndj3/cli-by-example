package main

import (
	"clis-by-example/pkg/secret/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestCommand_Run(t *testing.T) {
	safe := &mocks.Safe{}
	safe.On("SetName", mock.Anything)
	safe.On("GetSecretName").Return("secretName")
	safe.On("GetName").Return("name")

	cmd := NewCommand(safe)
	err := cmd.Run([]string{"a"})
	if err != nil {
		t.Errorf("err should be nil but is, %s", err.Error())
	}
}
