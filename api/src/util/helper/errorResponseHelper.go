package helper

import (
	"api/src/util/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message      string `json:"message"`
	ErrorMessage string `json:"errorMessage"`
}

func ErrorResponseMethod(ctx *gin.Context, err error) {
	errorResponse := ErrorResponse{}
	errorResponse.Message = findErrorType(err)
	errorResponse.ErrorMessage = err.Error()

	ctx.JSON(http.StatusInternalServerError, errorResponse)
}

func findErrorType(err error) string {
	var errorType constant.ErrorType

	switch err.Error() {
	case string(constant.MaxAttemptGenerateCode):
		errorType = constant.MaxAttempt
	default:
		errorType = constant.DateError
	}

	return string(errorType)
}
