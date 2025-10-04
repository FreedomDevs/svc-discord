package warns

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"svc-discord/errdefs"

	"github.com/bwmarrin/discordgo"
)

func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !slices.Contains([]string{"warn", "removewarn"}, i.ApplicationCommandData().Name) {
		return
	}

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
		roleToAdd, err := GiveWarn(target, guild, s)
		switch err {
		case errdefs.ErrMaxWarningsReached:
			result = fmt.Sprintf("❗ %s уже получил максимальное количество предупреждений", target.DisplayName())
		case errdefs.ErrRoleNotFound:
			var rnf *errdefs.RoleNotFoundError
			errors.As(err, &rnf)
			result = fmt.Sprintf("⚠️ Роль с ID %s не найдена", rnf.RoleID)
		case errdefs.ErrRoleAddFailed:
			result = "❗Ошибка при добавлении роли, подробнее в логах"
		case nil:
			result = fmt.Sprintf("✅%s получил предупреждение: %s", target.DisplayName(), roleToAdd)
		}
	case "removewarn":
		removeWarnResult, err := RemoveWarn(target, guild, s)
		switch err {
		case errdefs.ErrRoleNotFound:
			var rnf *errdefs.RoleNotFoundError
			errors.As(err, &rnf)
			result = fmt.Sprintf("⚠️ Роль с ID %s не найдена", rnf.RoleID)
		case errdefs.ErrRoleAddFailed:
			result = "❗Ошибка при добавлении роли, подробнее в логах"
		case errdefs.ErrNoWarnsToRemove:
			result = fmt.Sprintf("ℹ️ У %s нет предупреждений", target.DisplayName())
		case nil:
			switch removeWarnResult.Status {
			case WarnDecreased:
				result = fmt.Sprintf("✅ %s: уровень предупреждения снижен до %s", target.DisplayName(), removeWarnResult.Role)
			case WarnRemoved:
				result = fmt.Sprintf("✅ %s: все предупреждения сняты (удалена роль %s)", target.DisplayName(), removeWarnResult.Role)
			default:
				result = "❗Получен неизвестный код успешного выполнения от removewarn"
			}
		default:
			log.Println("Произошла неизвестная ошибка при исполнении removewarn", err)
			result = "❗Произошла неизвестная ошибка при исполнении removewarn"
		}
	}

	// Ответ пользователю
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: result,
		},
	})
}
