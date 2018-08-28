package error_define

type ErrStruct struct {
	Code int
	Msg  string
}

var ErrorUnkown = ErrStruct{-101, "unknown error"}

// inner
var ErrorDatabase = ErrStruct{-201, "database error"}

// input
var ErrorInvalidJSON = ErrStruct{-301, "Invalid json input"}
var ErrValidation     = ErrStruct{403001, "验证错误"}
