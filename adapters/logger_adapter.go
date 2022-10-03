package adapters

// LoggerAdapter - Describe Logger Adapter interface that must be implemented by logger services
type LoggerAdapter interface {
	Info(interface{})
	Warn(interface{})
	Error(interface{})
	Debug(interface{})
}
