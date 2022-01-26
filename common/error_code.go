package common

const (
	// ErrorPanic panic
	ErrorPanic = 1 + iota
	// ErrorInvalidParams 参数不合法
	ErrorInvalidParams
	// ErrorMarshal 序列化错误
	ErrorMarshal
	// ErrorUnmarshal 反序列化错误
	ErrorUnmarshal
	// ErrorInner 内部错误
	ErrorInner
)

// BaseResponse 基础返回
type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
