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
	"fmt"
	"net/http"
	"testing"

	"github.com/elastic/beats/libbeat/version"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/apm-server/beater/beatertest"
)

//TODO: check if it is breaking to return 404 body not always in plain text
func TestRootHandler(t *testing.T) {
	t.Run("404", func(t *testing.T) {
		c, w := beatertest.ContextWithResponseRecorder(http.MethodGet, "/abc/xyz")
		RootHandler()(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, `{"error":"404 page not found"}`+"\n", w.Body.String())
	})

	t.Run("ok", func(t *testing.T) {
		c, w := beatertest.ContextWithResponseRecorder(http.MethodGet, "/")
		RootHandler()(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	t.Run("unauthorized", func(t *testing.T) {
		c, w := beatertest.ContextWithResponseRecorder(http.MethodGet, "/")
		RootHandler()(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	t.Run("authorized", func(t *testing.T) {
		c, w := beatertest.ContextWithResponseRecorder(http.MethodGet, "/")
		c.Authorized = true
		RootHandler()(c)

		assert.Equal(t, http.StatusOK, w.Code)
		body := fmt.Sprintf("{\"build_date\":\"0001-01-01T00:00:00Z\",\"build_sha\":\"%s\",\"version\":\"%s\"}\n",
			version.Commit(), version.GetDefaultVersion())
		assert.Equal(t, body, w.Body.String())
	})
}
