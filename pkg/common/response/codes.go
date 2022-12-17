package response

type ErrorCode string

const (
	CodePanic ErrorCode = "CODE_PANIC"
	CodeUnknown ErrorCode = "CODE_UNKNOWN"
	CodeBadRequest ErrorCode = "CODE_BAD_REQUEST"
	InputReadError ErrorCode = "INPUT_READ_ERROR"
	DatetimeOperationFAILED ErrorCode = "DATETIME_OPERATION_FAILED"
)