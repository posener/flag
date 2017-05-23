# flag

[![Build Status](https://travis-ci.org/posener/flag.svg?branch=master)](https://travis-ci.org/posener/flag)
[![codecov](https://codecov.io/gh/posener/flag/branch/master/graph/badge.svg)](https://codecov.io/gh/posener/flag)
[![GoDoc](https://godoc.org/github.com/posener/flag?status.svg)](http://godoc.org/github.com/posener/flag)
[![Go Report Card](https://goreportcard.com/badge/github.com/posener/flag)](https://goreportcard.com/report/github.com/posener/flag)

Like `flag`, but with bash completion support.

## Features

* Fully compatible with standard library `flag` package
* Bash completions for flag names and flag values
* Additional flag types provided:
  * [`File`/`Dir`](./flag_path.go): file completion flag
  * [`Bool`](./flag_bool.go): bool flag (that does not complete)
  * [`Choice`](./flag_choice.go): choices flag
  * [`StringCompleter`](./flag_completer.go): custom completions
  * Any other value that implements the [`Completer`](./complete.go) interface.

## Example

Here is an [example](./example/example.go)

## Usage

```diff
import (
-	"flag"
+	"github.com/posener/flag"
)

var (
-	file = flag.String("file", "", "file value")
+	file = flag.File("file", "*.md", "", "file value")
-	dir  = flag.String("dir", "", "dir value")
+	dir  = flag.Dir("dir", "*", "", "dir value")
	b    = flag.Bool("bool", false, "bool value")
	s    = flag.String("any", "", "string value")
-	opts = flag.String("choose", "", "some items to choose from")
+	opts = flag.Choice("choose", []string{"work", "dring}, "", "some items to choose from")
)

func main() {
+	flag.SetInstallFlags("complete", "uncomplete")
	flag.Parse()
+	if flag.Complete() {  // runs bash completion if necessary
+		return  // return from main without executing the rest of the command
+	}
    ...
}
```
