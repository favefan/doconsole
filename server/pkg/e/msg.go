package e

// MsgFlags sets message for given flag
var MsgFlags = map[int]string{
	Success:                           "处理成功",
	Error:                             "处理失败",
	InvalidParams:                     "请求参数错误",
	ErrorAuth:                         "Token错误",
	ErrorDockerDaemonConnectionFailed: "连接Docker失败，请检查Docker运行状态",
	ErrorAuthCheckTokenFail:           "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:        "Token已超时",
	ErrorAuthToken:                    "Token生成失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
