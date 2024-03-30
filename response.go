package gozerox

import "net/http"
import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
)

var Success = errors.CodeMsg{Code: 0, Msg: "ok"}
var LoginStatusExpired = errors.CodeMsg{Code: 1001, Msg: "auth status expired"}

var BadRequest = errors.CodeMsg{Code: 4001, Msg: "bad request"}

var InternalServerError = errors.CodeMsg{Code: 5001, Msg: "internal server error"}

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

		case BadRequest.Error():
			body.Code = BadRequest.Code
			body.Msg = BadRequest.Msg

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
