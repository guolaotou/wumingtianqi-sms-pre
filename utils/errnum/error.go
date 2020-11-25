package errnum

import "fmt"

var (
	// common errors
	OK                  = &Er{Code: 0, Message: "OK"}
	InternalServerError = &Er{Code: 40001, Message: "Internal server error"}
	ParamsError         = &Er{Code: 40002, Message: "Params error"}
	DbError             = &Er{Code: 40003, Message: "Db error"}
	S2SError            = &Er{Code: 40004, Message: "S2S error"}
	WxError             = &Er{Code: 40005, Message: "Wx error"}
	StrangeError        = &Er{Code: 40006, Message: "Strange error"}
	RemainingNotEnough  = &Er{Code: 40007, Message: "Remaining Not Enough"}
	UserAlreadyVip      = &Er{Code: 40008, Message: "User Already Vip"}
	ErrParsingPostJson  = &Er{Code: 40009, Message: "ErrParsingPostJson"}
	ErrNoAuth           = &Er{Code: 40010, Message: "No Auth"}
	// login
	ErrTokenNotExist = &Er{Code: 50100, Message: "token not exist"}

	// order
	ErrTelOrderChanceInsufficient = &Er{Code: 50200, Message:"Tel Order Chance Insufficient"}
	ErrOrderNotFound = &Er{Code: 50201, Message:"Order Not Found"}


	// Continue add err types ...
)

// 定义两种错误类型，一个是Er, 一个是Err，分别用来返回前端错误以及用作后台log记录详细错误日志

// internal error
type Er struct {
	Code    int
	Message string
}

func (er *Er) Error() string {
	return er.Message
}

// api error
type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func New(er *Er, err error) *Err {
	return &Err{Code: er.Code, Message: er.Message, Err: err}
}

func (err *Err) AddMsg(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) AddMsgF(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Er:
		return typed.Code, typed.Message
	case *Err:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
