/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package service

import (
	"database/sql"
	"errors"

	"github.com/TencentBlueKing/gopkg/errorx"
	"github.com/jmoiron/sqlx"

	"iam/pkg/database"
	"iam/pkg/database/dao"
)

// ErrConcurrencyConflict ...
var ErrConcurrencyConflict = errors.New("concurrency conflict")

// createOrUpdateGroupAuthType ...
func (s *groupService) createOrUpdateGroupAuthType(
	tx *sqlx.Tx,
	systemID string,
	groupPK, authType int64,
) (created bool, count int64, err error) {
	errorWrapf := errorx.NewLayerFunctionErrorWrapf(GroupSVC, "createOrUpdateGroupAuthType")

	groupSystemAuthType, err := s.authTypeManger.GetBySystemGroup(systemID, groupPK)
	if errors.Is(err, sql.ErrNoRows) {
		groupSystemAuthType = dao.GroupSystemAuthType{
			SystemID: systemID,
			GroupPK:  groupPK,
			AuthType: authType,
		}
		err = s.authTypeManger.CreateWithTx(tx, groupSystemAuthType)
		if err == nil {
			return true, 1, nil
		}

		if database.IsMysqlDuplicateEntryError(err) {
			return false, 0, ErrConcurrencyConflict
		}
	}

	if err != nil {
		err = errorWrapf(
			err,
			"groupSystemAuthTypeManager.GetBySystemGroup systemID=`%s` groupPK=`%d` fail",
			systemID,
			groupPK,
		)
		return false, 0, err
	}

	// 类型相同, 不需要更新
	if groupSystemAuthType.AuthType == authType {
		return false, 0, nil
	}

	groupSystemAuthType.AuthType = authType
	count, err = s.authTypeManger.UpdateWithTx(tx, groupSystemAuthType)
	if err != nil {
		err = errorWrapf(
			err, "groupSystemAuthTypeManager.UpdateWithTx groupSystemAuthType=`%+v` fail",
			groupSystemAuthType,
		)
		return false, 0, err
	}

	// 并发更新冲突
	if count == 0 {
		return false, 0, ErrConcurrencyConflict
	}

	return false, count, nil
}

// listGroupAuthSystem 查询group已授权的系统
func (s *groupService) listGroupAuthSystem(groupPK int64) ([]string, error) {
	errorWrapf := errorx.NewLayerFunctionErrorWrapf(GroupSVC, "listGroupAuthSystem")

	groupSystemAuthTypes, err := s.authTypeManger.ListByGroup(groupPK)
	if err != nil {
		err = errorWrapf(
			err,
			"groupSystemAuthTypeManager.ListByGroup groupPK=`%d` fail",
			groupPK,
		)
		return nil, err
	}

	systems := make([]string, 0, len(groupSystemAuthTypes))
	for _, groupSystemAuthType := range groupSystemAuthTypes {
		systems = append(systems, groupSystemAuthType.SystemID)
	}

	return systems, nil
}