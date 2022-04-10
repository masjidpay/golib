package common

import (
	"errors"
	http2 "github.com/masjidpay/utilities/common/http"
	"net/http"
)

var ErrInvalidRequest = errors.New("invalid request")
var ErrUnprocessableEntity = errors.New("unprocessable data")
var ErrNotFoundData = errors.New("data not found")
var ErrUserAlreadyExist = errors.New("user already exist")
var ErrUserNotFound = errors.New("user not found")
var ErrPasswordNotMatch = errors.New("password doesn't match")

func InjectErrors(handlerCtx *http2.HandlerContext) {
	handlerCtx.AddError(ErrNotFoundData, setErrResp(ErrNotFoundData.Error(), http.StatusNotFound))
	handlerCtx.AddError(ErrUnprocessableEntity, setErrResp(ErrUnprocessableEntity.Error(), http.StatusUnprocessableEntity))
	handlerCtx.AddError(ErrInvalidRequest, setErrResp(ErrInvalidRequest.Error(), http.StatusBadRequest))
	handlerCtx.AddError(ErrUserAlreadyExist, setErrResp(ErrUserAlreadyExist.Error(), http.StatusConflict))
	handlerCtx.AddError(ErrUserNotFound, setErrResp(ErrUserNotFound.Error(), http.StatusNotFound))
	handlerCtx.AddError(ErrPasswordNotMatch, setErrResp(ErrPasswordNotMatch.Error(), http.StatusConflict))

}

func setErrResp(message string, statusCode int) *http2.ErrorResponse {
	return &http2.ErrorResponse{
		Response: http2.Response{
			ResponseDesc: message,
		},
		HttpStatus: statusCode,
	}
}
