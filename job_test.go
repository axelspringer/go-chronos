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

	var testJob string

	BeforeEach(func() {
		httpClient, mux, server = fakeTestServer()
		testJob = "demo"

		client = chronos.New("https://localhost/", httpClient)
	})

	AfterEach(func() {
		defer server.Close()
	})

	Describe("Delete Job", func() {
		var ok bool

		BeforeEach(func() {
			mux.HandleFunc("/v1/scheduler/job/demo", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("DELETE"))
				w.WriteHeader(204)
			})

			ok, _, _ = client.Job.Delete(testJob)
		})

		It("should be deleted job", func() {
			Ω(ok).Should(Equal(true))
		})
	})

	Describe("Trigger a Job", func() {
		var ok bool

		BeforeEach(func() {
			mux.HandleFunc("/v1/scheduler/job/demo", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("PUT"))
				w.WriteHeader(204)
			})

			ok, _, _ = client.Job.Start(testJob, &chronos.JobStartParams{})
		})

		It("should have started the job", func() {
			Ω(ok).Should(Equal(true))
		})
	})
})
