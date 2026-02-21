package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/abisekhsubedi/mycli/internal/commands"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	go func() {
		<-ctx.Done()
		if ctx.Err() == context.Canceled {
			return
		}
		fmt.Fprintln(os.Stderr, "\nReceived signal, shutting down...")
		os.Exit(1)
	}()

	if err := commands.Execute(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
