package response

type ErrorCode string

const (
	CodePanic ErrorCode = "CODE_PANIC"
	CodeUnknown ErrorCode = "CODE_UNKNOWN"
	CodeBadRequest ErrorCode = "CODE_BAD_REQUEST"
)