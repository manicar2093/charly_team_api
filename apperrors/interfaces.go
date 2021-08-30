package apperrors

type HandableErrors interface {
	Error() string
	StatusCode() int
}
