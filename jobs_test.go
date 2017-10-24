package chronos_test

import (
	"net/http"
	"net/http/httptest"

	chronos "github.com/axelspringer/go-chronos"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jobs", func() {

	var client *chronos.Client
	var httpClient *http.Client
	var mux *http.ServeMux
	var server *httptest.Server

	var testJobs []*chronos.Job
	var testJob *chronos.Job

	BeforeEach(func() {
		httpClient, mux, server = fakeTestServer()

		testJob = &chronos.Job{
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
		testJobs = make([]*chronos.Job, 0)
		testJobs = append(testJobs, testJob)

		mux.HandleFunc("/v1/scheduler/jobs", func(w http.ResponseWriter, r *http.Request) {
			Ω(r.Method).Should(Equal("GET"))
			w.Header().Set("Content-Type", "application/json")
			handleJSONEncode(w, testJobs)
		})

		client = chronos.New("https://localhost/", httpClient)
	})

	AfterEach(func() {
		defer server.Close()
	})

	Describe("List Jobs", func() {
		var jobs *[]chronos.Job

		BeforeEach(func() {
			jobs, _, _ = client.Jobs.List()
		})

		It("should contain 1 job", func() {
			Ω(*jobs).Should(HaveLen(1))
		})
	})
})
