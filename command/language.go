package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/nlp"
)

var languagecmd = &cobra.Command{
	Use:   "languages",
	Short: "Show supported languages",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		fmt.Println("Supported Languages:")
		for _, lang := range nlp.SupportedLanguages {
			fmt.Printf("  %s (%s)\n", lang.Code, lang.Name)
		}

		return
	},
}
