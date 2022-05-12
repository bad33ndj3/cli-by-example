package main

import (
	"clis-by-example/pkg/secret"
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	safe := secret.New("generics")
	cmd := NewCommand(safe)

	err := cmd.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}

// Command is the module around our command
type Command struct {
	fs   *flag.FlagSet
	safe secret.Safe

	nameFlag string
}

// NewCommand inits the command on flag.CommandLine
func NewCommand(safe secret.Safe) *Command {
	s := &Command{
		fs:   flag.CommandLine,
		safe: safe,
	}

	s.fs.StringVarP(&s.nameFlag, "name", "n", "", "a name for our safe")

	return s
}

// Run executes our command and prints our safe's name and secretName
func (s *Command) Run(args []string) error {
	err := s.fs.Parse(args)
	if err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	// our business logic
	s.safe.SetName(s.nameFlag)
	log.Println(fmt.Sprintf("the name of the safe is: \"%s\"", s.safe.GetName()))
	log.Println(fmt.Sprintf("the secret in the safe is: \"%s\"", s.safe.GetSecretName()))

	return nil
}
