// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errorx

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknownError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}

// 为某个枚举单独设置错误码
func IsBus(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_BUS.String() && e.Code == 1
}

// 为某个枚举单独设置错误码
func ErrorBus(format string, args ...interface{}) *errors.Error {
	return errors.New(1, ErrorReason_BUS.String(), fmt.Sprintf(format, args...))
}