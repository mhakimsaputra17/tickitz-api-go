package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseBody is a common structure for all responses
type ResponseBody struct {
	Status string `json:"status"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

// SuccessResponse returns a successful response with data
func SuccessResponse (c *gin.Context, statusCode int, message string, data interface{}){
	c.JSON(statusCode, ResponseBody{
		Status: "success",
		Message: message,
		Data: data,
	})
}

// ErrorResponse returns the error response
func ErrorResponse (c *gin.Context, statusCode int, message string, err interface{}){
	c.JSON(statusCode, ResponseBody{
		Status: "error",
		Message: message,
		Errors: err,
	})
}

// More specific helper functions based on HTTP status

func OkResponse(c *gin.Context, message string, data interface{}){
	SuccessResponse(c, http.StatusOK, message, data)
}

func CreatedResponse(c *gin.Context, message string, data interface{}) {
	SuccessResponse(c, http.StatusCreated, message, data)
}

func BadRequestResponse(c *gin.Context, message string, err interface{}) {
	ErrorResponse(c, http.StatusBadRequest, message, err)
}

func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

func ConflictResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusConflict, message, nil)
}

func ServerErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message, nil)
}

