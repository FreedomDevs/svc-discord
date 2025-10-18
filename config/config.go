package config

import "strings"

var discordBotToken string
var guildID string
var warnLevels []string
var warnAccessRoles []string
var isProd bool

func GetDiscordBotToken() string {
	return discordBotToken
}
func GetGuildID() string {
	return guildID
}
func GetWarnLevels() []string {
	return warnLevels
}
func GetWarnAccessRoles() []string {
	return warnAccessRoles
}
func GetIsProd() bool {
	return isProd
}

func LoadEnvVars() {
	discordBotToken = "Bot " + getEnvOrPanic("DISCORD_BOT_TOKEN")
	guildID = getEnvOrDefault("GUILD_ID", "1157195772115292252")

	exampleWarnLevels := "1391167522446774272,1391168547538866217,1391168616271056957"
	warnLevels = strings.Split(getEnvOrDefault("WARN_LEVELS", exampleWarnLevels), ",")

	exampleAccessRoles := "1274669048902193186,1391097009032794183,1274668939519070259,1274669011384143873"
	warnAccessRoles = strings.Split(getEnvOrDefault("WARN_ACCESS_ROLES", exampleAccessRoles), ",")

	isProd = parseBoolSafe(getEnvOrDefault("IS_PROD", "false"))
}
