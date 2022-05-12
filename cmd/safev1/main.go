package main

import (
	"clis-by-example/pkg/secret"
	"flag" // todo: replace flag with GNU compliant
	"log"
)

func main() {
	// todo: change to flagSet
	name := flag.String("n", "", "a name for our safe")
	flag.Parse()

	if name != nil && *name == "" {
		log.Fatalln("flag \"n\" should be set")
	}

	safe := secret.Secret{
		Name: *name,
	}

	log.Printf("the name of the safe is: \"%s\"", safe.GetName())
}
