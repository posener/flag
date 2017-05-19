package flag

import (
	"os"
	"strings"
)

const envLine = "COMP_LINE"

// line describes command line arguments
type line struct {
	// all lists of all arguments in command line (not including the command itself)
	all []string
	// completed lists of all completed arguments in command line,
	// If the last one is still being typed - no space after it,
	// it won't appear in this list of arguments.
	completed []string
	// last argument in command line, the one being typed, if the last
	// character in the command line is a space, this argument will be empty,
	// otherwise this would be the last word.
	last string
	// lastCompleted is the last argument that was fully typed.
	// If the last character in the command line is space, this would be the
	// last word, otherwise, it would be the word before that.
	lastCompleted string
}

func readLine() *line {
	l, ok := os.LookupEnv(envLine)
	if !ok {
		return nil
	}
	words := strings.Split(l, " ")
	completed := removeLast(words[1:])
	return &line{
		all:           words[1:],
		completed:     completed,
		last:          last(words),
		lastCompleted: last(completed),
	}
}

func (a line) from(i int) line {
	if i > len(a.all) {
		i = len(a.all)
	}
	a.all = a.all[i:]

	if i > len(a.completed) {
		i = len(a.completed)
	}
	a.completed = a.completed[i:]
	return a
}

func removeLast(a []string) []string {
	if len(a) > 0 {
		return a[:len(a)-1]
	}
	return a
}

func last(args []string) (last string) {
	if len(args) > 0 {
		last = args[len(args)-1]
	}
	return
}
