package logger

type LogLevel string
type LogOutput string

const (
	LogLevelDebug    LogLevel  = "DEBUG"
	LogLevelInfo     LogLevel  = "INFO"
	LogLevelWarn     LogLevel  = "WARN"
	LogLevelError    LogLevel  = "ERROR"
	LogOutputConsole LogOutput = "CONSOLE"
	LogOutputFile    LogOutput = "FILE"
)
