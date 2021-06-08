package command

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "texcheck",
	Short:   "A command-line tool for error correction of LaTeX files",
}


func init() {

	rootCmd.AddCommand(serverCmd())
	rootCmd.AddCommand(newFileCmd())
	rootCmd.AddCommand(languagecmd)
}

// Execute is a executor of command-line tool
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
