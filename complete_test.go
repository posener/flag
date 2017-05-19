package flag_test

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/posener/flag"
)

var once = sync.Once{}

func chdir() {
	once.Do(func() {
		os.Chdir("tests")
	})
}

func TestComplete(t *testing.T) {
	// NoParallel: this test modifies environment variables and stdout
	chdir()

	tests := []struct {
		line       string
		want       []string
		file       string
		dir        string
		bool       bool
		string     string
		parseError bool
	}{
		{
			line: "command ",
			want: []string{"-file", "-dir", "-bool", "-any"},
		},
		{
			line:       "command -file",
			want:       []string{"-file"},
			parseError: true,
		},
		{
			line: "command -file ",
			want: []string{"./", "sub/", "readme.md"},
		},
		{
			line: "command -file .",
			want: []string{"./", "./sub/", "./readme.md"},
			file: ".",
		},
		{
			line: "command -file r",
			want: []string{"readme.md"},
			file: "r",
		},
		{
			line: "command -file readme.md",
			want: []string{"readme.md"},
			file: "readme.md",
		},
		{
			line: "command -file s",
			want: []string{"sub/", "sub/readme.md"},
			file: "s",
		},
		{
			line: "command -file x",
			want: []string{},
			file: "x",
		},
		{
			line: "command -dir ",
			want: []string{"./", "sub/"},
		},
		{
			line: "command -dir s",
			want: []string{"sub/"},
			dir:  "s",
		},
		{
			line: "command -dir x",
			want: []string{},
			dir:  "x",
		},
		{
			line: "command -dir .",
			want: []string{"./", "./sub/"},
			dir:  ".",
		},
		{
			line: "command -bool ",
			want: []string{"-file", "-dir", "-bool", "-any"},
			bool: true,
		},
		{
			line: "command -any ",
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {

			fs := flag.NewFlagSet("command", 0)

			file := fs.File("file", "*.md", "", "file value")
			dir := fs.Dir("dir", "*", "", "dir value")
			b := fs.Bool("bool", false, "bool value")
			s := fs.String("any", "", "string value")

			os.Setenv("COMP_LINE", tt.line)
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			defer r.Close()
			os.Stdout = w
			if !fs.Complete() {
				t.Error("failed to complete")
			}
			w.Close()
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				t.Fatal(err)
			}
			got := []string{}
			for _, word := range strings.Split(string(buf), "\n") {
				if word != "" && notTestFlag(word) {
					got = append(got, word)
				}
			}

			sort.Strings(got)
			sort.Strings(tt.want)
			assert.Equal(t, tt.want, got, t.Name())

			err = fs.Parse(strings.Split(tt.line, " ")[1:])
			if err == nil && tt.parseError {
				t.Error("parse did not parseError when expected", t.Name())
			} else if err != nil && !tt.parseError {
				t.Error("parse parseError when not expected", r, t.Name())
			}
			assert.Equal(t, tt.file, *file, t.Name())
			assert.Equal(t, tt.dir, *dir, t.Name())
			assert.Equal(t, tt.bool, *b, t.Name())
			assert.Equal(t, tt.string, *s, t.Name())
		})
	}
}

// ExampleComplete shows how bash completion works
func ExampleComplete() {
	chdir()

	// define a flag
	flag.File("file", "*.md", "", "file value")

	// set the environment variable used for setting the completion
	// look that there is a space after the '-file' flag, indicating that the user
	// want to complete by that flag
	os.Setenv("COMP_LINE", "demo -file ")

	// complete
	flag.Complete()

	// Unordered output: readme.md
	// sub/
	// ./
}

func notTestFlag(f string) bool {
	return !strings.HasPrefix(f, "-") || !strings.Contains(f, "test.")
}
