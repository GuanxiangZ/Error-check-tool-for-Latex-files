package command

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/config"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/parser"
	"os"
)





func newFileCmd() *cobra.Command {
	cfg := config.GetFileCmdConfig()

	cmd := &cobra.Command{
		Use:     "file",
		Short:   "Do grammatical errors check for LaTeX file",
		Example: "texcheck file --input demo.tex --output pdf",
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			var (
				file    *os.File
				//content string
			)

			file, err = os.Open(cfg.Input)
			if err != nil {
				logrus.Error(err)
				return
			}

			defer file.Close()

			_, err = parser.ProcessTeX(file, cfg.Input, cfg.Language, true)
			if err != nil {
				logrus.Error(err)
				return
			}

			//parser.Report(cfg.Input)

			return err
		},
	}

	cmd.Flags().StringVarP(&cfg.Input, "input", "i", "", "The path of the LaTeX file")
	cmd.Flags().StringVarP(&cfg.Output, "output", "o", "stdout", "Specify the report file type (json | pdf | stdout)")
	cmd.Flags().StringVar(&cfg.Outdir, "outdir", "", "The path to store the report file")
	cmd.Flags().StringVarP(&cfg.Language, "language", "l", "en-US", "Specify the language")
	cmd.Flags().StringVarP(&cfg.Addr, "addr", "a", "https://api.languagetool.org", "Specify the LanguageTool server url address")

	return cmd
}
