package flag

import (
	gflag "flag"
	"fmt"
	"strings"
)

// Complete tries to complete the command line that was already
// entered. It returns true on success, on which the main program
// should exit
func Complete() bool {
	return CommandLine.Complete()
}

// Complete ctries to complete the command line according to a FlagSet.
// It returns true on success, on which the main program should exit.
func (f *FlagSet) Complete() bool {
	l := readLine()
	if l == nil {
		return cli.Run()
	}

	options := f.complete(l)

	for _, option := range options {
		fmt.Println(option)
	}
	return true
}

func (f *FlagSet) complete(line *line) []string {
	var options []string

	// flag completion according to last completed argument in command line
	if flg := f.lastFlag(line.lastCompleted); flg != nil {
		if cmp, ok := flg.Value.(Completer); ok {
			// last flag implements the Completer interface, complete according
			// to its decision
			var only bool
			options, only = cmp.Complete(line.last)
			if only {
				return options
			}
		} else {
			// standard library flag, we want to return an empty
			// stringSet of completion, since anything can come after this
			// kind of flag, and we should not complete any other
			// flags after it.
			return []string{}
		}
	}

	// add all flag names to the complete options
	f.VisitAll(func(flg *gflag.Flag) {
		name := "-" + flg.Name
		if strings.HasPrefix(name, line.last) {
			options = append(options, name)
		}
	})
	return options
}

func (f *FlagSet) lastFlag(lastCompleted string) *gflag.Flag {
	if !strings.HasPrefix(lastCompleted, "-") {
		return nil
	}
	return f.Lookup(lastCompleted[1:])
}

// Completer interface is for the Complete function
type Completer interface {
	// Complete according to the given last string in the command line.
	// returns list of options to complete the last word in the command
	// line, and a boolean indicating if those options should be the
	// only shown options.
	Complete(last string) (options []string, only bool)
}

// CompleteFn is function implementing the Completer interface
type CompleteFn func(string) ([]string, bool)

// Complete implements the Complete interface
func (c CompleteFn) Complete(last string) ([]string, bool) {
	if c == nil {
		return nil, false
	}
	return c(last)
}
