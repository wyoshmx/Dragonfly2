/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package timeutils provides utilities supplementing the standard 'time' package.
package structutils

import (
	"encoding/json"
)

func StructToMap(t interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}
