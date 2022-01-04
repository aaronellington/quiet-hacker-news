package forge

import "net/http"

// Log Levels
const (
	LogLevelError   = "error"
	LogLevelWarning = "warning"
	LogLevelInfo    = "info"
	LogLevelDebug   = "debug"
)

// Logger Service
type Logger interface {
	// Error logs an error message
	Error(r *http.Request, message string, context map[string]interface{}) error
	// Info logs an info message
	Info(r *http.Request, message string, context map[string]interface{}) error
	// Warning logs a warning message
	Warning(r *http.Request, message string, context map[string]interface{}) error
	// Debug logs a debug message
	Debug(r *http.Request, message string, context map[string]interface{}) error
}
