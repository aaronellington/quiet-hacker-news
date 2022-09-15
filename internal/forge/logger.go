package forge

import (
	"encoding/json"
	"log"
)

// Log is foobar
type Log struct {
	Severity string                 `json:"severity"`
	Message  string                 `json:"message"`
	Context  map[string]interface{} `json:"context"`
}

// Logger is foobar
type Logger interface {
	Emergency(message string, context map[string]interface{})
	Alert(message string, context map[string]interface{})
	Critical(message string, context map[string]interface{})
	Error(message string, context map[string]interface{})
	Warning(message string, context map[string]interface{})
	Notice(message string, context map[string]interface{})
	Info(message string, context map[string]interface{})
	Debug(message string, context map[string]interface{})
	StandardLogger() *log.Logger
}

// LoggerJSON is foobar
type LoggerJSON struct {
	Encoder *json.Encoder
}

// Emergency is foobar
func (logger *LoggerJSON) Emergency(message string, context map[string]interface{}) {
	logger.log("EMERGENCY", message, context)
}

// Alert is foobar
func (logger *LoggerJSON) Alert(message string, context map[string]interface{}) {
	logger.log("ALERT", message, context)
}

// Critical is foobar
func (logger *LoggerJSON) Critical(message string, context map[string]interface{}) {
	logger.log("CRITICAL", message, context)
}

// Error is foobar
func (logger *LoggerJSON) Error(message string, context map[string]interface{}) {
	logger.log("ERROR", message, context)
}

// Warning is foobar
func (logger *LoggerJSON) Warning(message string, context map[string]interface{}) {
	logger.log("WARNING", message, context)
}

// Notice is foobar
func (logger *LoggerJSON) Notice(message string, context map[string]interface{}) {
	logger.log("NOTICE", message, context)
}

// Info is foobar
func (logger *LoggerJSON) Info(message string, context map[string]interface{}) {
	logger.log("INFO", message, context)
}

// Debug is foobar
func (logger *LoggerJSON) Debug(message string, context map[string]interface{}) {
	logger.log("DEBUG", message, context)
}

func (logger *LoggerJSON) log(severity string, message string, context map[string]interface{}) {
	if context == nil {
		context = map[string]interface{}{}
	}

	logger.Encoder.Encode(Log{
		Severity: severity,
		Message:  message,
		Context:  context,
	})
}

// Write is foobar
func (logger *LoggerJSON) Write(b []byte) (int, error) {
	logger.Alert("Standard Library Log", map[string]interface{}{
		"log": string(b),
	})

	return 0, nil
}

// StandardLogger is foobar
func (logger *LoggerJSON) StandardLogger() *log.Logger {
	return log.New(logger, "", 0)
}
