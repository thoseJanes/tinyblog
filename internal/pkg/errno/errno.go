package errno

import "fmt"



type Errno struct{
	HttpCode int
	ErrCode string
	Message string
}


func (e *Errno) Error() string {
	return e.Message
}

func (e *Errno) SetMessage(format string, args ...interface{}) *Errno {
	e.Message = fmt.Sprintf(format, args...)
	return e
}



func Decode(e error) (httpCode int, errCode string, message string) {
	var ret *Errno
	if e == nil {
		ret = Ok
	}else if err, ok := e.(*Errno); ok {
		ret = err
	}else{
		ret = InternalServerErr
	}

	httpCode, errCode, message = ret.HttpCode, ret.ErrCode, ret.Message
	return
}