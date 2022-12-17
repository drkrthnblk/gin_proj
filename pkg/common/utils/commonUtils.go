package utils

import (
	"io"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gin_proj/pkg/common/response"
)

func GetUUID() uuid.UUID {
	return uuid.New()
}

func GetUTCTimeFromLocal(datetimeStr string) (*time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	datetime, err := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, loc)
	if err != nil {
		errData := []gin.H{}
		errData = append(errData, gin.H{"error": err.Error()})
		return nil, response.NewGenericApiError(
			response.DatetimeOperationFAILED, response.ErrSomethingWentWrong, 500, &errData)
	}
	loc, _ = time.LoadLocation("UTC")
	utcDateTime := datetime.In(loc)
	return &utcDateTime, nil
}

func ReadRequestBody(c *gin.Context) (*[]byte, error) {
	// var requestData map[string]interface{}
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errData := []gin.H{}
		errData = append(errData, gin.H{"error": err.Error()})
		return nil, response.NewGenericApiError(
			response.InputReadError, response.ErrSomethingWentWrong, 500, &errData)
	}
	return &jsonData, nil
}