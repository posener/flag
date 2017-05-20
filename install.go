package flag

import (
	"os"

	"github.com/posener/complete/cmd"
)

var cli = cmd.CLI{}

// SetInstallFlags adds to a flagset install and uninstall flags which
// invokes installation of bash completion in the user's home folder.
func (f *FlagSet) SetInstallFlags(installName, uninstallName string) {
	cli.Name = os.Args[0]
	cli.InstallName = installName
	cli.UninstallName = uninstallName
	cli.AddFlags(f.FlagSet)
}

// SetInstallFlags adds to a flagset install and uninstall flags which
// invokes installation of bash completion in the user's home folder.
func SetInstallFlags(installName, uninstallName string) {
	CommandLine.SetInstallFlags(installName, uninstallName)
}
