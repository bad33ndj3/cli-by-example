package main

import (
	"clis-by-example/pkg/secret"
	"fmt"
	"log"

	flag "github.com/spf13/pflag"
)

const safeCommand = "safe"

type SafeCommand struct {
	fs   *flag.FlagSet
	safe secret.Safe

	nameFlag string
}

func NewSafeCommand(safe secret.Safe) SubCommand {
	s := &SafeCommand{
		fs:   flag.NewFlagSet(safeCommand, flag.ExitOnError),
		safe: safe,
	}
	s.fs.StringVarP(&s.nameFlag, "name", "n", "", "a name for our safe")

	return s
}

func (s *SafeCommand) Init(args []string) error {
	err := s.fs.Parse(args)
	if err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	return nil
}

func (s *SafeCommand) Run() error {
	// our business logic
	s.safe.SetName(s.nameFlag)
	log.Println(fmt.Sprintf("the name of the safe is: \"%s\"", s.safe.GetName()))
	log.Println(fmt.Sprintf("the secret in the safe is: \"%s\"", s.safe.GetSecretName()))

	return nil
}
