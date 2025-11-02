package autorole

import (
	"log"
	"svc-discord/config"

	"github.com/bwmarrin/discordgo"
)

func onGuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	guildID := m.GuildID
	userID := m.User.ID
	roleID := config.GetOnJoinRole()

	err := s.GuildMemberRoleAdd(guildID, userID, roleID)
	if err != nil {
		log.Println("Не удалось выдать роль:", err)
		return
	}

	log.Println("Роль выдана пользователю:", userID)
}
