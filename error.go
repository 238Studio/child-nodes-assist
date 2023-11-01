package util

//进行错误处理

import (
	"errors"
	"runtime"
)

// CustomError 是自定义的错误类型，包括错误消息和堆栈信息。
type CustomError struct {
	ErrorLevel   int    //错误等级
	ErrorType    int    //错误类型
	ErrorMessage string //错误信息
	Stack        string //调用堆栈

	CreatedAt int64 `gorm:"autoUpdateTime:milli"` //创建时间

	isNeedSpecialHandle bool //是否需要特殊处理
}

// NewError 创建一个新的 CustomError。
// 传入：错误等级 错误类型 错误信息
// 传出：CustomError
func NewError(errorLevel int, errorType int, isNeedSpecialHandle bool, err error) *CustomError {
	// 获取堆栈信息
	stack := getStackTrace()
	return &CustomError{
		ErrorLevel:          errorLevel,
		ErrorType:           errorType,
		ErrorMessage:        err.Error(),
		Stack:               stack,
		isNeedSpecialHandle: isNeedSpecialHandle,
	}
}

// Error 实现了错误接口的 Error 方法。
// 传入：CustomError
// 传出：错误信息
func (e *CustomError) Error() string {
	return e.ErrorMessage + "\n" + e.Stack
}

// getStackTrace 获取当前堆栈信息。
// 传入：无
// 传出：堆栈信息
func getStackTrace() string {
	buf := make([]byte, 4096)
	runtime.Stack(buf, false)
	stackTrace := string(buf)
	return stackTrace
}

// ErrToStruct 将错误转换为结构体
// 传入：错误信息
// 传出：错误结构体
func ErrToStruct(err error) *CustomError {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr
	}
	return nil
}

// ExtractStackTrace 从错误消息中提取堆栈信息。
// 传入：错误信息
// 传出：堆栈信息
func ExtractStackTrace(err error) string {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.Stack
	}
	return ""
}

// ExtractErrorMessage 从错误消息中提取错误消息。
// 传入：错误信息
// 传出：错误消息
func ExtractErrorMessage(err error) string {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.ErrorMessage
	}
	return err.Error()
}

// ExtractErrorType 从错误消息中提取错误类型。
// 传入：错误信息
// 传出：错误类型
func ExtractErrorType(err error) int {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.ErrorType
	}
	return -1
}

// ExtractErrorLevel 从错误消息中提取错误等级。
// 传入：错误信息
// 传出：错误等级
func ExtractErrorLevel(err error) int {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.ErrorLevel
	}
	return -1
}

// ExtractIsNeedSpecialHandle 从错误消息中提取是否需要特殊处理。
// 传入：错误信息
// 传出：是否需要特殊处理
func ExtractIsNeedSpecialHandle(err error) bool {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.isNeedSpecialHandle
	}
	return false
}
