package util

const (
	OK      = "0000" // 正常，处理成功
	Error   = "9999" // 异常，处理失败
	Unknown = "8888" // 未知，处理未知

	RecordNotFound        = "1001" // 未找到记录
	ParamConvertFailed    = "1002" // 参数转换失败
	StatusNoContent       = 204    // RFC 9110, 15.3.5
	StatusResetContent    = 205    // RFC 9110, 15.3.6
	StatusPartialContent  = 206    // RFC 9110, 15.3.7
	StatusMultiStatus     = 207    // RFC 4918, 11.1
	StatusAlreadyReported = 208    // RFC 5842, 7.1
	StatusIMUsed          = 226    // RFC 3229, 10.4.1
)
