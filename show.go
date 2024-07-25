package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewShowCmd(driver string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show a complete network config",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			switch driver {
			case driverHCN:
				err = showHcnNetwork(args[0])
			case driverHNS:
				err = showHnsNetwork(args[0])
			default:
				err = fmt.Errorf("unsupported driver %q", driver)
			}

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	cmd.Args = cobra.ExactArgs(1)

	return cmd
}

func showHcnNetwork(nwName string) error {
	nw, err := HcnGetNetwork(nwName)
	if err != nil {
		return fmt.Errorf("failed to get HCN network %q: %w", nwName, err)
	}

	PrintHcnNetwork(nw)

	return nil
}

func showHnsNetwork(nwName string) error {
	nw, err := HnsGetNetwork(nwName)
	if err != nil {
		return fmt.Errorf("failed to get HNS network %q: %w", nwName, err)
	}

	PrintHnsNetwork(nw)

	return nil
}
