/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package condition

import (
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Any", func() {
	var c *AnyCondition
	BeforeEach(func() {
		c = &AnyCondition{
			baseCondition{
				Key:   "ok",
				Value: []interface{}{"a", "b"},
			},
		}
	})

	It("new", func() {
		condition, err := newAnyCondition("ok", []interface{}{"a", "b"})
		assert.NoError(GinkgoT(), err)
		assert.NotNil(GinkgoT(), condition)
	})

	It("New", func() {
		c := NewAnyCondition()
		assert.Equal(GinkgoT(), "Any", c.GetName())
		assert.Equal(GinkgoT(), []string{}, c.GetKeys())
	})

	It("GetName", func() {
		assert.Equal(GinkgoT(), "Any", c.GetName())
	})

	It("GetKeys", func() {
		assert.Empty(GinkgoT(), c.GetKeys())
	})

	It("Eval", func() {
		assert.True(GinkgoT(), c.Eval(intCtx(1)))
		assert.True(GinkgoT(), c.Eval(boolCtx(false)))
		assert.True(GinkgoT(), c.Eval(listCtx{1, 2}))
		assert.True(GinkgoT(), c.Eval(errCtx(1)))
	})

	It("Translate", func() {
		ec, err := c.Translate()
		assert.NoError(GinkgoT(), err)
		expected := map[string]interface{}{
			"op":    "any",
			"field": "ok",
			"value": []interface{}{"a", "b"},
		}
		assert.Equal(GinkgoT(), expected, ec)
	})
})
