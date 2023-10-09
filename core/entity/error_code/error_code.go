package error_code

type ErrorCode int

const (
	Success ErrorCode = iota
	InvalidRequest
	Duplicate
	InternalError
)

type ErrorMessage int

const (
	SuccessErrMsg ErrorMessage = iota
	InternalErrMsg
)

func (e ErrorMessage) String() string {
	switch e {
	case SuccessErrMsg:
		return "success"
	case InternalErrMsg:
		return "internal error"
	}
	return "unknown"
}
