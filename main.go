package main

import (
	"log"
	"os"
	"os/signal"
	"svc-discord/config"
	"svc-discord/utils"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Println("Запуск svc-discord..")
	config.LoadEnvVars()

	dg, err := discordgo.New(config.GetDiscordBotToken())
	if err != nil {
		log.Fatalln("Неверный токен:", err)
	}

	user, err := dg.User("@me")
	if err != nil {
		log.Fatalln("Ошибка при проверке токена:", err)
	}

	log.Printf("Токен рабочий! Бот: %s#%s\n", user.Username, user.Discriminator)

	dg.AddHandler(onInteractionCreate)

	err = dg.Open()
	if err != nil {
		panic(err)
	}
	log.Println("Bot is running... Press CTRL+C to exit.")

	utils.ClearAllCommands(dg)

	// Регистрация команд
	_, err = dg.ApplicationCommandCreate(dg.State.User.ID, config.GetGuildID(), &discordgo.ApplicationCommand{
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
		log.Fatalln("Cannot create slash command:", err)
	}

	_, err = dg.ApplicationCommandCreate(dg.State.User.ID, config.GetGuildID(), &discordgo.ApplicationCommand{
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
		log.Fatalln("Cannot create slash command:", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
