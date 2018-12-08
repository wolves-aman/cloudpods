package cloudprovider

import (
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
)

type ICloudProviderFactory interface {
	GetProvider(providerId, providerName, url, account, secret string) (ICloudProvider, error)
	GetId() string
	ValidateChangeBandwidth(instanceId string, bandwidth int64) error
}

type ICloudProvider interface {
	GetSubAccounts() (jsonutils.JSONObject, error)
	GetId() string
	GetName() string
	GetIRegions() []ICloudRegion
	GetSysInfo() (jsonutils.JSONObject, error)
	IsPublicCloud() bool

	GetIRegionById(id string) (ICloudRegion, error)

	GetIHostById(id string) (ICloudHost, error)
	GetIVpcById(id string) (ICloudVpc, error)
	GetIStorageById(id string) (ICloudStorage, error)
	GetIStoragecacheById(id string) (ICloudStoragecache, error)

	GetBalance() (float64, error)
}

var providerTable map[string]ICloudProviderFactory

func init() {
	providerTable = make(map[string]ICloudProviderFactory)
}

func RegisterFactory(factory ICloudProviderFactory) {
	providerTable[factory.GetId()] = factory
}

func GetProviderDriver(provider string) (ICloudProviderFactory, error) {
	factory, ok := providerTable[provider]
	if ok {
		return factory, nil
	}
	log.Errorf("Provider %s not registerd", provider)
	return nil, fmt.Errorf("No such provider %s", provider)
}

func GetProvider(providerId, providerName, accessUrl, account, secret, provider string) (ICloudProvider, error) {
	driver, err := GetProviderDriver(provider)
	if err != nil {
		return nil, err
	}
	return driver.GetProvider(providerId, providerName, accessUrl, account, secret)
}

func IsSupported(provider string) bool {
	_, ok := providerTable[provider]
	return ok
}
