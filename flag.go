package flag

import gflag "flag"

// FlagSet is a copy of the standard library FlagSet struct.
type FlagSet struct {
	*gflag.FlagSet
}

// NewFlagSet creates a new FlagSet
func NewFlagSet(name string, h gflag.ErrorHandling) *FlagSet {
	return &FlagSet{gflag.NewFlagSet(name, h)}
}

// CommandLine is a copy of CommandLine variable from flag standard library
var CommandLine = &FlagSet{gflag.CommandLine}

// Copy functions from standard library
var (
	Parse  = gflag.Parse
	Parsed = gflag.Parsed

	String      = gflag.String
	StringVar   = gflag.StringVar
	Int         = gflag.Int
	IntVar      = gflag.IntVar
	Int64       = gflag.Int64
	Int64Var    = gflag.Int64Var
	Float64     = gflag.Float64
	Float64Var  = gflag.Float64Var
	DurationVar = gflag.DurationVar
	UintVar     = gflag.UintVar
	NArg        = gflag.NArg
	Args        = gflag.Args
)
