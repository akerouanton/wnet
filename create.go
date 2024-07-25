package main

import (
	"fmt"
	"github.com/Microsoft/hcsshim"
	"gopkg.in/yaml.v3"
	"os"

	"github.com/Microsoft/hcsshim/hcn"
	"github.com/spf13/cobra"
)

func NewCreateCmd(driver string) *cobra.Command {
	var debug bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new network",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			switch driver {
			case driverHCN:
				err = runHcnCreate(args[0], debug)
			case driverHNS:
				err = runHnsCreate(args[0], debug)
			default:
				err = fmt.Errorf("unsupported driver %s", driver)
			}

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	cmd.Args = cobra.ExactArgs(1)
	cmd.Flags().BoolVar(&debug, "debug", false, "show the network spec before sending it to HNS/HCN")

	return cmd
}

func runHnsCreate(specFile string, debug bool) error {
	spec, err := os.ReadFile(specFile)
	if err != nil {
		return fmt.Errorf("failed to read spec file: %w", err)
	}

	nw := &hcsshim.HNSNetwork{}
	if err := yaml.Unmarshal(spec, nw); err != nil {
		return fmt.Errorf("failed to parse network spec: %w", err)
	}

	if debug {
		PrintHnsNetwork(nw)
	}

	nwId, err := HnsCreateNetwork(nw)
	if err != nil {
		return fmt.Errorf("failed to create network: %w", err)
	}

	fmt.Println(nwId)
	return nil
}

func runHcnCreate(specFile string, debug bool) error {
	spec, err := os.ReadFile(specFile)
	if err != nil {
		return fmt.Errorf("failed to read spec file: %w", err)
	}

	nw := &hcn.HostComputeNetwork{SchemaVersion: hcn.V2SchemaVersion()}
	if err := yaml.Unmarshal(spec, nw); err != nil {
		return fmt.Errorf("failed to parse network spec: %w", err)
	}

	if debug {
		PrintHcnNetwork(nw)
	}

	createdNw, err := nw.Create()
	if err != nil {
		return fmt.Errorf("failed to create network: %w", err)
	}

	fmt.Println(createdNw.Id)
	return nil
}
