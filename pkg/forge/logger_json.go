package forge

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// LoggerJSON is a JSON logger
type LoggerJSON struct {
	Writer io.Writer
	e      *json.Encoder
}

type loggerJSONLog struct {
	Date      time.Time              `json:"date"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"requestID"`
	Context   map[string]interface{} `json:"context"`
}

// Error logs an error message
func (logger *LoggerJSON) Error(r *http.Request, message string, context map[string]interface{}) error {
	return logger.log(r, LogLevelError, message, context)
}

// Warning logs a warning message
func (logger *LoggerJSON) Warning(r *http.Request, message string, context map[string]interface{}) error {
	return logger.log(r, LogLevelWarning, message, context)
}

// Info logs an info message
func (logger *LoggerJSON) Info(r *http.Request, message string, context map[string]interface{}) error {
	return logger.log(r, LogLevelInfo, message, context)
}

// Debug logs a debug message
func (logger *LoggerJSON) Debug(r *http.Request, message string, context map[string]interface{}) error {
	return logger.log(r, LogLevelDebug, message, context)
}

func (logger *LoggerJSON) log(r *http.Request, level string, message string, context map[string]interface{}) error {
	// Panic if there is no writer
	if logger.Writer == nil {
		panic("no writer has been set")
	}

	if logger.e == nil {
		logger.e = json.NewEncoder(logger.Writer)
	}

	if context == nil {
		context = make(map[string]interface{})
	}

	logMessage := loggerJSONLog{
		Date:      time.Now(),
		Level:     level,
		Message:   message,
		Context:   context,
		RequestID: LoggerGetRequestID(r),
	}

	if err := logger.e.Encode(logMessage); err != nil {
		return err
	}

	return nil
}
