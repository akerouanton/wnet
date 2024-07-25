package main

import (
	"encoding/json"
	"fmt"

	"github.com/Microsoft/hcsshim"
)

func HnsListNetworks() ([]hcsshim.HNSNetwork, error) {
	return hcsshim.HNSListNetworkRequest("GET", "", "")
}

func HnsGetNetwork(nwName string) (*hcsshim.HNSNetwork, error) {
	return hcsshim.GetHNSNetworkByName(nwName)
}

func HnsCreateNetwork(nw *hcsshim.HNSNetwork) (string, error) {
	config, err := json.Marshal(nw)
	if err != nil {
		return "", fmt.Errorf("marshal HNSNetwork: %w", err)
	}

	resp, err := hcsshim.HNSNetworkRequest("POST", "", string(config))
	if err != nil {
		return "", fmt.Errorf("create HNSNetwork: %w", err)
	}

	return resp.Id, nil
}

func HnsDeleteNetwork(netName string) error {
	hnsNet, err := hcsshim.GetHNSNetworkByName(netName)
	if err != nil {
		return fmt.Errorf("GetHNSNetworkByName: %v", err)
	}

	if _, err = hcsshim.HNSNetworkRequest("DELETE", hnsNet.Id, ""); err != nil {
		return fmt.Errorf("delete HNS network: %v", err)
	}

	return nil
}

func PrintHnsNetwork(nw *hcsshim.HNSNetwork) {
	fmt.Printf("Network %q (ID: %s):\n", nw.Name, nw.Id)
	fmt.Printf("\tType: %s\n", nw.Type)
	fmt.Printf("\tNetwork adapter: %s\n", nw.NetworkAdapterName)
	fmt.Printf("\tSource MAC: %s\n", nw.SourceMac)
	fmt.Printf("\tManagement IP: %s\n", nw.ManagementIP)

	fmt.Printf("\tPolicies:\n")
	for _, policy := range nw.Policies {
		fmt.Printf("\t\t- %s\n", policy)
	}

	fmt.Printf("\tMacPool:\n")
	for _, macRange := range nw.MacPools {
		fmt.Printf("\t\t- Range Start: %s\n", macRange.StartMacAddress)
		fmt.Printf("\t\t  Range End:   %s\n", macRange.EndMacAddress)
	}

	fmt.Printf("\tSubnets:\n")
	for _, subnet := range nw.Subnets {
		fmt.Printf("\t\t- Prefix: %s\n", subnet.AddressPrefix)
		fmt.Printf("\t\t  Gateway: %+v\n", subnet.GatewayAddress)
		fmt.Printf("\t\t  Policies: %s\n", subnet.Policies)
	}

	fmt.Printf("\tDNS:\n")
	fmt.Printf("\t\t- Suffix: %s\n", nw.DNSSuffix)
	fmt.Printf("\t\t  Server List: %v\n", nw.DNSServerList)
	fmt.Printf("\t\t  Server Compartment: %d\n", nw.DNSServerCompartment)
	fmt.Printf("\t\t  Automatic DNS: %t\n", nw.AutomaticDNS)
}
