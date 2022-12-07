package service

import "log"

func NewLoggerService(logInfo *log.Logger, logErr *log.Logger) LoggerService {
	return &loggerService{
		logInfo: logInfo,
		logErr:  logErr,
	}
}

type loggerService struct {
	logInfo *log.Logger
	logErr  *log.Logger
}

// Error implements LoggerService
func (s *loggerService) Error(format string, a ...any) {
	s.logErr.Printf(format, a...)
}

// Info implements LoggerService
func (s *loggerService) Info(format string, a ...any) {
	s.logInfo.Printf(format, a...)
}
