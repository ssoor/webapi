package gzip

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/ssoor/webapi"
)

const (
	BestCompression    = gzip.BestCompression
	BestSpeed          = gzip.BestSpeed
	DefaultCompression = gzip.DefaultCompression
	NoCompression      = gzip.NoCompression
)

// New returns a new Gzip middleware with DefaultCompression level.
func New() *Middleware {
	return &Middleware{
		Level: DefaultCompression,
	}
}

// Middleware represents a Gzip middleware.
type Middleware struct {
	Level                int
	gzipResponseWriter   *ResponseWriter
	originResponseWriter http.ResponseWriter
}

func (this *Middleware) OnRequest(req *http.Request, writer http.ResponseWriter) http.ResponseWriter {
	if !strings.Contains(req.Header().Get("Accept-Encoding"), "gzip") {
		return writer
	}

	gzipResponseWriter, err := NewResponseWriter(writer, this.Level)

	if err != nil {
		return writer
	}

	writer.Header().Set("Content-Encoding", "gzip")
	writer.Header().Set("Vary", "Accept-Encoding")
	writer.Header().Del("Content-Length")

	this.originResponseWriter = writer
	this.gzipResponseWriter = gzipResponseWriter

	return gzipResponseWriter
}

func (this Middleware) OnResponse(req *http.Request, writer http.ResponseWriter) http.ResponseWriter {

	this.gzipResponseWriter.Close()

	return this.originResponseWriter
}

// NewResponseWriter returns a new ResponseWriter.
func NewResponseWriter(writer http.ResponseWriter, level int) (*ResponseWriter, error) {
	gzipWriter, err := gzip.NewWriterLevel(writer, level)
	if err != nil {
		return nil, err
	}

	return &ResponseWriter{
		ResponseWriter: writer,
		gzipWriter:     gzipWriter,
	}, nil
}

// ResponseWriter represents a Gzip ResponseWriter.
type ResponseWriter struct {
	http.ResponseWriter
	gzipWriter *gzip.Writer
}

func (this ResponseWriter) Close() error {
	return this.gzipWriter.Close()
}

func (this ResponseWriter) Write(buffer []byte) (int, error) {
	return this.gzipWriter.Write(buffer)
}
