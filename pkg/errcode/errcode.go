package errcode

import (
	"fmt"
	"runtime"
)

type ErrCode struct {
	Code  string
	Desc  string
	Http  int
	Trace string
	Err   error
}

const DefaultCode = "internal_error"

func NewDefaultErr() *ErrCode {
	_, file, line, ok := runtime.Caller(2)
	var trace string
	if ok {
		trace = fmt.Sprintf("%s:%d", file, line)
	}
	e := New(DefaultCode)
	e.Trace = trace
	return e
}
func New(code string) *ErrCode {
	_, file, line, ok := runtime.Caller(1)
	var trace string
	if ok {
		trace = fmt.Sprintf("%s:%d", file, line)
	}
	return &ErrCode{
		Code:  code,
		Trace: trace,
	}
}

func (e *ErrCode) WithMsg(msg ...string) *ErrCode {
	e.Desc = fmt.Sprint(msg)
	return e
}
func (e *ErrCode) WithCode(code string) *ErrCode {
	e.Code = code
	return e
}
func (e *ErrCode) WithHttp(status int) *ErrCode {
	e.Http = status
	return e
}
func (e *ErrCode) WithErr(err error) *ErrCode {
	e.Err = err
	return e
}

func (e *ErrCode) GetHttp() int {
	if e.Http >= 200 && e.Http < 510 {
		return e.Http
	}
	return 400
}

func (e *ErrCode) Error() string {
	if e == nil {
		return DefaultCode
	}
	return e.Code
}

func (e *ErrCode) Log() string {

	return fmt.Sprintf("code: %s | err: %v | description: %s | trace: %s", e.Code, e.Err, e.Desc, e.Trace)
}
