package errno


var(
	Ok = &Errno{HttpCode: 200, ErrCode: "", Message: ""}

	InternalServerErr = &Errno{HttpCode: 500, ErrCode: "InternalError", Message: "Internal server error."}

	ErrUnauthorized = &Errno{HttpCode: 401, ErrCode: "AuthFailure.Unauthorized", Message: "Unauthorized."}

	ErrTokenInvalid = &Errno{HttpCode: 401, ErrCode: "AuthFailure.TokenInvalid", Message: "Can't resolve json web token."}
)
