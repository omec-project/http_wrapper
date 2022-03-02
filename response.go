// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

package http_wrapper

import (
	"net/http"
)

type Response struct {
	Header http.Header
	Status int
	Body   interface{}
}

func NewResponse(code int, h http.Header, body interface{}) (ret *Response) {
	ret = &Response{}
	ret.Status = code
	ret.Header = h
	ret.Body = body
	return
}
