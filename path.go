package flag

import (
	"os"
	"path/filepath"
	"strings"
)

type path struct {
	path       *string
	pattern    string
	allowFiles bool
}

func newPathValue(val string, pattern string, allowFiles bool, ptr *string) *path {
	p := &path{
		path:       ptr,
		pattern:    pattern,
		allowFiles: allowFiles,
	}
	*(p.path) = val
	return p
}

func (p *path) Set(s string) error {
	*(p.path) = s
	return nil
}

func (p *path) Get() interface{} { return *(p.path) }

func (p *path) String() string {
	if p.path == nil {
		return ""
	}
	return *(p.path)
}

func (p *path) Complete(s string) ([]string, bool) {
	return fileCompleter(p.pattern, p.allowFiles)(s)
}

// File is a flag that is related to a file path on the filesystem
// according to a given pattern. for any file use "*" as the pattern.
func (f *FlagSet) File(name string, pattern string, value string, usage string) *string {
	p := new(string)
	f.FileVar(p, name, pattern, value, usage)
	return p
}

// FileVar is a File with a given pointer
func (f *FlagSet) FileVar(p *string, name string, pattern string, value string, usage string) {
	f.Var(newPathValue(value, pattern, true, p), name, usage)
}

// Dir is a flag the relates to a directory on the filesystem according to
// a given pattern. For any directory name use "*" as the pattern.
func (f *FlagSet) Dir(name string, pattern string, value string, usage string) *string {
	p := new(string)
	f.DirVar(p, name, pattern, value, usage)
	return p
}

// DirVar is a Dir with a given pointer
func (f *FlagSet) DirVar(p *string, name string, pattern string, value string, usage string) {
	f.Var(newPathValue(value, pattern, false, p), name, usage)
}

// File is a flag that is related to a file path on the filesystem
// according to a given pattern. for any file use "*" as the pattern.
func File(name string, pattern string, value string, usage string) *string {
	p := new(string)
	CommandLine.FileVar(p, name, pattern, value, usage)
	return p
}

// FileVar is a File with a given pointer
func FileVar(p *string, name string, pattern string, value string, usage string) {
	CommandLine.FileVar(p, name, pattern, value, usage)
}

// DirVar is a Dir with a given pointer
func DirVar(p *string, name string, pattern string, value string, usage string) {
	CommandLine.DirVar(p, name, pattern, value, usage)
}

// Dir is a flag the relates to a directory on the filesystem according to
// a given pattern. For any directory name use "*" as the pattern.
func Dir(name string, pattern string, value string, usage string) *string {
	p := new(string)
	CommandLine.DirVar(p, name, pattern, value, usage)
	return p
}

func fileCompleter(pattern string, allowFiles bool) CompleteFn {

	// search for files according to arguments,
	// if only one directory has matched the result, search recursively into
	// this directory to give more results.
	return func(last string) (options []string, only bool) {
		only = true
		options = predictFiles(last, pattern, allowFiles)

		// if the number of options is not 1, we either have many results or
		// have no results, so we return it.
		if len(options) != 1 {
			return
		}

		// only try deeper, if the one item is a directory
		if stat, err := os.Stat(options[0]); err != nil || !stat.IsDir() {
			return
		}

		return predictFiles(options[0], pattern, allowFiles), only
	}
}

func predictFiles(last string, pattern string, allowFiles bool) []string {
	if strings.HasSuffix(last, "/..") {
		return nil
	}

	dir := directory(last)
	files := listFiles(dir, pattern, allowFiles)

	// add dir
	files = append(files, dir)

	return files.complete(last)
}

type files []string

func (f files) complete(last string) []string {
	var options []string
	for _, fi := range f {
		fi = fixPathForm(last, fi)

		// test matching of file to the argument
		if matchFile(fi, last) {
			options = append(options, fi)
		}
	}
	return options
}

func matchFile(file, prefix string) bool {
	// special case for current directory completion
	if file == "./" && (prefix == "." || prefix == "") {
		return true
	}
	if prefix == "." && strings.HasPrefix(file, ".") {
		return true
	}

	file = strings.TrimPrefix(file, "./")
	prefix = strings.TrimPrefix(prefix, "./")

	return strings.HasPrefix(file, prefix)
}

func listFiles(dir, pattern string, allowFiles bool) files {
	// stringChoice of all file names
	list := files{}

	dirPattern := pattern

	// list files
	if allowFiles {
		if files, err := filepath.Glob(filepath.Join(dir, pattern)); err == nil {
			for _, f := range files {
				if stat, err := os.Stat(f); err != nil || !stat.IsDir() {
					list = append(list, f)
				}
			}
		}
		dirPattern = "*"
	}

	// list directories
	if files, err := filepath.Glob(filepath.Join(dir, dirPattern)); err == nil {
		for _, f := range files {
			if stat, err := os.Stat(f); err != nil || stat.IsDir() {
				list = append(list, f)
			}
		}
	}
	return list
}

// directory gives the directory of the current written
// last argument if it represents a file name being written.
// in case that it is not, we fall back to the current directory.
func directory(last string) string {
	if info, err := os.Stat(last); err == nil && info.IsDir() {
		return fixPathForm(last, last)
	}
	dir := filepath.Dir(last)
	if info, err := os.Stat(dir); err != nil || !info.IsDir() {
		return "./"
	}
	return fixPathForm(last, dir)
}

// fixPathForm changes a file name to a relative name
func fixPathForm(last string, file string) string {
	// get wording directory for relative name
	workDir, err := os.Getwd()
	if err != nil {
		return file
	}

	abs, err := filepath.Abs(file)
	if err != nil {
		return file
	}

	// if last is absolute, return path as absolute
	if filepath.IsAbs(last) {
		return fixDirPath(abs)
	}

	rel, err := filepath.Rel(workDir, abs)
	if err != nil {
		return file
	}

	// fix ./ prefix of path
	if rel != "." && strings.HasPrefix(last, ".") {
		rel = "./" + rel
	}

	return fixDirPath(rel)
}

func fixDirPath(path string) string {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() && !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return path
}
