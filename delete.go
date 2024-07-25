package main

import (
	"fmt"

	"github.com/Microsoft/hcsshim/hcn"
	"github.com/spf13/cobra"
)

func NewRmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove a network",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRm(args[0])
		},
		Args: cobra.ExactArgs(1),
	}

	return cmd
}

func runRm(networkName string) error {
	net, err := hcn.GetNetworkByName(networkName)
	if err != nil {
		return fmt.Errorf("GetNetworkByName: %w", err)
	}

	return net.Delete()

}
