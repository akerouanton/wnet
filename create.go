package main

import (
	"fmt"

	"github.com/Microsoft/hcsshim/hcn"
	"github.com/spf13/cobra"
)

type createFlags struct {
	networkName      string
	networkType      string
	ipamType         string
	subnetPrefix     string
	routeGateway     string
	routeDestination string
}

func NewCreateCmd() *cobra.Command {
	flags := createFlags{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new network",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(flags)
		},
	}

	cmd.Flags().StringVar(&flags.networkName, "network-name", "", "")
	cmd.Flags().StringVar(&flags.networkType, "network-type", "nat", "")
	cmd.Flags().StringVar(&flags.ipamType, "ipam-type", "static", "")
	cmd.Flags().StringVar(&flags.subnetPrefix, "subnet-prefix", "", "")
	cmd.Flags().StringVar(&flags.routeGateway, "route-gateway", "", "")
	cmd.Flags().StringVar(&flags.routeDestination, "route-destination", "", "")

	return cmd
}

func runCreate(flags createFlags) error {
	net := &hcn.HostComputeNetwork{
		Name: flags.networkName,
		Type: hcn.NetworkType(flags.networkType),
		Ipams: []hcn.Ipam{{
			Type: flags.ipamType,
			Subnets: []hcn.Subnet{{
				IpAddressPrefix: flags.subnetPrefix,
				Routes: []hcn.Route{
					{NextHop: flags.routeGateway, DestinationPrefix: flags.routeDestination},
				},
			}},
		}},
		SchemaVersion: hcn.V2SchemaVersion(),
	}

	net, err := net.Create()
	if err != nil {
		return fmt.Errorf("HostComputeNetwork.Create: %w", err)
	}

	fmt.Println(net.Id)

	return nil
}
