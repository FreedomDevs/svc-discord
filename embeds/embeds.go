package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var TicketEmbed = &discordgo.MessageEmbed{
	Title:       "🎫 Обращение в поддержку",
	Description: "Нажмите на кнопку ниже, чтобы обратиться в поддержку.\n\n📌 **Опишите вашу проблему максимально подробно — это ускорит решение.**",
	Color:       0x00FF00,
}

var TicketCreatedEmbed = &discordgo.MessageEmbed{
	Title:       "✅ Тикет создан",
	Description: "💬 Здесь вы можете обсудить свою проблему с администрацией сервера.\n⏳ **Ожидайте ответа, администраторы уведомлены.**",
	Color:       0x0000FF,
}

var ReportInfoEmbed = &discordgo.MessageEmbed{
	Title:       "📢 Подача жалобы",
	Description: "Если вы хотите сообщить о нарушении, нажмите кнопку ниже. 🔽\n\n⚠️ **Пожалуйста, убедитесь, что ваша жалоба содержит достоверную информацию.**",
	Color:       0xFFA500, // оранжевый
}

func ReportSubmittedEmbed(nick, reason string, reporter *discordgo.User) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "📨 Новая жалоба",
		Fields: []*discordgo.MessageEmbedField{
			{Name: "👤 **Нарушитель**", Value: nick, Inline: false},
			{Name: "📝 **Причина**", Value: reason, Inline: false},
			{Name: "📩 **Жалоба от**", Value: "<@" + reporter.ID + ">", Inline: false},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: reporter.AvatarURL("256"),
		},
		Color:     0xFF0000,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

var SuccessReportEmbed = &discordgo.MessageEmbed{
	Title:       "✅ Успех",
	Description: "📬 **Жалоба успешно отправлена!**\n💬 В случае необходимости с вами свяжутся для более подробного выяснения обстоятельств.",
	Color:       0xFFA500,
}

var ErrorReportChannelNotFoundEmbed = &discordgo.MessageEmbed{
	Title:       "❌ Ошибка",
	Description: "Канал для жалоб не найден.\n📢 Пожалуйста, сообщите об этом администратору сервера.",
	Color:       0xFF0000,
}

func SuccessEmbed(description string, title ...string) *discordgo.MessageEmbed {
	t := "✅ Успех"
	if len(title) > 0 {
		t = title[0]
	}
	return &discordgo.MessageEmbed{
		Title:       t,
		Description: description,
		Color:       0x00FF00,
	}
}

func ErrorEmbed(description string, title ...string) *discordgo.MessageEmbed {
	t := "❌ Ошибка"
	if len(title) > 0 {
		t = title[0]
	}
	return &discordgo.MessageEmbed{
		Title:       t,
		Description: description,
		Color:       0xFF0000,
	}
}

func InfoEmbed(description string, title ...string) *discordgo.MessageEmbed {
	t := "ℹ️ Информация"
	if len(title) > 0 {
		t = title[0]
	}
	return &discordgo.MessageEmbed{
		Title:       t,
		Description: description,
		Color:       0x00FFFF,
	}
}

func WarningEmbed(description string, title ...string) *discordgo.MessageEmbed {
	t := "⚠️ Предупреждение"
	if len(title) > 0 {
		t = title[0]
	}
	return &discordgo.MessageEmbed{
		Title:       t,
		Description: description,
		Color:       0xFFFF00,
	}
}

func CreateKickMenuEmbed(voiceChannelName string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "🚪 Выбор участника для кика",
		Description: "Выберите участника, которого хотите **кикнуть** из канала `" + voiceChannelName + "`.\n\n" +
			"⚠️ **Помните, что кик будет немедленным, хотя участник сможет вернутся обратно.**",
		Color: 0xFF0000,
	}
}

func CreateControlPanelEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "🛠️ Панель управления каналом",
		Description: "Используйте кнопки ниже для управления каналами:\n" +
			"➕ **Увеличить максимальное число пользователей на 1**, \n" +
			"➖ **Уменьшить максимальное число пользователей на 1**, \n" +
			"❌ **Кикнуть участника из канала**.",
		Color: 0x00FFFF,
	}
}
