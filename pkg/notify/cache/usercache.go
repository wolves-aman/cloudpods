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

package cache

import (
	"context"
	"database/sql"
	"fmt"

	"yunion.io/x/pkg/errors"

	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/auth"
	"yunion.io/x/onecloud/pkg/util/httputils"
)

type SUserCacheManager struct {
	db.SUserCacheManager
}

type SUser struct {
	db.SUser
}

func (user *SUser) GetModelManager() db.IModelManager {
	return UserCacheManager
}

var UserCacheManager *SUserCacheManager

func RegistUserCredCacheUpdater() {
	auth.RegisterAuthHook(onAuthCompleteUpdateCache)
}

func onAuthCompleteUpdateCache(userCred mcclient.TokenCredential) {
	UserCacheManager.updateUserCache(userCred)
}

func (ucm *SUserCacheManager) updateUserCache(userCred mcclient.TokenCredential) {
	ucm.Save(context.Background(), userCred.GetUserId(), userCred.GetUserName(),
		userCred.GetDomainId(), userCred.GetDomainName())
}

func (ucm *SUserCacheManager) dealErrFromKeystone(err error) error {
	if err != nil {
		if je, ok := err.(*httputils.JSONClientError); ok && je.Code == 404 {
			return sql.ErrNoRows
		} else {
			return errors.Wrap(err, "fetch User info from keystone")
		}
	}
	return nil
}

func (ucm *SUserCacheManager) FetchUsersByIDs(ctx context.Context, ids []string) (map[string]SUser, error) {
	q := ucm.Query().In("id", ids)
	users := make([]SUser, 0)
	err := db.FetchModelObjects(ucm, q, &users)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]SUser)

	for i := range users {
		ret[users[i].Id] = users[i]
	}

	// check that id is exist
	for _, id := range ids {
		if _, ok := ret[id]; ok {
			continue
		}
		user, err := ucm.FetchUserFromKeystone(ctx, id)
		if err != nil {
			continue
		}
		ret[id] = SUser{*user}
	}
	return ret, nil
}

func (ucm *SUserCacheManager) FetchUserByIDOrName(ctx context.Context, idStr string) (*SUser, error) {
	user, err := ucm.SUserCacheManager.FetchUserByIdOrName(ctx, idStr)
	if err != nil {
		return nil, err
	}
	return &SUser{*user}, nil
}

func (ucm *SUserCacheManager) FetchUserLikeName(ctx context.Context, name string, noExpireCheck bool) ([]SUser,
	error) {

	if !noExpireCheck {
		// todo
		return nil, fmt.Errorf("FetchUserLikeName with check Not Implement")
	}
	q := ucm.Query().Like("name", "%"+name+"%")
	users := make([]SUser, 0)
	err := db.FetchModelObjects(ucm, q, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
