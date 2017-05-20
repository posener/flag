package flag

import (
	"fmt"
	"strings"
)

type stringSet struct {
	chosen  *string
	choices []string
}

func newSet(val string, choices []string, ptr *string) *stringSet {
	p := &stringSet{
		chosen:  ptr,
		choices: choices,
	}
	*(p.chosen) = val
	return p
}

func (s *stringSet) Set(v string) error {
	for _, c := range s.choices {
		if c == v {
			*(s.chosen) = v
			return nil
		}
	}
	return fmt.Errorf("invalid choice %s, allowed stringSet: %s", s, s.choices)
}

func (s *stringSet) Get() interface{} { return *(s.chosen) }

func (s *stringSet) String() string {
	if s.chosen == nil {
		return ""
	}
	return *(s.chosen)
}

func (s *stringSet) Complete(last string) ([]string, bool) {
	options := make([]string, 0, len(s.choices))
	for _, c := range s.choices {
		if strings.HasPrefix(c, last) {
			options = append(options, c)
		}
	}
	return options, true
}

// StringSet is a flag that has a stringSet of possible string values
func (f *FlagSet) StringSet(name string, choices []string, value string, usage string) *string {
	p := new(string)
	f.StringSetVar(p, name, choices, value, usage)
	return p
}

// StringSetVar is a StringSet with a given pointer
func (f *FlagSet) StringSetVar(p *string, name string, choices []string, value string, usage string) {
	f.Var(newSet(value, choices, p), name, usage)
}

// StringSet is a flag that has a stringSet of possible string values
func StringSet(name string, choices []string, value string, usage string) *string {
	p := new(string)
	CommandLine.StringSetVar(p, name, choices, value, usage)
	return p
}

// StringSetVar is a StringSet with a given pointer
func StringSetVar(p *string, name string, choices []string, value string, usage string) {
	CommandLine.StringSetVar(p, name, choices, value, usage)
}
