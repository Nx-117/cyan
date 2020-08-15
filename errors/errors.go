package errors

/**
自定义Error包
*/

import (
	"encoding/json"
)

type Err struct {
	Code int
	Msg  string
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func New(code int, msg string) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}
