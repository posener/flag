package flag

import (
	"fmt"
	"strings"
)

type choice struct {
	chosen  *string
	choices []string
}

func newChoice(val string, choices []string, ptr *string) *choice {
	p := &choice{
		chosen:  ptr,
		choices: choices,
	}
	*(p.chosen) = val
	return p
}

func (s *choice) Set(v string) error {
	for _, c := range s.choices {
		if c == v {
			*(s.chosen) = v
			return nil
		}
	}
	return fmt.Errorf("invalid choice %s, allowed choice: %s", s, s.choices)
}

func (s *choice) Get() interface{} { return *(s.chosen) }

func (s *choice) String() string {
	if s.chosen == nil {
		return ""
	}
	return *(s.chosen)
}

func (s *choice) Complete(last string) ([]string, bool) {
	options := make([]string, 0, len(s.choices))
	for _, c := range s.choices {
		if strings.HasPrefix(c, last) {
			options = append(options, c)
		}
	}
	return options, true
}

// Choice is a flag that has a choice of possible string values
func (f *FlagSet) Choice(name string, choices []string, value string, usage string) *string {
	p := new(string)
	f.ChoiceVar(p, name, choices, value, usage)
	return p
}

// ChoiceVar is a Choice with a given pointer
func (f *FlagSet) ChoiceVar(p *string, name string, choices []string, value string, usage string) {
	f.Var(newChoice(value, choices, p), name, usage)
}

// Choice is a flag that has a choice of possible string values
func Choice(name string, choices []string, value string, usage string) *string {
	p := new(string)
	CommandLine.ChoiceVar(p, name, choices, value, usage)
	return p
}

// ChoiceVar is a Choice with a given pointer
func ChoiceVar(p *string, name string, choices []string, value string, usage string) {
	CommandLine.ChoiceVar(p, name, choices, value, usage)
}
