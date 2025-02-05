/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package pdp

import (
	"errors"
	"reflect"

	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"

	"iam/pkg/abac/pdp/evaluation"
	"iam/pkg/abac/pdp/translate"
	pdptypes "iam/pkg/abac/pdp/types"
	"iam/pkg/abac/types"
	"iam/pkg/abac/types/request"
	"iam/pkg/logging/debug"
)

var _ = Describe("Entrance", func() {

	Describe("Eval", func() {
		var entry *debug.Entry
		var req *request.Request
		var ctl *gomock.Controller
		var patches *gomonkey.Patches
		BeforeEach(func() {
			//entry = debug.EntryPool.Get()
			ctl = gomock.NewController(GinkgoT())
			req = &request.Request{
				System: "test",
				Resources: []types.Resource{{
					System: "test",
				}},
			}

			patches = gomonkey.NewPatches()
			patches.ApplyMethod(reflect.TypeOf(req), "ValidateActionResource",
				func(_ *request.Request) bool {
					return true
				})
			patches.ApplyMethod(reflect.TypeOf(req), "HasSingleLocalResource",
				func(_ *request.Request) bool {
					return true
				})
		})
		AfterEach(func() {
			ctl.Finish()
			patches.Reset()
		})

		It("FillAction error", func() {
			patches.ApplyFunc(fillActionDetail, func(req *request.Request) error {
				return errors.New("fill action fail")
			})

			ok, err := Eval(req, entry, false)
			assert.False(GinkgoT(), ok)
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "fill action fail")
		})

		It("ValidateAction error", func() {
			patches.ApplyFunc(fillActionDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyMethod(reflect.TypeOf(req), "ValidateActionResource",
				func(_ *request.Request) bool {
					return false
				})

			ok, err := Eval(req, entry, false)
			assert.False(GinkgoT(), ok)
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "request resources not match action")
		})

		It("FillSubject error", func() {
			patches.ApplyFunc(fillActionDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyFunc(fillSubjectDetail, func(req *request.Request) error {
				return errors.New("fill subject fail")
			})

			ok, err := Eval(req, entry, false)
			assert.False(GinkgoT(), ok)
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "fill subject fail")

		})

		It("QueryPolicies error", func() {
			patches.ApplyFunc(fillActionDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyFunc(fillSubjectDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyFunc(queryPolicies, func(system string,
				subject types.Subject,
				action types.Action,
				withoutCache bool,
				entry *debug.Entry,
			) (policies []types.AuthPolicy, err error) {
				return nil, errors.New("queryPolicies fail")
			})

			ok, err := Eval(req, entry, false)
			assert.False(GinkgoT(), ok)
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "queryPolicies fail")
		})
		//
		It("ok, QueryPolicies single pass", func() {
			patches.ApplyFunc(fillActionDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyFunc(fillSubjectDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyFunc(queryPolicies, func(system string,
				subject types.Subject,
				action types.Action,
				withoutCache bool,
				entry *debug.Entry,
			) (policies []types.AuthPolicy, err error) {
				return []types.AuthPolicy{}, nil
			})
			patches.ApplyFunc(evaluation.EvalPolicies, func(
				ctx *pdptypes.ExprContext, policies []types.AuthPolicy,
			) (isPass bool, policyID int64, err error) {
				return true, 1, nil
			})

			ok, err := Eval(req, entry, false)
			assert.True(GinkgoT(), ok)
			assert.NoError(GinkgoT(), err)
		})

		It("fail, QueryPolicies single fail", func() {
			patches.ApplyFunc(fillActionDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyFunc(fillSubjectDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyFunc(queryPolicies, func(system string,
				subject types.Subject,
				action types.Action,
				withoutCache bool,
				entry *debug.Entry,
			) (policies []types.AuthPolicy, err error) {
				return []types.AuthPolicy{}, nil
			})
			patches.ApplyFunc(evaluation.EvalPolicies, func(
				ctx *pdptypes.ExprContext, policies []types.AuthPolicy,
			) (isPass bool, policyID int64, err error) {
				return false, -1, errors.New("eval fail")
			})

			ok, err := Eval(req, entry, false)
			assert.False(GinkgoT(), ok)
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "eval fail")
		})

		It("fail, QueryPolicies filter error", func() {
			patches.ApplyFunc(fillActionDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyFunc(fillSubjectDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyMethod(reflect.TypeOf(req), "HasSingleLocalResource",
				func(_ *request.Request) bool {
					return false
				})
			patches.ApplyFunc(queryPolicies, func(system string,
				subject types.Subject,
				action types.Action,
				withoutCache bool,
				entry *debug.Entry,
			) (policies []types.AuthPolicy, err error) {
				return []types.AuthPolicy{}, nil
			})
			patches.ApplyFunc(filterPoliciesByEvalResources, func(
				r *request.Request,
				policies []types.AuthPolicy,
			) (filteredPolicies []types.AuthPolicy, err error) {
				return nil, errors.New("test")
			})

			ok, err := Eval(req, entry, false)
			assert.False(GinkgoT(), ok)
			assert.Error(GinkgoT(), err, "test")
		})

		It("ok, QueryPolicies filter success", func() {
			patches.ApplyFunc(fillActionDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyFunc(fillSubjectDetail, func(req *request.Request) error {
				return nil
			})
			patches.ApplyMethod(reflect.TypeOf(req), "HasSingleLocalResource",
				func(_ *request.Request) bool {
					return false
				})
			patches.ApplyFunc(queryPolicies, func(system string,
				subject types.Subject,
				action types.Action,
				withoutCache bool,
				entry *debug.Entry,
			) (policies []types.AuthPolicy, err error) {
				return []types.AuthPolicy{}, nil
			})
			patches.ApplyFunc(filterPoliciesByEvalResources, func(
				r *request.Request,
				policies []types.AuthPolicy,
			) (filteredPolicies []types.AuthPolicy, err error) {
				return []types.AuthPolicy{{}}, nil
			})
			defer patches.Reset()

			ok, err := Eval(req, entry, false)
			assert.True(GinkgoT(), ok)
			assert.NoError(GinkgoT(), err)
		})
	})

	Describe("Query", func() {
		var entry *debug.Entry
		var req *request.Request
		var ctl *gomock.Controller
		var patches *gomonkey.Patches
		BeforeEach(func() {
			entry = debug.EntryPool.Get()
			ctl = gomock.NewController(GinkgoT())
			req = &request.Request{
				System: "test",
				Resources: []types.Resource{{
					System: "test",
				}},
			}
		})
		AfterEach(func() {
			ctl.Finish()
			patches.Reset()
		})

		It("filter error", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return nil, errors.New("test")
			})

			expr, err := Query(req, entry, false, false)
			assert.Nil(GinkgoT(), expr)
			assert.Error(GinkgoT(), err)
		})

		It("filter empty", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return []types.AuthPolicy{}, nil
			})

			expr, err := Query(req, entry, false, false)
			assert.Equal(GinkgoT(), expr, EmptyPolicies)
			assert.NoError(GinkgoT(), err)
		})

		It("get resourceType error", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return []types.AuthPolicy{{}}, nil
			})
			patches.ApplyMethod(reflect.TypeOf(req), "GetQueryResourceTypes",
				func(_ *request.Request) ([]types.ActionResourceType, error) {
					return nil, errors.New("test")
				})

			expr, err := Query(req, entry, false, false)
			assert.Nil(GinkgoT(), expr)
			assert.Error(GinkgoT(), err, "test")
		})

		It("translate error", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return []types.AuthPolicy{{}}, nil
			})
			patches.ApplyMethod(reflect.TypeOf(req), "GetQueryResourceTypes",
				func(_ *request.Request) ([]types.ActionResourceType, error) {
					return []types.ActionResourceType{}, nil
				})
			patches.ApplyFunc(translate.PoliciesTranslate, func(policies []types.AuthPolicy,
				resourceTypes []types.ActionResourceType,
			) (map[string]interface{}, error) {
				return nil, errors.New("test")
			})

			expr, err := Query(req, entry, false, false)
			assert.Nil(GinkgoT(), expr)
			assert.Error(GinkgoT(), err, "test")
		})

		It("ok", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return []types.AuthPolicy{{}}, nil
			})
			patches.ApplyMethod(reflect.TypeOf(req), "GetQueryResourceTypes",
				func(_ *request.Request) ([]types.ActionResourceType, error) {
					return []types.ActionResourceType{}, nil
				})
			patches.ApplyFunc(translate.PoliciesTranslate, func(policies []types.AuthPolicy,
				resourceTypes []types.ActionResourceType,
			) (map[string]interface{}, error) {
				return map[string]interface{}{}, nil
			})

			expr, err := Query(req, entry, false, false)
			assert.Equal(GinkgoT(), expr, map[string]interface{}{})
			assert.NoError(GinkgoT(), err)

		})

	})

	Describe("QueryByExtResources", func() {
		var entry *debug.Entry
		var req *request.Request
		var ctl *gomock.Controller
		var patches *gomonkey.Patches
		BeforeEach(func() {
			entry = debug.EntryPool.Get()
			ctl = gomock.NewController(GinkgoT())
			req = &request.Request{
				System: "test",
				Resources: []types.Resource{{
					System: "test",
				}},
			}
		})
		AfterEach(func() {
			ctl.Finish()
			patches.Reset()
		})

		It("filter error", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return nil, errors.New("test")
			})

			expr, resources, err := QueryByExtResources(req, []types.ExtResource{}, entry, false)
			assert.Nil(GinkgoT(), expr)
			assert.Nil(GinkgoT(), resources)
			assert.Error(GinkgoT(), err)
		})

		It("filter empty", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return []types.AuthPolicy{}, nil
			})

			expr, resources, err := QueryByExtResources(req, []types.ExtResource{}, entry, false)
			assert.Equal(GinkgoT(), expr, EmptyPolicies)
			assert.Equal(GinkgoT(), resources, []types.ExtResourceWithAttribute{})
			assert.Nil(GinkgoT(), err)
		})

		It("query error", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return []types.AuthPolicy{{}}, nil
			})
			patches.ApplyFunc(queryExtResourceAttrs, func(
				resource *types.ExtResource,
				policies []types.AuthPolicy,
			) (resources []map[string]interface{}, err error) {
				return nil, errors.New("test")
			})

			expr, resources, err := QueryByExtResources(req, []types.ExtResource{{}}, entry, false)
			assert.Nil(GinkgoT(), expr)
			assert.Nil(GinkgoT(), resources)
			assert.Error(GinkgoT(), err)
		})

		It("get ResourceType error", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return []types.AuthPolicy{{}}, nil
			})
			patches.ApplyMethod(reflect.TypeOf(req), "GetQueryResourceTypes",
				func(_ *request.Request) ([]types.ActionResourceType, error) {
					return nil, errors.New("test")
				})

			expr, resources, err := QueryByExtResources(req, []types.ExtResource{}, entry, false)
			assert.Nil(GinkgoT(), expr)
			assert.Nil(GinkgoT(), resources)
			assert.Error(GinkgoT(), err)
		})

		It("translate error", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return []types.AuthPolicy{{}}, nil
			})
			patches.ApplyMethod(reflect.TypeOf(req), "GetQueryResourceTypes",
				func(_ *request.Request) ([]types.ActionResourceType, error) {
					return []types.ActionResourceType{}, nil
				})
			patches.ApplyFunc(translate.PoliciesTranslate, func(policies []types.AuthPolicy,
				resourceTypes []types.ActionResourceType,
			) (map[string]interface{}, error) {
				return nil, errors.New("test")
			})

			expr, resources, err := QueryByExtResources(req, []types.ExtResource{}, entry, false)
			assert.Nil(GinkgoT(), expr)
			assert.Nil(GinkgoT(), resources)
			assert.Error(GinkgoT(), err, "test")
		})

		It("ok", func() {
			patches = gomonkey.ApplyFunc(queryFilterPolicies, func(
				r *request.Request,
				entry *debug.Entry,
				willCheckRemoteResource, // 是否检查请求的外部依赖资源完成性
				withoutCache bool,
			) ([]types.AuthPolicy, error) {
				return []types.AuthPolicy{{}}, nil
			})
			patches.ApplyMethod(reflect.TypeOf(req), "GetQueryResourceTypes",
				func(_ *request.Request) ([]types.ActionResourceType, error) {
					return []types.ActionResourceType{}, nil
				})
			patches.ApplyFunc(translate.PoliciesTranslate, func(policies []types.AuthPolicy,
				resourceTypes []types.ActionResourceType,
			) (map[string]interface{}, error) {
				return map[string]interface{}{}, nil
			})

			expr, resources, err := QueryByExtResources(req, []types.ExtResource{}, entry, false)
			assert.Equal(GinkgoT(), expr, map[string]interface{}{})
			assert.Equal(GinkgoT(), resources, []types.ExtResourceWithAttribute{})
			assert.NoError(GinkgoT(), err)
		})

	})

})
