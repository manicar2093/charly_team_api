package tokenclaimsgenerator

type TokenClaimsGeneratorRequest struct {
	UserUUID string `validate:"required"`
}

type TokenClaimsGeneratorResponse struct {
	Claims map[string]string
}
