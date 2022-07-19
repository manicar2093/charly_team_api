package services_test

import (
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/manicar2093/charly_team_api/internal/services"
)

var _ = Describe("Passwords", func() {
	var (
		passService    *services.BcryptImpl
		passwordToHash string
	)

	BeforeEach(func() {
		passService = &services.BcryptImpl{}
		passwordToHash = faker.Password()
	})

	Describe("Digest", func() {
		It("generate a password hash from a string", func() {
			got, err := passService.Digest(passwordToHash)

			Expect(err).ToNot(HaveOccurred())
			Expect(got).ToNot(BeEmpty())
		})
	})

	Describe("Compare", func() {
		It("sends a nil error when pass is ok", func() {
			var (
				plainPassword  = "bCLxihZAbTLcaGEGeRsoBnuwWxNJqQTSpeDMxswQKJoJdbyWxD"
				hashedPassword = "$2a$10$oPHltqIW/TtZ2GlMz7RfM.sglZIgPoaglBQa3P8mJMtkjcGG4aVDm"
			)

			err := passService.Compare(hashedPassword, plainPassword)

			Expect(err).ToNot(HaveOccurred())
		})
	})

})
