package kit

import (
	"log"

	"github.com/go-kit/kit/log"
	"github.com/grpc-ecosystem/go-grpc-middleware.v2/interceptors/logging"
)

type Logger struct {
	*log.Logger
}

func New(logger *log.Logger) *Logger {
	return &Logger{logger}
}

func (l *Logger) Log(lvl logging.Level, msg string) {

}

func (l *Logger) With(fields ...string) logging.Logger {
	return New(l.With(fields...))
}
