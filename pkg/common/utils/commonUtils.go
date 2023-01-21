package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	requestdetails "gin_proj/pkg/common/requestDetails"
	"gin_proj/pkg/common/response"
	"gin_proj/pkg/common/restclient"
	"io"
	"log"
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

func SendAPIRequest[P requestdetails.ApiPayload](endpoint, method *string, payload *P, headers, params *map[string]string) (*[]byte, error) {
	marshaledBody, _ := json.Marshal(payload)
	data := bytes.NewBuffer(marshaledBody)
	restRequest, _ := http.NewRequest(*method, *endpoint, data)

	if headers != nil {
		for key, value := range *headers {
			restRequest.Header.Add(key, value)
		}
	}

	if params != nil {
		qParams := restRequest.URL.Query()
		for key, value := range *params {
			qParams.Add(key, value)
		}
		restRequest.URL.RawQuery = qParams.Encode()
	}

	response, err := restclient.Client.Do(restRequest)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
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
		return nil, fmt.Errorf("something went wrong while sending digio request %s", jsonData)
	}
	return &body, nil
}

type JSON json.RawMessage

// // Scan scan value into Jsonb, implements sql.Scanner interface
// func (j *JSON) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
// 	}

// 	result := json.RawMessage{}
// 	err := json.Unmarshal(bytes, &result)
// 	*j = JSON(result)
// 	return err
// }

// // Value return json value, implement driver.Valuer interface
// func (j JSON) Value() (driver.Value, error) {
// 	if len(j) == 0 {
// 		return nil, nil
// 	}
// 	return json.RawMessage(j).MarshalJSON()
// }

func ParseJSONToModel(src, dest interface{}) error {
	var data []byte

	if b, ok := src.([]byte); ok {
		data = b
	} else if s, ok := src.(string); ok {
		data = []byte(s)
	}
	return json.Unmarshal(data, dest)
}

// func ParseModelToJSON(j JSON)(driver.Value, error) {
// 	if len(j) == 0 {
// 		return nil, nil
// 	}
// 	return json.Marshal(j)
// }
