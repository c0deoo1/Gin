// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !jsoniter

package json

// 这里利用了go build 的tags特性，默认使用标准库的Json来实现
import "encoding/json"

var (
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
