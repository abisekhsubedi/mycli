package commands

import (
	"context"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "mycli",
		Short:         "A CLI application",
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(NewGreetCmd())

	return rootCmd
}

func Execute(ctx context.Context) error {
	return NewRootCmd().ExecuteContext(ctx)
}
