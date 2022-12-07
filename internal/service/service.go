package service

type LoggerService interface {
	Info(format string, a ...any)
	Error(format string, a ...any)
}
