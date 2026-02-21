package commands

import (
	"strings"

	"github.com/abisekhsubedi/mycli/internal/logger"
	"github.com/spf13/cobra"
)

func NewGreetCmd() *cobra.Command {
	var uppercase bool

	cmd := &cobra.Command{
		Use:   "greet <name>",
		Short: "Greet a user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			if uppercase {
				name = strings.ToUpper(name)
			}
			logger.Success("Hello, %s!", name)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&uppercase, "uppercase", "u", false, "Print name in uppercase")

	return cmd
}
