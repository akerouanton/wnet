package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "wnet",
		Short: "wnet is a small utility that can be used to manage HNS / HCN networks",
	}

	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(NewCreateCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
