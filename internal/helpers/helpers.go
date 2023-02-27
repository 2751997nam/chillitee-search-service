package helpers

import (
	"fmt"
	"logger-service/internal/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, result any, status int) {
	c.IndentedJSON(status, types.Response{
		Status: "successful",
		Result: result,
	})
}

func AnyToString(value any) string {
	if value != nil {
		return fmt.Sprint(value)
	}

	return ""
}

func ResponseSuccessWithMessage(c *gin.Context, result any, message string) {
	c.IndentedJSON(http.StatusOK, types.Response{
		Status:  "successful",
		Result:  result,
		Message: message,
	})
}

func ResponseWithMeta(c *gin.Context, result any, meta types.Meta) {
	c.IndentedJSON(http.StatusOK, types.Response{
		Status: "successful",
		Result: result,
		Meta:   meta,
	})
}

func ResponseFail(c *gin.Context, message string, status int) {
	c.IndentedJSON(status, types.Response{
		Status:  "fail",
		Message: message,
	})
}

func AnyToInt(value any) int {
	result, _ := strconv.Atoi(fmt.Sprint(value))
	return result
}

func AnyToUint(value any) uint {
	result, _ := strconv.ParseInt(fmt.Sprint(value), 10, 64)
	return uint(result)
}

func AnyToFloat(value any) float32 {
	result, _ := strconv.ParseFloat(fmt.Sprint(value), 64)
	return float32(result)
}
