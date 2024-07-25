package main

import (
	"github.com/spf13/cobra"
)

func NewRmCmd(driver string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove a network",
		RunE: func(cmd *cobra.Command, args []string) error {
			if driver == driverHCN {
				return HcnDeleteNetwork(args[0])
			}
			return HnsDeleteNetwork(args[0])
		},
		Args: cobra.ExactArgs(1),
	}

	return cmd
}
