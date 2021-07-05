package testinglogur

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type T interface {
	Log(...interface{})
}

// This is the logger interface suggested by https://github.com/logur/logur
type MyLogger interface {
	Trace(msg string, fields ...map[string]interface{})
	Debug(msg string, fields ...map[string]interface{})
	Info(msg string, fields ...map[string]interface{})
	Warn(msg string, fields ...map[string]interface{})
	Error(msg string, fields ...map[string]interface{})
}

type logger func(...interface{})

// Get returns a logger that matches the interface proposed by log
func Get(t T) MyLogger {
	return logger(t.Log)
}

func (l logger) Trace(msg string, fields ...map[string]interface{}) {
	if len(fields) == 0 {
		l(msg)
		return
	}
	out := make([]string, 1, len(fields)*3+1)
	out[0] = msg
	for _, field := range fields {
		keys := make([]string, 0, len(field))
		for k := range field {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			enc, err := json.Marshal(field[k])
			if err == nil {
				out = append(out, k+"="+string(enc))
			} else {
				out = append(out, k+"="+fmt.Sprint(field[k]))
			}
		}
	}
	l(strings.Join(out, ", "))
}

func (l logger) Debug(msg string, fields ...map[string]interface{}) { l.Trace(msg, fields...) }
func (l logger) Info(msg string, fields ...map[string]interface{})  { l.Trace(msg, fields...) }
func (l logger) Warn(msg string, fields ...map[string]interface{})  { l.Trace(msg, fields...) }
func (l logger) Error(msg string, fields ...map[string]interface{}) { l.Trace(msg, fields...) }
