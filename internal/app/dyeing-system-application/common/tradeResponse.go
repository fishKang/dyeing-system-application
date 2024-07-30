package common

import "wk.com/dyeing-system-application/internal/app/dyeing-system-application/util"

type TradeResponse struct {
	Returncode string `json:"returncode"` // 返回码
	Returnmsg  string `json:"returnmsg"`  // 返回信息
	Returninfo string `json:"returninfo"` // 异常信息
	Data       any    `json:"data"`       // 返回正常信息
}

func CreateSuccessResponse(data any) *TradeResponse {
	var tradeResponse TradeResponse
	tradeResponse.Returncode = util.OK
	tradeResponse.Returnmsg = "处理成功"
	tradeResponse.Data = data
	return &tradeResponse
}

func CreateFailResponse(code string, msg string, info string) *TradeResponse {
	var tradeResponse TradeResponse
	tradeResponse.Returncode = code
	tradeResponse.Returnmsg = msg
	tradeResponse.Returninfo = info
	return &tradeResponse
}
