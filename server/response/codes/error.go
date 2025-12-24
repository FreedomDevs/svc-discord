package codes

import "svc-discord/server/response"

var (
	ErrGuildIsNull   = response.ErrorResponseCode{Code: "GUILD_IS_NULL", Message: "Был запрошен дискорд сервер, но получен nil", Status: 500}
	ErrInternalError = func(err error) response.ErrorResponseCode {
		return response.ErrorResponseCode{Code: "INTERNAL_ERROR", Message: err.Error(), Status: 500}
	}
)
