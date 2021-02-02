package main

import (
	"io"
	"os"

	"github.com/nomad-software/crypt/aes"
	"github.com/nomad-software/crypt/cli"
	"github.com/nomad-software/crypt/output"
)

func main() {
	options := cli.ParseOptions()

	if options.Help {
		options.PrintUsage()

	} else if options.Valid() {
		data, err := io.ReadAll(os.Stdin)
		output.OnError(err, "Could not read stdin.")

		password, err := cli.ReadPassword("enter password: ")
		output.OnError(err, "Could not read password from stdin.")

		if options.Encode {
			data, err = aes.Encrypt(password, data)
			output.OnError(err, "Could not encode input.")
		} else if options.Decode {
			data, err = aes.Decrypt(password, data)
			output.OnError(err, "Could not decode input.")
		}

		_, err = os.Stdout.Write(data)
		output.OnError(err, "Could not write output.")
	}
}
