package metrics

import (
	"net/http"
	"strconv"
)

type LoggedResponseWriter struct {
	http.ResponseWriter
	http.Hijacker
	status int
}

func (w *LoggedResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *LoggedResponseWriter) Status() string {
	status := w.status
	if status == 0 {
		status = http.StatusOK
	}
	return strconv.Itoa(status)
}
