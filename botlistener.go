package main

import (
	"github.com/bwmarrin/discordgo"
)

func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	guild, err := s.Guild(i.GuildID)
	if err != nil {
		panic(err)
	}

	author := i.Member
	targetOption := i.ApplicationCommandData().Options
	if len(targetOption) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ User not found.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	user := targetOption[0].UserValue(s) // *discordgo.User
	if user == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ User not found.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	target, err := s.GuildMember(guild.ID, user.ID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ User not on discord server.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	// Проверка прав через кастомную функцию
	if !HasAccessToWarn(author) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ You do not have permission to use this command.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	var result string
	switch i.ApplicationCommandData().Name {
	case "warn":
		result = GiveWarn(target, guild, s)
	case "removewarn":
		result = RemoveWarn(target, guild, s)
	default:
		result = "Unknown command."
	}

	// Ответ пользователю
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: result,
		},
	})
}
