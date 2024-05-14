// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azure

import (
	"fmt"
	"strings"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SAppHybirdConnection struct {
	Id         string
	Name       string
	Type       string
	Properties struct {
		ServiceBusNamespace string
		RelayName           string
		Hostname            string
		Port                int
	}
}

func (self *SAppHybirdConnection) GetGlobalId() string {
	return strings.ToLower(self.Id)
}

func (self *SAppHybirdConnection) GetName() string {
	return self.Name
}

func (self *SAppHybirdConnection) GetHostname() string {
	return self.Properties.Hostname
}

func (self *SAppHybirdConnection) GetNamespace() string {
	return self.Properties.ServiceBusNamespace
}

func (self *SAppHybirdConnection) GetPort() int {
	return self.Properties.Port
}

func (self *SRegion) GetAppHybirConnections(appId string) ([]SAppHybirdConnection, error) {
	res := fmt.Sprintf("%s/hybridConnectionRelays", appId)
	resp, err := self.list_v2(res, "2023-12-01", nil)
	if err != nil {
		return nil, err
	}
	ret := []SAppHybirdConnection{}
	err = resp.Unmarshal(&ret, "value")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *SAppSite) GetHybirdConnections() ([]cloudprovider.IAppHybirdConnection, error) {
	connections, err := self.region.GetAppHybirConnections(self.Id)
	if err != nil {
		return nil, err
	}
	ret := []cloudprovider.IAppHybirdConnection{}
	for i := range connections {
		ret = append(ret, &connections[i])
	}
	return ret, nil
}
