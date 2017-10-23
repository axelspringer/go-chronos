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

package main

import (
	"fmt"
	"net/http"
	"os"

	chronos "github.com/axelspringer/go-chronos"
)

func main() {
	chronosURL := "https://chronos.tortuga.services/v1/scheduler/"

	httpClient := &http.Client{}

	chronosClient := chronos.New(chronosURL, httpClient)
	jobs, _, _ := chronosClient.Jobs.List()
	for _, job := range *jobs {
		fmt.Println(job.Name)
	}

	job := &chronos.Job{
		Schedule: "R10/2012-10-01T05:52:00Z/PT2S",
		Name:     "demo",
		Epsilon:  "PT15M",
		Command:  "echo 'Test'",
		Mem:      64,
		CPUs:     0.1,
		Container: &chronos.Container{
			Image:       "alpine:3.6",
			Network:     "USER",
			NetworkName: "mesos",
		},
	}

	if ok, _, _ := chronosClient.Job.New(job); ok {
		fmt.Printf("Successfully scheduled: %v", job.Name)
	} else {
		os.Exit(-1)
	}

	jobs, _, _ = chronosClient.Jobs.Search(&chronos.JobsSearchParams{Name: "demo", Command: "echo"})
	for _, job := range *jobs {
		fmt.Println(job.Name)
	}

	if ok, _, _ := chronosClient.Job.Start(job.Name, &chronos.JobStartParams{}); ok {
		fmt.Printf("Successfully startet: %v", job.Name)
	}

	if ok, _, _ := chronosClient.Job.MarkSuccess(job.Name); ok {
		fmt.Printf("Successfully job as success %v", ok)
	}
}
