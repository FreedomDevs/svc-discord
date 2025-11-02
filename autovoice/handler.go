package autovoice

import (
	"log"
	"strings"
	"svc-discord/config"

	"github.com/bwmarrin/discordgo"
)

func onVoiceStateUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	if v.GuildID != config.GetGuildID() {
		log.Printf("ERROR: Попытка входа в войс с неизвестного GuildID, %s\n", v.GuildID)
		return
	}

	if v.BeforeUpdate != nil {
		channel, err := s.Channel(v.BeforeUpdate.ChannelID)
		if err != nil {
			log.Printf("Ошибка при обработке ивента войса: %s", err.Error())
			return
		}

		if strings.HasPrefix(channel.Name, config.GetChannelPrefix()) && channel.MemberCount == 0 {
			_, err := s.ChannelDelete(v.BeforeUpdate.ChannelID)
			if err != nil {
				log.Printf("Ошибка при удалении войса после неактивности: %s", err.Error())
				return
			}
		}
	}

	if v.ChannelID == config.GetVoiceCreateChannelID() {
		channelData := &discordgo.GuildChannelCreateData{
			Name:      config.GetChannelPrefix() + v.Member.DisplayName(),
			Type:      discordgo.ChannelTypeGuildVoice,
			ParentID:  config.GetAutoVoiceCategoryID(),
			UserLimit: 5,
		}

		channel, err := s.GuildChannelCreateComplex(config.GetGuildID(), *channelData)
		if err != nil {
			log.Println("Ошибка при создании канала:", err)
			return
		}

		log.Println("Создан канал: " + channel.Name)

		err = s.GuildMemberMove(config.GetGuildID(), v.UserID, &channel.ID)
		if err != nil {
			log.Println("Не удалось переместить человека из одного войса в другой, " + err.Error())
			return
		}

		log.Println("Войс создан и чел перемещён")
	}
}
