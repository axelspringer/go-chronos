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

const (
	// ChronosAPI is a constant representing the Chronos REST API entrypoint
	ChronosAPI = "v1/scheduler/"
	// ChronosAPIJobs is a constant representing the jobs resource
	ChronosAPIJobs = "jobs"
	// ChronosAPIJobsSearch is a constant representing the search resource
	ChronosAPIJobsSearch = "jobs/search"
	// ChronosAPIJob is a constant representing the job resource
	ChronosAPIJob = "job/"
	// ChronosAPIJobTaskKill is a constant representing the task resource
	ChronosAPIJobTaskKill = "task/kill/"
	// ChronosAPIJobSuccess is a constant representing the jobs success resource
	ChronosAPIJobSuccess = "success/"
	// ChronosAPIIso8601 is a constant representing the spec for date format
	ChronosAPIIso8601 = "iso8601"
)
