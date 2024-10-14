package response

import "github.com/gin-gonic/gin"

type Response struct {
	HttpCode int    `json:"code"`
	Message  string `json:"message"`
	Payload  any    `json:"payload"`
}

func NewResponse(httpCode int, message string, payload any) *Response {
	return &Response{
		HttpCode: httpCode,
		Message:  message,
		Payload:  payload,
	}
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.HttpCode, r)
}
