package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/ssoor/webapi"
)

const (
	_DefaultProject = "GO-RESTFUL"
)

// New returns a log middleware uses stdout for outputing.
func New() *Middleware {
	return &Middleware{
		Writer:  os.Stdout,
		Project: _DefaultProject,
	}
}

// Middleware represents a log middleware.
type Middleware struct {
	Writer  io.Writer
	Project string
	timer   time.Time
}

func (this *Middleware) Processing(context *restful.Context) {
	this.timer = time.Now()
}

func (this Middleware) Processed(context *restful.Context) {
	end := time.Now()
	spent := end.Sub(this.timer)
	method := context.Request.Method()
	ip := context.Request.IPAddress()
	status := context.Status
	path := context.Request.RelativePath()

	fmt.Fprintf(this.Writer, "[%s] %v | %3d | %12v | %s | %-7s %s\n",
		this.Project,
		end.Format("2006-01-02 15:04:05"),
		status,
		spent,
		ip,
		method,
		path)
}
