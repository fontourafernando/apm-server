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

package beater

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/elastic/apm-server/beater/beatertest"
	"github.com/elastic/apm-server/beater/request"
)

func TestRumHandler_MonitoringMiddleware(t *testing.T) {
	h, err := rumHandler(DefaultConfig(beatertest.MockBeatVersion()), beatertest.NilReporter)
	require.NoError(t, err)
	c, _ := beatertest.DefaultContextWithResponseRecorder()

	// send GET request resulting in 403 Forbidden error as RUM is disabled by default
	expected := map[request.ResultID]int{
		request.IDRequestCount:            1,
		request.IDResponseCount:           1,
		request.IDResponseErrorsCount:     1,
		request.IDResponseErrorsForbidden: 1}

	equal, result := beatertest.CompareMonitoringInt(h, c, expected, serverMetrics, IntakeResultIDToMonitoringInt)
	assert.True(t, equal, result)
}

//TODO: add more integration tests for the actual middleware
