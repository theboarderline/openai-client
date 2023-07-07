package turbo_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/openai-client/turbo"
	"os"
)

var _ = Describe("Vertexai", func() {

	It("can get a chat response", func() {
		c, _ := turbo.NewClient(os.Getenv("OPENAI_API_KEY"), os.Getenv("OPENAI_ORG"))
		req := &turbo.Request{
			Model: turbo.ModelGpt35Turbo,
			Messages: []*turbo.Message{
				{
					Role:    turbo.RoleUser,
					Content: "Who is the best basketball player of all time?",
				},
			},
		}

		response, err := c.GetChat(req)
		Expect(err).NotTo(HaveOccurred())
		Expect(response.Choices[0].Message.Content).NotTo(BeEmpty())
	})

})
