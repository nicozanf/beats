// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package auditd

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "auditd", asset.ModuleFieldsPri, AssetAuditd); err != nil {
		panic(err)
	}
}

// AssetAuditd returns asset data.
// This is the base64 encoded gzipped contents of module/auditd.
func AssetAuditd() string {
	return "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"
}