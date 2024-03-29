package gozerox

import "net/http"
import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
)

var Success = errors.CodeMsg{Code: 0, Msg: "ok"}
var InternalServerError = errors.CodeMsg{Code: 500, Msg: "internal server error"}
var LoginStatusExpired = errors.CodeMsg{Code: 2001, Msg: "auth status expired"}

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		switch err.Error() {
		case LoginStatusExpired.Error():
			body.Code = LoginStatusExpired.Code
			body.Msg = LoginStatusExpired.Msg
		default:
			body.Code = InternalServerError.Code
			body.Msg = InternalServerError.Msg
		}
	} else {
		body.Code = Success.Code
		body.Msg = Success.Msg
		body.Data = resp
	}

	httpx.OkJson(w, body)
	return

}
