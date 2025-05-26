package DarkThroneApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"reflect"
)

// ApiRequestConfig holds configuration for API requests, such as the base URL and logger.
type ApiRequestConfig struct {
	BaseURL string
	Logger  *slog.Logger
}

// ApiRequest represents an API request with generic request and response types.
type ApiRequest[Req any, Resp any] struct {
	Method   string
	Endpoint string
	Headers  map[string]string
	Body     Req
	Config   *ApiRequestConfig // Optional config for request-level settings
}

// isZeroValue checks if a value is the zero value for its type.
func isZeroValue[T any](v T) bool {
	return reflect.ValueOf(v).IsZero()
}

// logRequest logs the API request details if a logger is configured.
func (req *ApiRequest[Req, Resp]) logRequest(msg string) {
	if req.Config != nil && req.Config.Logger != nil {
		req.Config.Logger.Debug(msg,
			"endpoint", req.Endpoint,
			"method", req.Method,
			"body", req.Body,
			"headers", req.Headers,
		)
	}
}

// logResponse logs the API response details if a logger is configured.
func (req *ApiRequest[Req, Resp]) logResponse(msg string, resp Resp) {
	if req.Config != nil && req.Config.Logger != nil {
		req.Config.Logger.Debug(msg,
			"endpoint", req.Endpoint,
			"method", req.Method,
			"response", resp,
		)
	}
}

// logError logs an error related to the API request if a logger is configured.
func (req *ApiRequest[Req, Resp]) logError(msg string, err error) {
	if req.Config != nil && req.Config.Logger != nil {
		req.Config.Logger.Error(msg,
			"endpoint", req.Endpoint,
			"method", req.Method,
			"error", err,
		)
	}
}

// GetUrl constructs the full URL for the API request.
func (req *ApiRequest[Req, Resp]) GetUrl() string {
	if req.Config == nil || req.Config.BaseURL == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", req.Config.BaseURL, req.Endpoint)
}

// DoRequest executes the API request and returns the response or an error.
func (req ApiRequest[Req, Resp]) DoRequest() (Resp, error) {
	var zero Resp
	req.logRequest("Executing API request")
	if req.Config == nil || req.Config.BaseURL == "" {
		errorString := fmt.Errorf("ApiRequest.Config.BaseURL is required")
		req.logError("Request Error", errorString)
		return zero, errorString
	}

	var bodyReader io.Reader
	if !isZeroValue(req.Body) {
		if data, err := json.Marshal(req.Body); err != nil {
			return zero, err
		} else {
			bodyReader = bytes.NewBuffer(data)
		}
	}

	httpReq, err := http.NewRequest(req.Method, req.GetUrl(), bodyReader)
	if err != nil {
		return zero, err
	}

	if _, ok := req.Headers["Content-Type"]; !ok {
		if req.Headers == nil {
			req.Headers = make(map[string]string)

		}
		req.Headers["Content-Type"] = "application/json"
	}

	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	req.logResponse("API request response", zero)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		req.logError("Non-OK HTTP status", fmt.Errorf("Non-OK HTTP status: %s", resp.Status))
		return zero, fmt.Errorf("Non-OK HTTP status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, err
	}

	var result Resp
	resultType := reflect.TypeOf(result)
	var unmarshalTarget any = &result
	if resultType != nil {
		switch resultType.Kind() {
		case reflect.Map:
			// If Resp is a map, initialize it and unmarshal into its address
			result = reflect.MakeMap(resultType).Interface().(Resp)
			unmarshalTarget = &result
		case reflect.Ptr:
			if resultType.Elem().Kind() == reflect.Map && reflect.ValueOf(result).IsNil() {
				mapValue := reflect.MakeMap(resultType.Elem())
				ptr := reflect.New(resultType.Elem())
				ptr.Elem().Set(mapValue)
				result = ptr.Interface().(Resp)
				unmarshalTarget = result
			}
		}
	}
	err = json.Unmarshal(body, unmarshalTarget)
	return result, err
}
