/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package util

import (
	"fmt"
	"strconv"
)

var sentryOn bool

const MaxPoolInteger = int64(1000000)

var int64ToStringMap map[int64]string // stringToInt64Map map[string]int64

// InitErrorReport init the sentryEnabled var
func InitErrorReport(sentryEnabled bool) {
	sentryOn = sentryEnabled

	int64ToStringMap = make(map[int64]string, MaxPoolInteger)

	var i int64
	k := MaxPoolInteger
	for i = 0; i < k; i++ {
		s := strconv.FormatInt(i, 10)
		int64ToStringMap[i] = s
		// stringToInt64Map[s] = i
	}
}

func ConvInt64ToString(i int64) string {
	if i < MaxPoolInteger {
		value, ok := int64ToStringMap[i]
		if !ok {
			fmt.Println("not exits", i)
		}
		return value
	}

	return strconv.FormatInt(i, 10)
}
