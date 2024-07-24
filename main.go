package main

import (
	"fmt"
	"os"

	"github.com/Microsoft/hcsshim/hcn"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "wnet",
		Short: "wnet is a small utility that can be used to manage HNS / HCN networks",
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List existing networks",
		RunE:  runList,
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runList(cmd *cobra.Command, args []string) error {
	nets, err := hcn.ListNetworks()
	if err != nil {
		return fmt.Errorf("ListNetworks: %w", err)
	}

	for _, net := range nets {
		fmt.Printf("Network %q (ID: %s):\n", net.Name, net.Id)
		fmt.Printf("\tType: %s\n", net.Type)

		fmt.Printf("\tPolicies:\n")
		for _, policy := range net.Policies {
			fmt.Printf("\t\t- Type: %s\n", policy.Type)
			fmt.Printf("\t\t  Settings: %s\n", policy.Settings)
		}

		fmt.Printf("\tMacPool:\n")
		for _, macRange := range net.MacPool.Ranges {
			fmt.Printf("\t\t- Range Start: %s -- Range End: %s\n", macRange.StartMacAddress, macRange.EndMacAddress)
		}

		fmt.Printf("\tDNS:\n")
		fmt.Printf("\t\t- Domain: %s\n", net.Dns.Domain)
		fmt.Printf("\t\t  Search: %v\n", net.Dns.Search)
		fmt.Printf("\t\t  ServerList: %v\n", net.Dns.ServerList)
		fmt.Printf("\t\t  Options: %v\n", net.Dns.Options)

		fmt.Printf("\tIPAMs:\n")
		for _, ipam := range net.Ipams {
			fmt.Printf("\t\t- Type: %s\n", ipam.Type)
			fmt.Printf("\t\t  Subnets: %v\n", ipam.Subnets)
		}

		fmt.Printf("\tFlags: %s (%d)\n", netFlagsToStr(net.Flags), net.Flags)
		fmt.Printf("\tHealth:\n")
		fmt.Printf("\t\tData: %v\n", net.Health.Data)

		fmt.Printf("\t\tExtra:\n")
		fmt.Printf("\t\t\tResources: %s\n", net.Health.Extra.Resources)
		fmt.Printf("\t\t\tSharedContainers: %s\n", net.Health.Extra.SharedContainers)
		fmt.Printf("\t\t\tLayeredOn: %s\n", net.Health.Extra.LayeredOn)
		fmt.Printf("\t\t\tSwitchGuid: %s\n", net.Health.Extra.SwitchGuid)
		fmt.Printf("\t\t\tUtilityVM: %s\n", net.Health.Extra.UtilityVM)
		fmt.Printf("\t\t\tVirtualMachine: %s\n", net.Health.Extra.VirtualMachine)

		fmt.Printf("\t\tSchemaVersion: %+v\n\n", net.SchemaVersion)
	}

	return nil
}

func netFlagsToStr(netflags hcn.NetworkFlags) string {
	switch netflags {
	case hcn.None:
		return "None"
	case hcn.EnableNonPersistent:
		return "EnableNonPersistent"
	case hcn.DisableHostPort:
		return "DisableHostPort"
	case hcn.EnableIov:
		return "EnableIov"
	default:
		return "Unknown flag"
	}
}
