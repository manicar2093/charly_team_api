package token_test

import (
	"github.com/manicar2093/charly_team_api/internal/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateToken", func() {

	var (
		service *token.Creator
	)

	BeforeEach(func() {
		service = token.NewCreator()
	})

	It("creates a token with required claims", func() {
		var claims = map[string]interface{}{
			"user":        "user",
			"permissions": []string{"perm1", "perm2", "perm3"},
		}

		token, err := service.Gen(claims)

		Expect(err).ToNot(HaveOccurred())
		Expect(token).ToNot(BeEmpty())
	})

})
