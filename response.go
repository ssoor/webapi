package webapi

import (
	"net/http"
)

// Context represents a context.
type Response struct {
	Status int
	Data   interface{}
	Header http.Header
}

func NewResponse(rawstatus int, rawdata interface{}, rawheader http.Header) *Response {
	return &Response{
		Status: rawstatus,
		Data:   rawdata,
		Header: rawheader,
	}
}
