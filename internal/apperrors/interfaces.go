package apperrors

type HandableErrors interface {
	error
	StatusCode() int
}
