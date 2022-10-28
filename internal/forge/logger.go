package forge

import (
	"encoding/json"
	"log"
)

type Log struct {
	Severity string                 `json:"severity"`
	Message  string                 `json:"message"`
	Context  map[string]interface{} `json:"context"`
}

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

type LoggerJSON struct {
	Encoder *json.Encoder
}

func (logger *LoggerJSON) Emergency(message string, context map[string]interface{}) {
	logger.log("EMERGENCY", message, context)
}

func (logger *LoggerJSON) Alert(message string, context map[string]interface{}) {
	logger.log("ALERT", message, context)
}

func (logger *LoggerJSON) Critical(message string, context map[string]interface{}) {
	logger.log("CRITICAL", message, context)
}

func (logger *LoggerJSON) Error(message string, context map[string]interface{}) {
	logger.log("ERROR", message, context)
}

func (logger *LoggerJSON) Warning(message string, context map[string]interface{}) {
	logger.log("WARNING", message, context)
}

func (logger *LoggerJSON) Notice(message string, context map[string]interface{}) {
	logger.log("NOTICE", message, context)
}

func (logger *LoggerJSON) Info(message string, context map[string]interface{}) {
	logger.log("INFO", message, context)
}

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

func (logger *LoggerJSON) Write(b []byte) (int, error) {
	logger.Alert("Standard Library Log", map[string]interface{}{
		"log": string(b),
	})

	return 0, nil
}

func (logger *LoggerJSON) StandardLogger() *log.Logger {
	return log.New(logger, "", 0)
}
