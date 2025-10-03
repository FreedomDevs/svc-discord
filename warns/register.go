package warns

import (
	"log"
	"svc-discord/config"

	"github.com/bwmarrin/discordgo"
)

func Register(session *discordgo.Session) {
	session.AddHandler(onInteractionCreate)

	userID := session.State.User.ID

	// Регистрация команд
	log.Println("Регистрация команды /warn")
	_, err := session.ApplicationCommandCreate(userID, config.GetGuildID(), &discordgo.ApplicationCommand{
		Name:        "warn",
		Description: "Give a warn to a user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to warn",
				Required:    true,
			},
		},
	})
	if err != nil {
		log.Println("Не удалось зарегистрировать команду /warn:", err)
	} else {
		log.Println("Команда /warn зарегистрирована")
	}

	log.Println("Регистрация команды /removewarn")
	_, err = session.ApplicationCommandCreate(userID, config.GetGuildID(), &discordgo.ApplicationCommand{
		Name:        "removewarn",
		Description: "Remove a warn from a user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to remove warn from",
				Required:    true,
			},
		},
	})
	if err != nil {
		log.Println("Не удалось зарегистрировать команду /removewarn:", err)
	} else {
		log.Println("Команда /removewarn зарегистрирована")
	}
}
