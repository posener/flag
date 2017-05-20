package flag

import (
	gflag "flag"
	"os"

	"github.com/posener/complete/cmd"
)

var cli = cmd.CLI{}

// AddCompleteFlags adds to a flagset install and uninstall flags which
// invokes installation of bash completion in the user's home folder.
func AddCompleteFlags(f *gflag.FlagSet, installName, uninstallName string) {
	cli.Name = os.Args[0]
	cli.InstallName = installName
	cli.UninstallName = uninstallName
	cli.AddFlags(f)
}

// ParseInstallFlags parses and invokes the install or uninstall if those
// flags were given in the command line.
func ParseInstallFlags() bool {
	return cli.Run()
}
