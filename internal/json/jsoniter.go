// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build jsoniter

// 这里利用了go build 的tags特性，只有在传入go build -tags=jsoniter .的时候，才会编译这个文件
// jsoniter比标准库的实现快很多
package json

import "github.com/json-iterator/go"

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gin/json package.
	NewEncoder = json.NewEncoder
)
