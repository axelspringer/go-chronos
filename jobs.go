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

// newJobsService returns a new JobsService
func newJobsService(sling *sling.Sling) *JobsService {
	return &JobsService{
		sling.Path(ChronosAPIJobs),
	}
}

// List returns the list of jobs
func (j *JobsService) List() (*[]Job, *http.Response, error) {
	jobs := new([]Job)
	chronosError := new(Error)
	resp, err := j.sling.New().Get("").Receive(jobs, chronosError)

	return jobs, resp, err
}

// Search is searching for a job in Chronos
// params: are the search parameters for the job
func (j *JobsService) Search(params *JobsSearchParams) (*[]Job, *http.Response, error) {
	jobs := new([]Job)
	ChronosError := new(Error)
	resp, err := j.sling.New().Get(ChronosAPIJobsSearch).QueryStruct(params).Receive(jobs, ChronosError)
	defer resp.Body.Close()

	return jobs, resp, err
}
