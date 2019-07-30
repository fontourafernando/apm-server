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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/apm-server/beater/request"
)

func TestPanicHandler(t *testing.T) {

	setupContext := func() *request.Context {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		c := &request.Context{}
		c.Reset(w, r)
		return c
	}

	t.Run("NoPanic", func(t *testing.T) {
		h := func(c *request.Context) { c.WriteHeader(http.StatusAccepted) }
		c := setupContext()
		panicHandler(h)(c)
		require.Equal(t, http.StatusAccepted, c.StatusCode)
		assert.Empty(t, c.Err)
		assert.Empty(t, c.Stacktrace)
	})

	t.Run("HandlePanic", func(t *testing.T) {
		h := func(c *request.Context) { panic("panic xyz") }
		c := setupContext()
		panicHandler(h)(c)
		require.Equal(t, http.StatusInternalServerError, c.StatusCode)
		assert.Contains(t, c.Err, "panic xyz")
		assert.NotNil(t, c.Stacktrace)
	})

}
