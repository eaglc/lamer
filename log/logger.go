package log

import (
    "fmt"
    log2 "log"
    "os"
)

var callpath = 3
var l = NewLogger()

type Logger interface {
    Debug(v ...interface{})
    Debugf(format string, v ...interface{})

    Info(v ...interface{})
    Infof(format string, v ...interface{})

    Warn(v ...interface{})
    Warnf(format string, v ...interface{})

    Error(v ...interface{})
    Errorf(format string, v ...interface{})

    Fatal(v ...interface{})
    Fatalf(format string, v ...interface{})

    Panic(v ...interface{})
    Panicf(format string, v ...interface{})
}

func SetLogger(ll Logger)  {
    l = ll
}

func Debug(v ...interface{}) {
    l.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
    l.Debugf(format, v...)
}

func Info(v ...interface{}) {
    l.Info(v...)
}

func Infof(format string, v ...interface{}) {
    l.Infof(format, v...)
}

func Warn(v ...interface{}) {
    l.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
    l.Warnf(format, v...)
}

func Error(v ...interface{}) {
    l.Error(v...)
}

func Errorf(format string, v ...interface{}) {
    l.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
    l.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
    l.Fatalf(format, v...)
}

func Panic(v ...interface{}) {
    l.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
    l.Panicf(format, v...)
}

type logger struct {
    *log2.Logger
}

func (l *logger) Debug(v ...interface{}) {
    _ = l.Output(callpath, tag("d", fmt.Sprint(v...)))
}

func (l *logger) Debugf(format string, v ...interface{}) {
    _ = l.Output(callpath, tag("d", fmt.Sprintf(format, v...)))
}

func (l *logger) Info(v ...interface{}) {
    _ = l.Output(callpath, tag("i", fmt.Sprint(v...)))
}

func (l *logger) Infof(format string, v ...interface{}) {
    _ = l.Output(callpath, tag("i", fmt.Sprintf(format, v...)))
}

func (l *logger) Warn(v ...interface{}) {
    _ = l.Output(callpath, tag("w", fmt.Sprint(v...)))
}

func (l *logger) Warnf(format string, v ...interface{}) {
    _ = l.Output(callpath, tag("w", fmt.Sprintf(format, v...)))
}

func (l *logger) Error(v ...interface{}) {
    _ = l.Output(callpath, tag("e", fmt.Sprint(v...)))
}

func (l *logger) Errorf(format string, v ...interface{}) {
    _ = l.Output(callpath, tag("e", fmt.Sprintf(format, v...)))
}

func (l *logger) Fatal(v ...interface{}) {
    _ = l.Output(callpath, tag("f", fmt.Sprint(v...)))
}

func (l *logger) Fatalf(format string, v ...interface{}) {
    _ = l.Output(callpath, tag("f", fmt.Sprintf(format, v...)))
}

func (l *logger) Panic(v ...interface{}) {
    _ = l.Output(callpath, tag("p", fmt.Sprint(v...)))
}

func (l *logger) Panicf(format string, v ...interface{}) {
    _ = l.Output(callpath, tag("p", fmt.Sprintf(format, v...)))
}

func tag(lvl string, msg string) string {
    return fmt.Sprintf("%s: %s", lvl, msg)
}

func NewLogger() Logger {
    return &logger{
        log2.New(os.Stdout, "", log2.LstdFlags|log2.Lshortfile),
    }
}