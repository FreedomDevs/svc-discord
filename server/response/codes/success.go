package codes

import "svc-discord/server/response"

var (
	SuccessLiveOK       = response.SuccessResponseCode{Code: "LIVE_OK", Message: "svc-discord жив", Status: 200}
	SuccessReadyOK      = response.SuccessResponseCode{Code: "READY_OK", Message: "Сервис готов к приёму запросов", Status: 200}
	SuccessHealthOK     = response.SuccessResponseCode{Code: "HEALTH_OK", Message: "Сервис работает", Status: 200}
	SuccessUsersCountOK = response.SuccessResponseCode{Code: "USERS_COUNT_OK", Message: "Количество пользователей дискорд сервера", Status: 200}
)
