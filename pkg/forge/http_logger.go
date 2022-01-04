package forge

import (
	"context"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// HTTPLogger logs all request before passing off to the Handler
type HTTPLogger struct {
	Handler http.Handler
	Log     Logger
}

// ServerHTTP satisfies the http.Handler interface
func (logger *HTTPLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if logger.Log == nil {
		logger.Log = &LoggerJSON{
			Writer: os.Stdout,
		}
	}

	recorder := &statusRecorder{
		ResponseWriter: w,
		Status:         200,
	}

	if logger.Handler != nil {
		r = LoggerAddRequestID(r)
		logger.Handler.ServeHTTP(recorder, r)
	}

	logger.Log.Info(r, "HTTP Request", map[string]interface{}{
		"status":     recorder.Status,
		"remoteAddr": r.RemoteAddr,
		"method":     r.Method,
		"requestURI": r.RequestURI,
	})
}

type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

// loggerRequestID is the type used for storing the request id key
type loggerRequestID string

// loggerRequestIDKey is the "string" key is used for storing the request id
const loggerRequestIDKey loggerRequestID = "request-id"

// LoggerAddRequestID adds the request ID to the request
func LoggerAddRequestID(r *http.Request) *http.Request {
	ctx := r.Context()
	requestID := uuid.New()
	ctx = context.WithValue(ctx, loggerRequestIDKey, requestID.String())
	return r.WithContext(ctx)
}

// LoggerGetRequestID gets the request ID off of the request if there is one
func LoggerGetRequestID(r *http.Request) string {
	if r == nil {
		return ""
	}

	ctx := r.Context()
	requestIDRaw := ctx.Value(loggerRequestIDKey)
	requestID, ok := requestIDRaw.(string)
	if !ok {
		return ""
	}

	return requestID
}
