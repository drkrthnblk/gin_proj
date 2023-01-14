package requestdetails

import (
	"net/http"
)

type ApiPayload interface {
	map[string]interface{} | []map[string]interface{}
}

type RequestDetails[P ApiPayload] struct {
	Endpoint string
	Method   string
	Headers  map[string]interface{}
	Params   map[string]interface{}
	Data     P
}

func CreateRequestDetails[P ApiPayload](endpoint, method string, payload P, headers, params map[string]interface{}) (interface{}, error) {
	switch method {
	case http.MethodGet:
		d := make(map[string]interface{})
		d["key_1"] = "value_1"
		d["key_2"] = 1
		// r := RequestDetails[map[string]interface{}]{
		// 	Endpoint: endpoint,
		// 	Method:   method,
		// 	Headers:  headers,
		// 	Params:   params,
		// 	Data:  payload,
		// }

		return nil, nil
	}
	return nil, nil
}
