package client

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplaceordering/armmarketplaceordering"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// clientFactory implements ARMClientProvider interface.
type clientFactory struct {
}

// NewClientsProvider creates a new instance of ARMClientProvider.
func NewClientsProvider() ARMClientProvider {
	return clientFactory{}
}

func (c clientFactory) CreateResourceGroupsClient(connectConfig ConnectConfig) (*armresources.ResourceGroupsClient, error) {
	tokenCredential, err := connectConfig.CreateTokenCredential()
	if err != nil {
		return nil, err
	}
	factory, err := armresources.NewClientFactory(connectConfig.SubscriptionID, tokenCredential, nil)
	if err != nil {
		return nil, err
	}
	return factory.NewResourceGroupsClient(), nil
}

func (c clientFactory) CreateVirtualMachinesClient(connectConfig ConnectConfig) (*armcompute.VirtualMachinesClient, error) {
	tokenCredential, err := connectConfig.CreateTokenCredential()
	if err != nil {
		return nil, err
	}
	factory, err := armcompute.NewClientFactory(connectConfig.SubscriptionID, tokenCredential, nil)
	if err != nil {
		return nil, err
	}
	return factory.NewVirtualMachinesClient(), nil
}

func (c clientFactory) CreateNetworkInterfacesClient(connectConfig ConnectConfig) (*armnetwork.InterfacesClient, error) {
	tokenCredential, err := connectConfig.CreateTokenCredential()
	if err != nil {
		return nil, err
	}
	factory, err := armnetwork.NewClientFactory(connectConfig.SubscriptionID, tokenCredential, nil)
	if err != nil {
		return nil, err
	}
	return factory.NewInterfacesClient(), nil
}

func (c clientFactory) CreateSubnetClient(connectConfig ConnectConfig) (*armnetwork.SubnetsClient, error) {
	tokenCredential, err := connectConfig.CreateTokenCredential()
	if err != nil {
		return nil, err
	}
	factory, err := armnetwork.NewClientFactory(connectConfig.SubscriptionID, tokenCredential, nil)
	if err != nil {
		return nil, err
	}
	return factory.NewSubnetsClient(), nil
}

func (c clientFactory) CreateDisksClient(connectConfig ConnectConfig) (*armcompute.DisksClient, error) {
	tokenCredential, err := connectConfig.CreateTokenCredential()
	if err != nil {
		return nil, err
	}
	factory, err := armcompute.NewClientFactory(connectConfig.SubscriptionID, tokenCredential, nil)
	return factory.NewDisksClient(), nil
}

func (c clientFactory) CreateResourceGraphClient(connectConfig ConnectConfig) (*armresourcegraph.Client, error) {
	tokenCredential, err := connectConfig.CreateTokenCredential()
	if err != nil {
		return nil, err
	}
	factory, err := armresourcegraph.NewClientFactory(tokenCredential, nil)
	if err != nil {
		return nil, err
	}
	return factory.NewClient(), nil
}

func (c clientFactory) CreateImagesClient(connectConfig ConnectConfig) (*armcompute.ImagesClient, error) {
	tokenCredential, err := connectConfig.CreateTokenCredential()
	if err != nil {
		return nil, err
	}
	factory, err := armcompute.NewClientFactory(connectConfig.SubscriptionID, tokenCredential, nil)
	if err != nil {
		return nil, err
	}
	return factory.NewImagesClient(), nil
}

func (c clientFactory) CreateMarketPlaceAgreementsClient(connectConfig ConnectConfig) (*armmarketplaceordering.MarketplaceAgreementsClient, error) {
	tokenCredential, err := connectConfig.CreateTokenCredential()
	if err != nil {
		return nil, err
	}
	factory, err := armmarketplaceordering.NewClientFactory(connectConfig.SubscriptionID, tokenCredential, nil)
	if err != nil {
		return nil, err
	}
	return factory.NewMarketplaceAgreementsClient(), nil
}
