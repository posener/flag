# flag

## Sub Packages

* [example](./example)

#### Examples

##### Complete

ExampleComplete shows how bash completion works

```golang
package main

import (
	"github.com/posener/flag"
	"os"
	"sync"
)

var once = sync.Once{}

func chdir() {
	once.Do(func() {
		os.Chdir("tests")
	})
}

func main() {
	chdir()

	// define a flag
	flag.File("file", "*.md", "", "file value")

	// set the environment variable used for setting the completion
	// look that there is a space after the '-file' flag, indicating that the user
	// want to complete by that flag
	os.Setenv("COMP_LINE", "demo -file ")

	// complete
	flag.Complete()

}

```

 Output:

```
readme.md
sub/
./

```


---

Created by [goreadme](https://github.com/apps/goreadme)
