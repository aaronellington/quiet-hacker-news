package forge

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// httpLoggerRequestID is the type used for storing the request id key
type httpLoggerRequestID string

// httpLoggerRequestIDKey is the "string" key is used for storing the request id
const httpLoggerRequestIDKey httpLoggerRequestID = "forge-request-id"

type HTTPLogger struct {
	Logger  Logger
	Handler http.Handler
}

func (httpLogger *HTTPLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	recorder := &statusRecorder{
		ResponseWriter: w,
		ResponseCode:   200,
	}
	requestWithID := httpLoggerAddRequestID(r)

	startTime := time.Now()
	httpLogger.Handler.ServeHTTP(recorder, requestWithID)
	duration := time.Since(startTime)

	httpLogger.Logger.Info("HTTP Request", map[string]interface{}{
		"requestMethod": r.Method,
		"requestPath":   r.URL.Path,
		"requestID":     HTTPLoggerGetRequestID(requestWithID),
		"requestQuery":  r.URL.Query().Encode(),
		"responseCode":  recorder.ResponseCode,
		"handleTime":    duration.Milliseconds(),
	})
}

type statusRecorder struct {
	http.ResponseWriter
	ResponseCode int
}

func (recorder *statusRecorder) WriteHeader(status int) {
	recorder.ResponseCode = status
	recorder.ResponseWriter.WriteHeader(status)
}

// HTTPLoggerGetRequestID gets the request ID off of the request if there is one
func HTTPLoggerGetRequestID(r *http.Request) string {
	if r == nil {
		return ""
	}

	requestIDRaw := r.Context().Value(httpLoggerRequestIDKey)
	requestID, ok := requestIDRaw.(string)
	if !ok {
		return ""
	}

	return requestID
}

func httpLoggerAddRequestID(r *http.Request) *http.Request {
	requestID := uuid.New()
	ctx := context.WithValue(r.Context(), httpLoggerRequestIDKey, requestID.String())
	return r.WithContext(ctx)
}
