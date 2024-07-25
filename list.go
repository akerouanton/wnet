package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/tabwriter"
)

func NewListCmd(driver string) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List existing networks",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			switch driver {
			case driverHCN:
				err = runHcnList()
			case driverHNS:
				err = runHnsList()
			default:
				err = fmt.Errorf("unsupported driver %q", driver)
			}

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}

func runHcnList() error {
	nws, err := HcnListNetworks()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tSubnets")
	for _, nw := range nws {
		subnets := make([]string, 0)
		for _, ipam := range nw.Ipams {
			for _, subnet := range ipam.Subnets {
				subnets = append(subnets, subnet.IpAddressPrefix)
			}
		}

		fmt.Fprintf(w, "%s\t%s\t%s\n", nw.Id, nw.Name, strings.Join(subnets, ", "))
	}

	w.Flush()

	return nil
}

func runHnsList() error {
	nws, err := HnsListNetworks()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tSubnets")
	for _, nw := range nws {
		subnets := make([]string, 0, len(nw.Subnets))
		for _, subnet := range nw.Subnets {
			subnets = append(subnets, subnet.AddressPrefix)
		}
		fmt.Fprintf(w, "%s\t%s\t%s\n", nw.Id, nw.Name, strings.Join(subnets, ", "))
	}

	w.Flush()

	return nil
}
