package api_test

import (
	"encoding/json"

	. "github.com/UESTC-ACM/acm-training/internal/api"
	"github.com/gogf/gf/frame/g"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response JSON encoding and decoding", func() {
	Context("normal response without error and only have simple JSON as data value", func() {
		response := Response{
			Data: g.Map{
				"status": "ok",
			},
		}
		jsonBytes := []byte(`{"data":{"status":"ok"}}`)

		It("Should be decoded without error fields", func() {
			encoded, err := json.Marshal(response)
			Ω(encoded).Should(Equal(jsonBytes))
			Ω(err).Should(BeNil())
		})

		It("Should be encoded successfully", func() {
			decoded := Response{}
			err := json.Unmarshal(jsonBytes, &decoded)
			Ω(decoded).Should(Equal(response))
			Ω(err).Should(BeNil())
		})
	})

	Context("normal response with error and no data field filled", func() {
		response := Response{
			ErrorCode:    "SYS_ERROR_CODE",
			ErrorMessage: "message of the error",
		}
		jsonBytes := []byte(`{"error_code":"SYS_ERROR_CODE","error_message":"message of the error"}`)

		It("Should be decoded without data fields", func() {
			encoded, err := json.Marshal(response)
			Ω(encoded).Should(Equal(jsonBytes))
			Ω(err).Should(BeNil())
		})

		It("Should be encoded successfully", func() {
			decoded := Response{}
			err := json.Unmarshal(jsonBytes, &decoded)
			Ω(decoded).Should(Equal(response))
			Ω(err).Should(BeNil())
		})
	})
})
