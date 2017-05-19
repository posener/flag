package flag

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		line          string
		completed     string
		last          string
		lastCompleted string
	}{
		{
			line:          "a b c",
			completed:     "b",
			last:          "c",
			lastCompleted: "b",
		},
		{
			line:          "a b ",
			completed:     "b",
			last:          "",
			lastCompleted: "b",
		},
		{
			line:          "",
			completed:     "",
			last:          "",
			lastCompleted: "",
		},
		{
			line:          "a",
			completed:     "",
			last:          "a",
			lastCompleted: "",
		},
		{
			line:          "a ",
			completed:     "",
			last:          "",
			lastCompleted: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {

			os.Setenv(envLine, tt.line)
			a := readLine()
			if a == nil {
				t.Fatal("a is nil", t.Name())
			}

			if got, want := strings.Join(a.completed, " "), tt.completed; got != want {
				t.Errorf("%s failed: Completed = %q, want %q", t.Name(), got, want)
			}
			if got, want := a.last, tt.last; got != want {
				t.Errorf("Last = %q, want %q", got, want)
			}
			if got, want := a.lastCompleted, tt.lastCompleted; got != want {
				t.Errorf("%s failed: LastCompleted = %q, want %q", t.Name(), got, want)
			}
		})
	}
}

func TestArgs_From(t *testing.T) {
	t.Parallel()
	tests := []struct {
		line         string
		from         int
		newLine      string
		newCompleted string
	}{
		{
			line:         "a b c",
			from:         0,
			newLine:      "b c",
			newCompleted: "b",
		},
		{
			line:         "a b c",
			from:         1,
			newLine:      "c",
			newCompleted: "",
		},
		{
			line:         "a b c",
			from:         2,
			newLine:      "",
			newCompleted: "",
		},
		{
			line:         "a b c",
			from:         3,
			newLine:      "",
			newCompleted: "",
		},
		{
			line:         "a b c ",
			from:         0,
			newLine:      "b c ",
			newCompleted: "b c",
		},
		{
			line:         "a b c ",
			from:         1,
			newLine:      "c ",
			newCompleted: "c",
		},
		{
			line:         "a b c ",
			from:         2,
			newLine:      "",
			newCompleted: "",
		},
		{
			line:         "",
			from:         0,
			newLine:      "",
			newCompleted: "",
		},
		{
			line:         "",
			from:         1,
			newLine:      "",
			newCompleted: "",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s/%d", tt.line, tt.from), func(t *testing.T) {

			os.Setenv(envLine, tt.line)
			a := readLine()
			n := a.from(tt.from)

			if got, want := strings.Join(n.all, " "), tt.newLine; got != want {
				t.Errorf("%s failed: all = %q, want %q", t.Name(), got, want)
			}
			if got, want := strings.Join(n.completed, " "), tt.newCompleted; got != want {
				t.Errorf("%s failed: completed = %q, want %q", t.Name(), got, want)
			}
		})
	}
}
