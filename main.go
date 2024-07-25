package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	driverHNS = "hns"
	driverHCN = "hcn"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "wnet",
		Short: "wnet is a small utility that can be used to manage HNS / HCN networks",
	}

	rootCmd.AddCommand(newDriverCmd(driverHNS))
	rootCmd.AddCommand(newDriverCmd(driverHCN))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newDriverCmd(driver string) *cobra.Command {
	driverCmd := &cobra.Command{
		Use:   driver,
		Short: "Manage networks using the " + strings.ToUpper(driver) + " API",
	}
	driverCmd.AddCommand(NewListCmd(driver))
	driverCmd.AddCommand(NewShowCmd(driver))
	driverCmd.AddCommand(NewCreateCmd(driver))
	driverCmd.AddCommand(NewDeleteCmd(driver))

	return driverCmd
}
