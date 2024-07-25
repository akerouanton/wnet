package main

import (
	"github.com/spf13/cobra"
)

func NewDeleteCmd(driver string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a network",
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
