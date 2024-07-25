package main

import (
	"github.com/spf13/cobra"
)

func NewListCmd(driver string) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List existing networks",
		RunE: func(cmd *cobra.Command, args []string) error {
			if driver == driverHCN {
				return runHcnList()
			}
			return runHnsList()
		},
	}
}

func runHcnList() error {
	nws, err := HcnListNetworks()
	if err != nil {
		return err
	}

	for _, nw := range nws {
		PrintHcnNetwork(nw)
	}

	return nil
}

func runHnsList() error {
	nws, err := HnsListNetworks()
	if err != nil {
		return err
	}

	for _, nw := range nws {
		PrintHnsNetwork(nw)
	}

	return nil
}
