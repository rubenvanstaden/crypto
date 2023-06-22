package main

import (
	"flag"
	"log"
)

func main() {

	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()

	nsec := ""

	if len(args) > 0 {
		// Generate a new private, public key pair.
		if args[0] == "gen" && len(args) == 1 {
			nsec = args[0]
		}
	}

	if nsec != "" {
		sk := GeneratePrivateKey()

		pk, err := GetPublicKey(sk)
		if err != nil {
			log.Fatal("unable to generate public key")
		}

		log.Printf("nsec: %s", sk)
		log.Printf("npub: %s", pk)
	}
}
