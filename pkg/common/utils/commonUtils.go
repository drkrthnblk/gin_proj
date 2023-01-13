package utils

import (
	"bytes"
	"encoding/json"
	requestdetails "gin_proj/pkg/common/requestDetails"
	"gin_proj/pkg/common/restclient"
	"gin_proj/pkg/common/response"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func SendAPIRequest[P requestdetails.ApiPayload](endpoint, method *string, payload *P, headers, params *map[string]interface{}) (*[]byte, error) {
	marshaledBody, _ := json.Marshal(payload)
	data := bytes.NewBuffer(marshaledBody)
	restRequest, _ := http.NewRequest(*method, *endpoint, data)

	if headers != nil {
		for key, value := range *headers {
			restRequest.Header.Add(key, value.(string))
		}
	}

	if params != nil {
		qParams := restRequest.URL.Query()
		for key, value := range *params {
			qParams.Add(key, value)
		}
		restRequest.URL.RawQuery = qParams.Enode()
	}

	response, err := restclient.Client.Do(restRequest)
	if err != nil {
		return nil, err
	}
	defer fun(Body io.ReadCloser){
		err := Body.Close()
		if err != nil {
			return nil, err
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		var resp map[string]interface{}
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}
		jsonData, err := json.Marshal(resp)
		if err != nil {
			return nil, err
		}
		return nil, Error("something went wrong while sending digio request")
	}
	return body, nil
}
