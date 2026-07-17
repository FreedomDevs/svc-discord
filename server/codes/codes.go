package codes

import "github.com/FreedomDevs/svcLibs/go/svcLibs"

var (
	ErrGuildIsNull = svcLibs.ErrorResponseCode{Code: "GUILD_IS_NULL", Message: "Был запрошен дискорд сервер, но получен nil", Status: 500}
)

var (
	SuccessUsersCountOK = svcLibs.SuccessResponseCode{Code: "USERS_COUNT_OK", Message: "Количество пользователей дискорд сервера", Status: 200}
)
