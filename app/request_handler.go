package app

import (
	"io"
)

type MethodType int

const (
	GET MethodType = iota
	POST
	PUT
	DELETE
)

func (mt MethodType) String() string {
	return []string{"GET", "POST", "PUT", "DELETE"}[mt]
}

type RequestHandler struct {
	Method       MethodType
	URL          string
	Login        string
	Headers      map[string]string
	Body         io.Reader
	OutputFilter string
}
