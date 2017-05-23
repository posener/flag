package flag

type stringCompleter struct {
	val *string
	Completer
}

func newCustom(val string, completer Completer, ptr *string) *stringCompleter {
	p := &stringCompleter{
		val:       ptr,
		Completer: completer,
	}
	*(p.val) = val
	return p
}

func (s *stringCompleter) Set(val string) error {
	*(s.val) = val
	return nil
}

func (s *stringCompleter) Get() interface{} { return *(s.val) }

func (s *stringCompleter) String() string {
	if s.val == nil {
		return ""
	}
	return *(s.val)
}

// StringCompleter is a string flag with custom bash completion
func (f *FlagSet) StringCompleter(name string, completer Completer, value string, usage string) *string {
	p := new(string)
	f.StringCompleterVar(p, name, completer, value, usage)
	return p
}

// StringCompleterVar is a StringCompleter with a given pointer
func (f *FlagSet) StringCompleterVar(p *string, name string, completer Completer, value string, usage string) {
	f.Var(newCustom(value, completer, p), name, usage)
}

// StringCompleter is a string flag with custom bash completion
func StringCompleter(name string, completer Completer, value string, usage string) *string {
	p := new(string)
	CommandLine.StringCompleterVar(p, name, completer, value, usage)
	return p
}

// StringCompleterVar is a StringCompleter with a given pointer
func StringCompleterVar(p *string, name string, completer Completer, value string, usage string) {
	CommandLine.StringCompleterVar(p, name, completer, value, usage)
}
