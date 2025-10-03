package utils

import "github.com/bwmarrin/discordgo"

func GetRoleByID(guild *discordgo.Guild, roleID string) *discordgo.Role {
	for _, role := range guild.Roles {
		if role.ID == roleID {
			return role
		}
	}
	return nil // роль не найдена
}
