package main

import (
	"flag"
	"fmt"
	"os"
)

const Version = "0.1.0"

func main() {
	flag.Usage = func() {
		usage := `Hello prints hello world!

Usage: hello [opts] [name]

Args:

  name
    	Name to greet (default "world")

Options:

  -help
    	Print help menu`
		// other flags defined below
		fmt.Println(usage)
		flag.PrintDefaults()
	}
	version := flag.Bool("version", false, "Print version")
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()
	if *version {
		fmt.Printf("hello version %s\n", Version)
		os.Exit(0)
	}
	name := flag.Arg(0)
	if name == "" {
		name = "world"
	}
	fmt.Printf("hello %s!\n", name)
}
