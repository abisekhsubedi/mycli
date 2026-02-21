package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	blue   = color.New(color.FgBlue).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
)

func Info(format string, a ...any) {
	fmt.Fprintf(os.Stdout, "%s %s\n", blue("info"), fmt.Sprintf(format, a...))
}

func Success(format string, a ...any) {
	fmt.Fprintf(os.Stdout, "%s %s\n", green("success"), fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "%s %s\n", yellow("warn"), fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "%s %s\n", red("error"), fmt.Sprintf(format, a...))
}
