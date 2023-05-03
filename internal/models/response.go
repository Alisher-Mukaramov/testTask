package models

type Response struct {
	Result bool   `json:"result"`
	Info   string `json:"info"`
	Code   int    `json:"code,omitempty"`
}

func (r *Response) SetStatus(result bool, info string, code int) {
	r.Result = result
	r.Info = info
	r.Code = code
}
