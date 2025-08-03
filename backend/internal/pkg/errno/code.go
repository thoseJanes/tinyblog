package errno


var(
	Ok = &Errno{HttpCode: 200, ErrCode: "", Message: ""}

	InternalServerErr = &Errno{HttpCode: 500, ErrCode: "InternalError", Message: "Internal server error."}

	ErrUnauthorized = &Errno{HttpCode: 401, ErrCode: "AuthFailure.Unauthorized", Message: "Unauthorized."}

	ErrTokenInvalid = &Errno{HttpCode: 401, ErrCode: "AuthFailure.TokenInvalid", Message: "Can't resolve json web token."}

	ErrInvalidParameter = &Errno{HttpCode: 400, ErrCode: "InvalidParamter", Message: "Failed to validate parameter."}

	ErrPageNotFound = &Errno{HttpCode: 404, ErrCode: "ResourceNotFound.PageNotFound", Message: "Page not found"}

	ErrUserAlreadyExist = &Errno{HttpCode: 400, ErrCode: "FailedOperation.UserAlreadyExist", Message: "User already exists."}

	ErrPasswordIncorrect = &Errno{HttpCode: 401, ErrCode: "AuthFailure.PasswordIncorrect", Message: "Password incorrect."};

	ErrSignToken = &Errno{HttpCode: 401, ErrCode: "AuthFailure.SignTokenErr", Message: "Sign token error."}

	ErrUserNotFound = &Errno{HttpCode: 404, ErrCode: "ResourceNotFound.UserNotFound", Message: "User not found"}

	ErrPostNotFound = &Errno{HttpCode: 404, ErrCode: "ResourceNotFound.PostNotFound", Message: "Post not found"}
)
