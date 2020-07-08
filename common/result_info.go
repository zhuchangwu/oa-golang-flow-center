package common

type ResultInfo struct {
	Code   int64 `json:"code"`
	ResMsg string
	ErrMsg string
	Obj    interface{}  `json:"objData"`
}

func CreateSuccessfulResultInfo(msg string, obj ...interface{}) *ResultInfo {
	return &ResultInfo{
		Code:   200,
		ResMsg: msg,
		ErrMsg: "",
		Obj:    obj,
	}
}

func CreateErrorResultInfo(msg string) *ResultInfo {
	return &ResultInfo{
		// 这里err同样设置成200，目的是通过前端的统一拦截，进入到具体到组件中具体处理这个错误信息
		Code:   200,
		ResMsg: "XXX",
		ErrMsg: msg,
	}
}
