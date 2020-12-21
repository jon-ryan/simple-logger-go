package simplelogger

import "log"

// NewSimpleLogger returns a new simple logger instance
func NewSimpleLogger(loggingLevel LoggingLevel) Simplelogger {
	return Simplelogger{loggingLevel: loggingLevel}
}

// LoggingLevel defines the enum type to set the logging level
type LoggingLevel int

// Enum for the different log levels
const (
	None LoggingLevel = iota
	Error
	Warning
	Info
)

// Simplelogger struct
type Simplelogger struct {
	loggingLevel LoggingLevel
}

// SetLoggingLevel sets a new logging level
func (s *Simplelogger) SetLoggingLevel(level LoggingLevel) {
	s.loggingLevel = level
}

// Info logs info information
func (s *Simplelogger) Info(msg string) {
	if s.loggingLevel == Info {
		log.Println("[INFO]", msg)
	}
}

// Warning logs warning information
func (s *Simplelogger) Warning(msg string) {
	if s.loggingLevel >= Warning {
		log.Println("[WARNING]", msg)
	}
}

// Error logs error information
func (s *Simplelogger) Error(msg string) {
	if s.loggingLevel >= Error {
		log.Println("[ERROR]", msg)
	}
}
