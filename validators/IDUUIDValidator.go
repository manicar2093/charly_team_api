package validators

type IDUUIDValidable interface {
	GetID() int32
	GetUUID() string
}

// isUpdateRequestValid check if ID or UUID is in biotest entity request
func IsUpdateRequestValid(toValidate IDUUIDValidable) bool {

	if toValidate.GetID() > 0 {
		return true
	}

	if toValidate.GetUUID() != "" {
		return true
	}

	return false
}
