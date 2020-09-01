// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"reflect"
	"sync"

	"github.com/go-playground/validator/v10"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ StructValidator = &defaultValidator{}

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
// 通过struct的tags来附加额外的元数据，做数据的校验
// github.com/go-playground/validator/v10 包含了一些常用的校验规则，同时支持自定义校验规则
// https://godoc.org/gopkg.in/go-playground/validator.v10
// 比如：Required规则
// This validates that the value is not the data types default zero value.
// For numbers ensures value is not zero. For strings ensures value is not "".
// For slices, maps, pointers, interfaces, channels and functions ensures the value is not nil.
func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	// 反射的用法
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	// 只会校验struct类型
	if valueType == reflect.Struct {
		// 真正使用到的时候才做初始化
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://godoc.org/gopkg.in/go-playground/validator.v8
func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

// 懒加载的方式，真正使用到的时候才做init初始化
func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}
