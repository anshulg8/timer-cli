package main

import (
	"os"
	"time"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "gifm",
		Short:        "Hello Gophers!",
		SilenceUsage: true,
	}

	cmd.AddCommand(PrintTimeCmd())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func PrintTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Hey Gophers! The current time is", prettyTime)
			return nil
		},
	}
}