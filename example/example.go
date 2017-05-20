package main

import (
	"fmt"

	"github.com/posener/flag"
)

var (
	file   = flag.File("file", "*.md", "", "file value")
	dir    = flag.Dir("dir", "*", "", "dir value")
	choice = flag.StringSet("choose", []string{"one", "two", "three"}, "", "choose between a set of values")
	b      = flag.Bool("bool", false, "bool value")
	s      = flag.String("any", "", "string value")
)

func main() {
	// runs bash completion if necessary
	if flag.Complete() {
		// return from main without executing the rest of the command
		return
	}
	// add flags for (un)installing bash completion
	flag.AddCompleteFlags(nil, "complete", "uncomplete")

	flag.Parse()

	// parse bash completion installation flags
	if flag.ParseInstallFlags() {
		// return from main if the bash completion was installed
		return
	}

	fmt.Println("file:", *file)
	fmt.Println("dir:", *dir)
	fmt.Println("choice:", *choice)
	fmt.Println("bool:", *b)
	fmt.Println("string:", *s)
}
