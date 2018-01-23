package chronos_test

import (
	"encoding/json"
	"io/ioutil"
	"path"

	chronos "github.com/axelspringer/go-chronos"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const FixtureBasePath = "_fixtures"

func getFixtureAsString(f string) string {
	p := path.Join(FixtureBasePath, f)
	b, err := ioutil.ReadFile(p)

	Ω(err).To(BeNil())
	return string(b)
}

var _ = Describe("Types", func() {

	Describe("Job type serialization/deserialization", func() {

		// Based on chronos rest api job example
		It("should pass fixture 1 around the world test", func() {
			fixData := getFixtureAsString("types_fixture1.json")
			Ω(fixData).ShouldNot(BeEmpty())

			job := &chronos.Job{}
			Ω(json.Unmarshal([]byte(fixData), job)).Should(BeNil())

			b, err := json.Marshal(job)
			Ω(err).Should(BeNil())
			Ω(string(b)).Should(MatchJSON([]byte(fixData)))
		})

		// Based on chronos rest api job example
		It("should pass fixture 2 around the world test", func() {
			fixData := getFixtureAsString("types_fixture2.json")
			Ω(fixData).ShouldNot(BeEmpty())

			job := &chronos.Job{}
			Ω(json.Unmarshal([]byte(fixData), job)).Should(BeNil())

			b, err := json.Marshal(job)
			Ω(err).Should(BeNil())
			Ω(string(b)).Should(MatchJSON([]byte(fixData)))
		})

		// This fixture should contain all parameter processed by chronos JobDeserializer
		It("should pass fixture 3 around the world test", func() {
			fixData := getFixtureAsString("types_fixture3.json")
			Ω(fixData).ShouldNot(BeEmpty())

			job := &chronos.Job{}
			Ω(json.Unmarshal([]byte(fixData), job)).Should(BeNil())

			b, err := json.Marshal(job)
			Ω(err).Should(BeNil())
			Ω(string(b)).Should(MatchJSON([]byte(fixData)))
		})

	})
})
