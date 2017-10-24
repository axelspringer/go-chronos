// Copyright 2017 Axel Springer SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chronos

import (
	"net/http"

	"github.com/dghubble/sling"
)

// New returns a new Client
func New(chronosURL string, httpClient *http.Client) *Client {
	return mustNew(&chronosURL, httpClient)
}

// mustNew wraps the creation of a new Client
func mustNew(chronosURL *string, httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(*chronosURL).Path(ChronosAPI)

	return &Client{
		sling: base,
		Jobs:  newJobsService(base.New()),
		Job:   newJobService(base.New()),
	}
}
