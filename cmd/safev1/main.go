package main

import (
	"clis-by-example/pkg/secret"
	"log"
	"os"

	flag "github.com/spf13/pflag" // todo: replace flag with GNU compliant
)

func main() {
	// todo: change to flagSet
	fs := flag.CommandLine
	name := fs.StringP("name", "n", "", "a name for our safe")
	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	if name != nil && *name == "" {
		log.Fatalln("flag \"n\" should be set")
	}

	safe := secret.Secret{
		Name: *name,
	}

	log.Printf("the name of the safe is: \"%s\"", safe.GetName())
}
