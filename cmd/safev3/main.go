package main

import (
	"clis-by-example/pkg/secret"
	"fmt"
	"log"
	"os"
)

type SubCommand interface {
	Run() error
	Init([]string) error
}

func findCommand(args []string) (SubCommand, error) {
	subCommand := args[0]
	subCommands := map[string]SubCommand{
		"safe":      NewSafeCommand(secret.New("hello123")),
		"supersafe": NewSafeCommand(secret.New("password321")),
	}

	for cmdName, cmd := range subCommands {
		if cmdName == subCommand {
			err := cmd.Init(args[1:])
			if err != nil {
				return nil, err
			}

			return cmd, nil
		}
	}

	return nil, fmt.Errorf("sub commando does not exist: %s", subCommand)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("a subcommand is required!")
	}

	cmd, err := findCommand(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
