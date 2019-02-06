package azureutil

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2016-09-01/network"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2016-01-01/storage"
)

// DeploymentContext contains references to various sources created and then
// used in creating other resources.
type DeploymentContext struct {
	VirtualNetworkExists   bool
	StorageAccount         *storage.AccountProperties
	PublicIPAddressID      string
	NetworkSecurityGroupID string
	SubnetID               string
	NetworkInterfaceID     string
	SSHPublicKey           string
	AvailabilitySetID      string
	FirewallRules          *[]network.SecurityRule
}
