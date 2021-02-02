package cli

import (
	"flag"
	"fmt"

	"github.com/nomad-software/crypt/output"
)

// Options contain the command line options passed to the program.
type Options struct {
	Encode bool
	Decode bool
	Help   bool
}

// ParseOptions parses the command line options.
func ParseOptions() *Options {
	var opt Options

	flag.BoolVar(&opt.Encode, "encode", false, "Encode input.")
	flag.BoolVar(&opt.Decode, "decode", false, "Decode input.")
	flag.BoolVar(&opt.Help, "help", false, "Show help.")
	flag.Parse()

	return &opt
}

// Valid checks command line options are valid.
func (opt *Options) Valid() bool {

	if opt.Encode && opt.Decode {
		output.Error("Encode or decode cannot be done at the same time.")
	}

	if !opt.Encode && !opt.Decode {
		output.Error("Either encode or decode must be specified.")
	}

	return true
}

// PrintUsage prints the usage of the program.
func (opt *Options) PrintUsage() {
	var banner = `     ____                  _
    / ___|_ __ _   _ _ __ | |_ 
   | |   | '__| | | | '_ \| __|
   | |___| |  | |_| | |_) | |_ 
    \____|_|   \__, | .__/ \__|
               |___/|_|
`
	fmt.Println(banner)
	flag.Usage()
}
