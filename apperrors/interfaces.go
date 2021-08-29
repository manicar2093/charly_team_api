package apperrors

type HandableErrors interface {
	StatusCode() int
}
