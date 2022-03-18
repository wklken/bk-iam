/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package util_test

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"

	"iam/pkg/util"
)

var _ = Describe("String", func() {
	Describe("TruncateBytes", func() {
		s := []byte("helloworld")

		DescribeTable("TruncateBytes cases", func(expected []byte, truncatedSize int) {
			assert.Equal(GinkgoT(), expected, util.TruncateBytes(s, truncatedSize))
		},
			Entry("truncated size less than real size", []byte("he"), 2),
			Entry("truncated size equals to real size", s, 10),
			Entry("truncated size greater than real size", s, 20),
		)
	})

	Describe("TruncateBytesToString", func() {
		s := []byte("helloworld")
		sStr := string(s)

		DescribeTable("TruncateBytesToString cases", func(expected string, truncatedSize int) {
			assert.Equal(GinkgoT(), expected, util.TruncateBytesToString(s, truncatedSize))
		},
			Entry("truncated size less than real size", "he", 2),
			Entry("truncated size equals to real size", sStr, 10),
			Entry("truncated size greater than real size", sStr, 20),
		)
	})
})

func BenchmarkStringSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("bk_job:script_execute:%s", "12345678")
	}
}

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "bk_job:script_execute:" + "12345678"
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s strings.Builder
		s.WriteString("bk_job:script_execute:")
		s.WriteString("12345678")
		_ = s.String()
	}
}

func BenchmarkStringBuilderWithPool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return new(strings.Builder)
		},
	}

	for i := 0; i < b.N; i++ {
		// var s strings.Builder
		s := pool.Get().(*strings.Builder)

		s.WriteString("bk_job:script_execute:")
		s.WriteString("12345678")
		_ = s.String()

		s.Reset()
		pool.Put(s)
	}
}

func BenchmarkIntStringSprintfD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", 123456)
	}
}

func BenchmarkIntToStringItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(123456)
	}
}

func BenchmarkInt64StringSprintfD(b *testing.B) {
	x := int64(123456)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", x)
	}
}

func BenchmarkInt64ToStringFormatInt(b *testing.B) {
	x := int64(123456)
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(x, 10)
	}
}

func BenchmarkInt64ToStringUseMap(b *testing.B) {
	m := map[int64]string{}
	largest := int64(1000000)
	var i int64
	for i = 0; i < largest; i++ {
		m[i] = strconv.FormatInt(i, 10)
	}

	// TODO: we can build a global map for this!!!!!

	x := int64(123456)
	for i := 0; i < b.N; i++ {
		_ = m[x]
		// strconv.FormatInt(x, 10)
	}
}

func BenchmarkStringAddInt64Sprintf(b *testing.B) {
	x := int64(123456)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s:%d", "abc", x)
	}
}

func BenchmarkStringAddInt64FormatInt(b *testing.B) {
	x := int64(123456)
	for i := 0; i < b.N; i++ {
		_ = "abc" + strconv.FormatInt(x, 10)
	}
}
