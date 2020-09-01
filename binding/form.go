// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"
)

const defaultMemory = 32 << 20

type formBinding struct{}
type formPostBinding struct{}
type formMultipartBinding struct{}

func (formBinding) Name() string {
	return "form"
}

func (formBinding) Bind(req *http.Request, obj interface{}) error {
	// 标准库的ParseForm函数的注释比较详细
	// ParseForm populates r.Form and r.PostForm.
	//
	// For all requests, ParseForm parses the raw query from the URL and updates
	// r.Form.
	//
	// For POST, PUT, and PATCH requests, it also parses the request body as a form
	// and puts the results into both r.PostForm and r.Form. Request body parameters
	// take precedence over URL query string values in r.Form.
	//
	// For other HTTP methods, or when the Content-Type is not
	// application/x-www-form-urlencoded, the request Body is not read, and
	// r.PostForm is initialized to a non-nil, empty value.
	//
	// If the request Body's size has not already been limited by MaxBytesReader,
	// the size is capped at 10MB.
	//
	// ParseMultipartForm calls ParseForm automatically.
	// ParseForm is idempotent.
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	if err := mapForm(obj, req.Form); err != nil {
		return err
	}
	return validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := mapForm(obj, req.PostForm); err != nil {
		return err
	}
	return validate(obj)
}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	if err := mappingByPtr(obj, (*multipartRequest)(req), "form"); err != nil {
		return err
	}

	return validate(obj)
}
