// +build appengine

// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

// 如果编译的时候增加了tags appengine 参数，表示运行的目标平台为Google AppEngine
func init() {
	defaultAppEngine = true
}
