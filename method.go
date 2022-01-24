package izlog

import "fmt"

func (l *logger) Debug(msg string) {
	l.write("[DEBUG]", msg)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.write("[DEBUG]", msg)
}

func (l *logger) Info(msg string) {
	l.write("[INFO]", msg)
}

func (l *logger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.write("[INFO]", msg)
}

func (l *logger) Error(msg string) {
	l.write("[ERROR]", msg)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.write("[ERROR]", msg)
}
