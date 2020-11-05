package main

import (
	"flag"
	"fmt"
	"os"
)

const Version = "0.1.2"

func main() {
	flag.Usage = func() {
		usage := `Hello prints hello world!

Usage: hello [opts] [name]

Args:

  name
    	Name to greet (default "world")

Options:
`
		// other flags defined below
		fmt.Println(usage)
		flag.PrintDefaults()
	}
	help := flag.Bool("help", false, "Print help menu")
	version := flag.Bool("version", false, "Print version")
	feature1 := flag.Bool("feature1", false, "Print feature1")
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if *version {
		fmt.Printf("hello version %s\n", Version)
		os.Exit(0)
	}
	if *feature1 {
		fmt.Println("hello feature1")
		os.Exit(0)
	}
	name := flag.Arg(0)
	if name == "" {
		name = "world"
	}
	fmt.Printf("hello %s!\n", name)
}
