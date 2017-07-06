package metrics

import (
	"net/http"
)

// LoggedResponseWriter define http.ResponseWriter with Status
type LoggedResponseWriter interface {
	http.ResponseWriter
	Status() int
}

// loggedResponseWriter is a http.ResponseWriter with Status.
type loggedResponseWriter struct {
	http.ResponseWriter
	status int
}

// loggedResponseWriterHijacker is a http.ResponseWriter and http.Hijacker
// with Status.
type loggedResponseWriterHijacker struct {
	LoggedResponseWriter
	http.Hijacker
}

// NewLoggedResponseWriter wraps the provided http.ResponseWriter with Status
// preserving the support to hijack the connection if supported by the provided
// http.ResponseWriter.
func NewLoggedResponseWriter(w http.ResponseWriter) LoggedResponseWriter {
	lw := &loggedResponseWriter{ResponseWriter: w}
	if hj, ok := w.(http.Hijacker); ok {
		return &loggedResponseWriterHijacker{LoggedResponseWriter: lw, Hijacker: hj}
	}

	return lw
}

// WriteHeader sends an HTTP response header with status code.
func (w *loggedResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// Status returns the written HTTP respons header status code or http.StatusOK
// if none was written.
func (w *loggedResponseWriter) Status() int {
	status := w.status
	if status == 0 {
		status = http.StatusOK
	}
	return status
}
