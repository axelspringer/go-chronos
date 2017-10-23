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

// newJobService returns a new JobsService
func newJobService(sling *sling.Sling) *JobService {
	return &JobService{
		sling,
	}
}

// New is adding a new job in chronos
func (j *JobService) New(job *Job) (bool, *http.Response, error) {
	success := new(interface{})
	resp, err := j.sling.New().Post(ChronosAPIIso8601).BodyJSON(job).ReceiveSuccess(success)

	return resp.StatusCode == 204, resp, err
}

// Delete is deleting a job
// job: is the name of the job to delete
func (j *JobService) Delete(job string) (bool, *http.Response, error) {
	success := new(interface{})
	resp, err := j.sling.New().Path(ChronosAPIJob).Delete(job).ReceiveSuccess(success)

	return resp.StatusCode == 204, resp, err
}

// KillAllTasks is killing all related tasks for a job
// job: is the name of the job of whoms tasks to be killed
func (j *JobService) KillAllTasks(job string) (bool, *http.Response, error) {
	success := new(interface{})
	resp, err := j.sling.New().Path(ChronosAPIJob).Path(ChronosAPIJobTaskKill).Delete(job).ReceiveSuccess(success)

	return resp.StatusCode == 204, resp, err
}

// Start is manually starting a job
// job: is the name of the job to start
// params: are the parameters of the job to start
func (j *JobService) Start(job string, params *JobStartParams) (bool, *http.Response, error) {
	success := new(interface{})
	resp, err := j.sling.New().Path(ChronosAPIJob).Put(job).QueryStruct(params).ReceiveSuccess(success)

	return resp.StatusCode == 204, resp, err
}

// MarkSuccess is marking a job as successful
// job: is the name of the job to mark as success
func (j *JobService) MarkSuccess(job string) (bool, *http.Response, error) {
	success := new(string)
	resp, err := j.sling.New().Path(ChronosAPIJob).Path(ChronosAPIJobSuccess).Put(job).ReceiveSuccess(success)

	// TODO: parse response
	return true, resp, err
}
