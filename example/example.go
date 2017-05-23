package main

import (
	"fmt"
	"strings"

	"github.com/posener/flag"
)

func complete(last string) ([]string, bool) {
	// here you can do whatever you want, http requests, read files, whatever!
	if strings.HasPrefix("custom-string", last) {
		return []string{"custom-string"}, true
	}
	return []string{}, true
}

var (
	file   = flag.File("file", "*.md", "", "file value")
	dir    = flag.Dir("dir", "*", "", "dir value")
	choice = flag.Choice("choose", []string{"one", "two", "three"}, "", "choose between a set of values")
	comp   = flag.StringCompleter("comp", flag.CompleteFn(complete), "", "custom complete function")
	b      = flag.Bool("bool", false, "bool value")
	s      = flag.String("any", "", "string value")
)

func main() {
	// add installation flags.
	// running 'example -complete' will install the completion script in the user's
	// home directory. running 'example -uncomplete' will uninstall it.
	flag.SetInstallFlags("complete", "uncomplete")
	flag.Parse()

	// runs bash completion if necessary
	if flag.Complete() {
		// return from main without executing the rest of the command
		return
	}

	fmt.Println("file:", *file)
	fmt.Println("dir:", *dir)
	fmt.Println("choice:", *choice)
	fmt.Println("string:", *comp)
	fmt.Println("bool:", *b)
	fmt.Println("string:", *s)
}
