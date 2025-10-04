package warns

import (
	"log"
	"slices"
	"svc-discord/config"
	"svc-discord/errdefs"
	"svc-discord/utils"

	"github.com/bwmarrin/discordgo"
)

func HasAccessToWarn(member *discordgo.Member) bool {
	for _, r := range member.Roles {
		if slices.Contains(config.GetWarnAccessRoles(), r) {
			return true
		}
	}
	return false
}

func GetCurrentWarnLevel(member *discordgo.Member) int {
	warnLevels := config.GetWarnLevels()

	for i := len(warnLevels) - 1; i >= 0; i-- { // идём с конца
		if slices.Contains(member.Roles, warnLevels[i]) {
			return i + 1
		}
	}
	return 0
}

func GiveWarn(target *discordgo.Member, guild *discordgo.Guild, session *discordgo.Session) (string, error) {
	warnLevels := config.GetWarnLevels()
	currentWarnLevel := GetCurrentWarnLevel(target)

	index := currentWarnLevel - 1
	if index+1 >= len(warnLevels) {
		return "", errdefs.ErrMaxWarningsReached
	}

	roleToAdd := utils.GetRoleByID(guild, warnLevels[index+1])
	if roleToAdd == nil {
		return "", &errdefs.RoleNotFoundError{RoleID: warnLevels[index-1]}
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
		return "", errdefs.ErrRoleAddFailed
	}

	return roleToAdd.Name, nil
}

type WarnResult int

const (
	WarnNone WarnResult = iota
	WarnRemoved
	WarnDecreased
)

type RemoveWarnResult struct {
	Status WarnResult
	Role   string
}

func RemoveWarn(target *discordgo.Member, guild *discordgo.Guild, session *discordgo.Session) (*RemoveWarnResult, error) {
	warnLevels := config.GetWarnLevels()
	currentWarnLevel := GetCurrentWarnLevel(target)

	index := currentWarnLevel - 1
	if index < 0 {
		return nil, errdefs.ErrNoWarnsToRemove
	}

	roleToRemove := utils.GetRoleByID(guild, warnLevels[index])
	err := session.GuildMemberRoleRemove(guild.ID, target.User.ID, roleToRemove.ID)
	if err != nil {
		log.Println("Ошибка при удалении роли:", err)
	}

	if index-1 < 0 {
		return &RemoveWarnResult{Role: roleToRemove.Name, Status: WarnRemoved}, nil
	}

	roleToAdd := utils.GetRoleByID(guild, warnLevels[index-1])
	if roleToAdd == nil {
		return nil, &errdefs.RoleNotFoundError{RoleID: warnLevels[index-1]}
	}

	err = session.GuildMemberRoleAdd(guild.ID, target.User.ID, roleToAdd.ID)
	if err != nil {
		log.Println("Ошибка при добавлении роли:", err)
		return nil, errdefs.ErrRoleAddFailed
	}

	return &RemoveWarnResult{Role: roleToAdd.Name, Status: WarnDecreased}, nil
}
