package utils

import (
	"fmt"
)

const (
	colorReset = "\033[0m"
	colorRed = "\033[31m"
	colorGreen = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan = "\033[36m"
	colorWhite = "\033[37m"
)

func Prompt(s string) string {
	return fmt.Sprintf("%s%s%s: ", colorBlue, s, colorReset)
}

func Error(s string) string {
	return fmt.Sprintf("%s%s%s", colorRed, s, colorReset)
}

func Suggestion(s string) string {
	return fmt.Sprintf("%s%s%s", colorCyan, s, colorReset)
}

func Index(idx int) string {
	return fmt.Sprintf("%s    %d) %s", colorYellow, idx, colorReset)
}