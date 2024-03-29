package gozerox

import "net/http"
import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err interface{}) {
	var body Body
	if err != nil {
		xerr, ok := err.(errors.CodeMsg)

		if ok {
			body.Code = xerr.Code
			body.Msg = xerr.Msg
			httpx.OkJson(w, body)
			return
		}

		err, ok := err.(error)

		if ok {
			httpx.Error(w, err)
			return
		}
	} else {
		body.Code = 0
		body.Msg = "ok"
		body.Data = resp
		httpx.OkJson(w, body)
		return
	}
}
