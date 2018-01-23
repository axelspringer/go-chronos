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

// TODO: I guess there is a bug in the chronos implementation
// Assumption is that externalVolume contains a size parameter to control the ext vol
// Ref https://docs.mesosphere.com/1.7/usage/storage/external-storage/#create-an-application-with-external-volumes

package chronos

import "github.com/dghubble/sling"

// Client is a Chronos client for making Chronos API requests
type Client struct {
	sling *sling.Sling
	// Chronos API
	Jobs *JobsService
	Job  *JobService
}

// Error contains an error from Chronos API
type Error struct {
}

// Argument contains an job argument
type Argument string

// EnvironmentVariable contains a single env kv pair
type EnvironmentVariable struct {
	Key   string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// Fetch contains a remote executor fetch descriptor
type Fetch struct {
	URI        string `json:"uri,omitempty"`
	Cache      bool   `json:"cache,omitempty"`
	Extract    bool   `json:"extract,omitempty"`
	Executable bool   `json:"executable,omitempty"`
	OutputFile string `json:"output_file,omitempty"`
}

// JobsService provices methods for Chronos jobs operations
type JobsService struct {
	sling *sling.Sling
}

// JobService provides methods for Chronos job operations
type JobService struct {
	sling *sling.Sling
}

// JobsSearchParams is the parameters for job search
type JobsSearchParams struct {
	Name    string `url:"name,omitempty"`
	Command string `url:"command,omitempty"`
	Any     string `url:"any,omitempty"`
}

// Constraint contains a string tuple in the form ["LHS", "Comparator", "RHS"]
type Constraint []string

// Jobs contains a multitude of Job
type Jobs []Job

// Job contains a Chronos jobs description
type Job struct {
	Name                   string                `json:"name"`
	Command                string                `json:"command"`
	Shell                  bool                  `json:"shell,omitempty"`
	Epsilon                string                `json:"epsilon,omitempty"`
	Executor               string                `json:"executor,omitempty"`
	ExecutorFlags          string                `json:"executorFlags,omitempty"`
	Retries                int                   `json:"retries,omitempty"`
	Owner                  string                `json:"owner,omitempty"`
	OwnerName              string                `json:"ownerName,omitempty"`
	Description            string                `json:"description,omitempty"`
	Async                  bool                  `json:"async,omitempty"`
	SuccessCount           int                   `json:"successCount,omitempty"`
	ErrorCount             int                   `json:"errorCount,omitempty"`
	LastSuccess            string                `json:"lastSuccess,omitempty"`
	LastError              string                `json:"lastError,omitempty"`
	CPUs                   float64               `json:"cpus,omitempty"`
	Disk                   float64               `json:"disk,omitempty"`
	Mem                    float64               `json:"mem,omitempty"`
	Disabled               bool                  `json:"disabled,omitempty"`
	Concurrent             bool                  `json:"concurrent,omitempty"`
	SoftError              bool                  `json:"softError,omitempty"`
	FetchList              []Fetch               `json:"fetch,omitempty"`
	DataProcessingJobType  bool                  `json:"dataProcessingJobType,omitempty"`
	ErrorsSinceLastSuccess int                   `json:"errorsSinceLastSuccess,omitempty"`
	URIs                   []string              `json:"uris,omitempty"`
	EnvironmentVariables   []EnvironmentVariable `json:"environmentVariables,omitempty"`
	Arguments              []Argument            `json:"arguments,omitempty"`
	HighPriority           bool                  `json:"highPriority,omitempty"`
	RunAsUser              string                `json:"runAsUser,omitempty"`
	Container              *Container            `json:"container,omitempty"`
	Schedule               string                `json:"schedule,omitempty"`
	ScheduleTimeZone       string                `json:"scheduleTimeZone,omitempty"`
	Constraints            []Constraint          `json:"constraints,omitempty"`
	Parents                []string              `json:"parents,omitempty"`
	TaskInfoData           string                `json:"taskInfoData,omitempty"`
}

// JobStartParams contains the parameters to start a job
type JobStartParams struct {
	Arguments Argument `url:"arguments,omitempty"`
}

// PortMapping contains a port map and protocol specifications
type PortMapping struct {
	HostPort      int    `json:"hostPort,omitempty"`
	ContainerPort int    `json:"containerPort,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

// NetworkInfo contains container infos to NetworkInfo
type NetworkInfo struct {
	Name         string        `json:"name"`
	Protocol     string        `json:"protocol,omitempty"`
	Labels       []Parameter   `json:"labels,omitempty"`
	PortMappings []PortMapping `json:"portMappings,omitempty"`
}

// ExternalVolume contains external volume specification
type ExternalVolume struct {
	Name     string      `json:"name"`
	Provider string      `json:"provider,omitempty"`
	Options  []Parameter `json:"options,omitempty"`
}

// Volume contains information about a container volume
type Volume struct {
	HostPath      string          `json:"hostPath,omitempty"`
	ContainerPath string          `json:"containerPath,omitempty"`
	Mode          string          `json:"mode,omitempty"`
	External      *ExternalVolume `json:"external,omitempty"`
}

// Parameter contains information about parameter
type Parameter struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// Container contains a Mesos container description
type Container struct {
	Type           string        `json:"type,omitempty"`
	Image          string        `json:"image,omitempty"`
	ForcePullImage bool          `json:"forcePullImage, omitempty"`
	Network        string        `json:"network,omitempty"`
	NetworkName    string        `json:"networkName,omitempty"`
	NetworkInfos   []NetworkInfo `json:"networkInfos,omitempty"`
	Volumes        []Volume      `json:"volumes,omitempty"`
	Parameters     []Parameter   `json:"parameters,omitempty"`
}
