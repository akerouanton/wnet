package main

import (
	"fmt"
	"strings"

	"github.com/Microsoft/hcsshim/hcn"
)

func HcnListNetworks() ([]hcn.HostComputeNetwork, error) {
	nets, err := hcn.ListNetworks()
	if err != nil {
		return []hcn.HostComputeNetwork{}, fmt.Errorf("HcnList: ListNetworks: %w", err)
	}

	return nets, nil
}

func HcnDeleteNetwork(netName string) error {
	nw, err := hcn.GetNetworkByName(netName)
	if err != nil {
		return fmt.Errorf("HcnDelete: GetNetworkByName: %w", err)
	}

	return nw.Delete()
}

func PrintHcnNetwork(nw hcn.HostComputeNetwork) {
	fmt.Printf("Network %q (ID: %s):\n", nw.Name, nw.Id)
	fmt.Printf("\tType: %s\n", nw.Type)

	fmt.Printf("\tPolicies:\n")
	for _, policy := range nw.Policies {
		fmt.Printf("\t\t- Type: %s\n", policy.Type)
		fmt.Printf("\t\t  Settings: %s\n", policy.Settings)
	}

	fmt.Printf("\tMacPool:\n")
	for _, macRange := range nw.MacPool.Ranges {
		fmt.Printf("\t\t- Range Start: %s -- Range End: %s\n", macRange.StartMacAddress, macRange.EndMacAddress)
	}

	fmt.Printf("\tDNS:\n")
	fmt.Printf("\t\t- Domain: %s\n", nw.Dns.Domain)
	fmt.Printf("\t\t  Search: %v\n", nw.Dns.Search)
	fmt.Printf("\t\t  ServerList: %v\n", nw.Dns.ServerList)
	fmt.Printf("\t\t  Options: %v\n", nw.Dns.Options)

	fmt.Printf("\tIPAMs:\n")
	for _, ipam := range nw.Ipams {
		fmt.Printf("\t\t- Type: %s\n", ipam.Type)
		fmt.Printf("\t\t  Subnets:\n")
		for _, subnet := range ipam.Subnets {
			fmt.Printf("\t\t\t- Prefix: %s\n", subnet.IpAddressPrefix)
			fmt.Printf("\t\t\t  Policies: %s\n", subnet.Policies)
			fmt.Printf("\t\t\t  Routes: %+v\n", subnet.Routes)
		}
	}

	fmt.Printf("\tFlags: %s (%d)\n", netFlagsToStr(nw.Flags), nw.Flags)
	fmt.Printf("\tHealth:\n")
	fmt.Printf("\t\tData: %v\n", nw.Health.Data)

	fmt.Printf("\t\tExtra:\n")
	fmt.Printf("\t\t\tResources: %s\n", nw.Health.Extra.Resources)
	fmt.Printf("\t\t\tSharedContainers: %s\n", nw.Health.Extra.SharedContainers)
	fmt.Printf("\t\t\tLayeredOn: %s\n", nw.Health.Extra.LayeredOn)
	fmt.Printf("\t\t\tSwitchGuid: %s\n", nw.Health.Extra.SwitchGuid)
	fmt.Printf("\t\t\tUtilityVM: %s\n", nw.Health.Extra.UtilityVM)
	fmt.Printf("\t\t\tVirtualMachine: %s\n", nw.Health.Extra.VirtualMachine)

	fmt.Printf("\t\tSchemaVersion: %+v\n\n", nw.SchemaVersion)
}

func netFlagsToStr(netflags hcn.NetworkFlags) string {
	flags := []string{}

	if netflags == hcn.None {
		flags = append(flags, "None")
	}
	if netflags&hcn.EnableNonPersistent > 0 {
		flags = append(flags, "EnableNonPersistent")
	}
	if netflags&hcn.DisableHostPort > 0 {
		flags = append(flags, "DisableHostPort")
	}
	if netflags&hcn.EnableIov > 0 {
		flags = append(flags, "EnableIov")
	}

	if len(flags) == 0 {
		flags = append(flags, fmt.Sprintf("Unknown flags (%d)", netflags))
	}

	return strings.Join(flags, ", ")
}
