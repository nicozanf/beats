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

package datastorecluster

import (
	"github.com/vmware/govmomi/vim25/mo"

	"github.com/elastic/elastic-agent-libs/mapstr"
)

func (m *DatastoreClusterMetricSet) mapEvent(datastoreCluster mo.StoragePod, data *metricData) mapstr.M {
	event := mapstr.M{
		"name": datastoreCluster.Name,
		"id":   datastoreCluster.Self.Value,
		"capacity": mapstr.M{
			"bytes": datastoreCluster.Summary.Capacity,
		},
		"free_space": mapstr.M{
			"bytes": datastoreCluster.Summary.FreeSpace,
		},
		"datastore": mapstr.M{
			"names": data.assetNames.outputDsNames,
			"count": len(data.assetNames.outputDsNames),
		},
	}

	if len(data.triggeredAlarms) > 0 {
		event.Put("triggered_alarms", data.triggeredAlarms)
	}

	return event
}