package response

type ErrorCode string

const (
	CodeUnnown ErrorCode = "CODE_UNKNOWN"
	CodeBadRequest ErrorCode = "CODE_BAD_REQUEST"
)