package main

import (
	"flag"
	"log"

	"github.com/rubenvanstaden/crypto"
)

func main() {

	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()

    account := ""
    key := ""

	if len(args) > 0 {

		if args[0] == "new" && len(args) == 1 {
			account = args[0]
		}

		if args[0] == "decode" && len(args) == 2 {
			key = args[1]
		}
	}

	if key != "" {
        _, pubkey, err := crypto.DecodeBech32(key)
		if err != nil {
			log.Fatal("unable to generate public key")
		}
		log.Printf("%s", pubkey)
    }

	if account != "" {
		sk := crypto.GeneratePrivateKey()
		pk, err := crypto.GetPublicKey(sk)
		if err != nil {
			log.Fatal("unable to generate public key")
		}

        ns, err := crypto.EncodePrivateKey(sk)
		if err != nil {
			log.Fatal("unable to generate public key")
		}

        np, err := crypto.EncodePublicKey(pk)
		if err != nil {
			log.Fatal("unable to generate public key")
		}

		log.Printf("nsec: %s", ns)
		log.Printf("npub: %s", np)
	}
}
