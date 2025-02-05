/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package dao

import (
	"database/sql"
	"errors"
	"fmt"

	"iam/pkg/database"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock

// ModelChangeEvent 操作
type ModelChangeEvent struct {
	database.AllowBlankFields

	PK        int64  `db:"pk"` // 自增列
	Type      string `db:"type"`
	Status    string `db:"status"`
	SystemID  string `db:"system_id"`
	ModelType string `db:"model_type"`
	ModelID   string `db:"model_id"`
	ModelPK   int64  `db:"model_pk"`
}

// ModelChangeEventManager define the event crud for model change
type ModelChangeEventManager interface {
	GetByTypeModel(eventType, status, modelType string, modelPK int64) (ModelChangeEvent, error)
	ListByStatus(status string) ([]ModelChangeEvent, error)
	UpdateStatusByPK(pk int64, status string) error
	BulkCreate(modelChangeEvents []ModelChangeEvent) error
}

type modelChangeEventManager struct {
	DB *sqlx.DB
}

// NewModelChangeEventManager create a ModelChangeEventManager
func NewModelChangeEventManager() ModelChangeEventManager {
	return &modelChangeEventManager{
		DB: database.GetDefaultDBClient().DB,
	}
}

// GetByTypeModel ...
func (m *modelChangeEventManager) GetByTypeModel(eventType, status, modelType string,
	modelPK int64) (modelChangeEvent ModelChangeEvent, err error) {
	err = m.selectOne(&modelChangeEvent, eventType, status, modelType, modelPK)
	if errors.Is(err, sql.ErrNoRows) {
		return modelChangeEvent, nil
	}
	return
}

// ListByStatus ...
func (m *modelChangeEventManager) ListByStatus(status string) (modelChangeEvents []ModelChangeEvent, err error) {
	err = m.selectByStatus(&modelChangeEvents, status)
	if errors.Is(err, sql.ErrNoRows) {
		return modelChangeEvents, nil
	}
	return
}

// UpdateStatusByPK ...
func (m *modelChangeEventManager) UpdateStatusByPK(pk int64, status string) error {
	modelChangeEvent := ModelChangeEvent{PK: pk, Status: status}
	// 1. parse the set sql string and update data
	expr, data, err := database.ParseUpdateStruct(modelChangeEvent, modelChangeEvent.AllowBlankFields)
	if err != nil {
		return fmt.Errorf("parse update struct fail. %w", err)
	}

	// 2. build sql
	updatedSQL := "UPDATE model_change_event SET " + expr + " WHERE pk=:pk"

	return m.update(updatedSQL, data)
}

// BulkCreate ...
func (m *modelChangeEventManager) BulkCreate(modelChangeEvents []ModelChangeEvent) error {
	return m.insert(modelChangeEvents)
}

func (m *modelChangeEventManager) selectOne(modelChangeEvent *ModelChangeEvent, eventType, status, modelType string,
	modelPK int64) error {
	query := `SELECT
		pk,
		type,
		status,
		system_id,
		model_type,
		model_id,
		model_pk
		FROM model_change_event
		WHERE type = ?
		AND status = ?
		AND model_type = ?
		AND model_pk = ?
		LIMIT 1`
	return database.SqlxGet(m.DB, modelChangeEvent, query, eventType, status, modelType, modelPK)
}

func (m *modelChangeEventManager) selectByStatus(modelChangeEvents *[]ModelChangeEvent, status string) error {
	query := `SELECT
		pk,
		type,
		status,
		system_id,
		model_type,
		model_id,
		model_pk
		FROM model_change_event
		WHERE status=?`
	return database.SqlxSelect(m.DB, modelChangeEvents, query, status)
}

func (m *modelChangeEventManager) update(updatedSQL string, data map[string]interface{}) error {
	_, err := database.SqlxUpdate(m.DB, updatedSQL, data)
	if err != nil {
		return err
	}
	return nil
}
func (m *modelChangeEventManager) insert(modelChangeEvents []ModelChangeEvent) error {
	query := `INSERT INTO model_change_event (
		type,
		status,
		system_id,
		model_type,
		model_id,
		model_pk
	) VALUES (:type, :status, :system_id, :model_type, :model_id, :model_pk)`
	return database.SqlxBulkInsert(m.DB, query, modelChangeEvents)
}
