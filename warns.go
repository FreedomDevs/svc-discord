package main

import (
	"fmt"
	"log"
	"slices"
	"strings"
	"svc-discord/utils"

	"github.com/bwmarrin/discordgo"
)

var warnLevels []string
var warnAccessRoles []string

func init() {
	exampleWarnLevels := "1391167522446774272,1391168547538866217,1391168616271056957"
	warnLevels = strings.Split(getEnv("WARN_LEVELS", exampleWarnLevels), ",")

	exampleAccessRoles := "1274669048902193186,1391097009032794183,1274668939519070259,1274669011384143873"
	warnAccessRoles = strings.Split(getEnv("WARN_ACCESS_ROLES", exampleAccessRoles), ",")
}

func HasAccessToWarn(member *discordgo.Member) bool {
	for _, r := range member.Roles {
		if slices.Contains(warnAccessRoles, r) {
			return true
		}
	}
	return false
}

func GetCurrentWarnLevel(member *discordgo.Member) int {
	for i := len(warnLevels) - 1; i >= 0; i-- { // идём с конца
		if slices.Contains(member.Roles, warnLevels[i]) {
			return i + 1
		}
	}
	return 0
}

func GiveWarn(target *discordgo.Member, guild *discordgo.Guild, session *discordgo.Session) string {
	currentWarnLevel := GetCurrentWarnLevel(target)

	index := currentWarnLevel - 1
	if index+1 >= len(warnLevels) {
		return fmt.Sprintf("❗ %s уже получил максимальное количество предупреждений", target.DisplayName())
	}

	roleToAdd := utils.GetRoleByID(guild, warnLevels[index+1])
	if roleToAdd == nil {
		return "⚠️ Роль с ID $nextWarnRoleId не найдена"
	}

	if index >= 0 {
		roleToRemove := utils.GetRoleByID(guild, warnLevels[index])
		err := session.GuildMemberRoleRemove(guild.ID, target.User.ID, roleToRemove.ID)
		if err != nil {
			log.Println("Ошибка при удалении роли:", err)
		}
	}

	err := session.GuildMemberRoleAdd(guild.ID, target.User.ID, roleToAdd.ID)
	if err != nil {
		log.Println("Ошибка при добавлении роли:", err)
		return "Ошибка при добавлении роли"
	}

	return fmt.Sprintf("✅%s получил предупреждение: %s", target.DisplayName(), roleToAdd.Name)
}

func RemoveWarn(target *discordgo.Member, guild *discordgo.Guild, session *discordgo.Session) string {
	currentWarnLevel := GetCurrentWarnLevel(target)

	index := currentWarnLevel - 1
	if index < 0 {
		return fmt.Sprintf("ℹ️ У %s нет предупреждений", target.DisplayName())
	}

	roleToRemove := utils.GetRoleByID(guild, warnLevels[index])
	err := session.GuildMemberRoleRemove(guild.ID, target.User.ID, roleToRemove.ID)
	if err != nil {
		log.Println("Ошибка при удалении роли:", err)
	}

	if index-1 < 0 {
		return fmt.Sprintf("✅ %s: все предупреждения сняты (удалена роль %s)", target.DisplayName(), roleToRemove.Name)
	}

	roleToAdd := utils.GetRoleByID(guild, warnLevels[index-1])
	if roleToAdd == nil {
		return "⚠️ Роль с ID $nextWarnRoleId не найдена"
	}

	err = session.GuildMemberRoleAdd(guild.ID, target.User.ID, roleToAdd.ID)
	if err != nil {
		log.Println("Ошибка при добавлении роли:", err)
		return "Ошибка при добавлении роли"
	}

	return fmt.Sprintf("✅ %s: уровень предупреждения снижен до %s", target.DisplayName(), roleToAdd.Name)
}
