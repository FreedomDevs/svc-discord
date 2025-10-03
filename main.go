package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Println("Запуск svc-discord..")
	token := "Bot " + getEnv("DISCORD_BOT_TOKEN", "nil")

	dg, err := discordgo.New(token)
	if err != nil {
		log.Fatalln("Неверный токен:", err)
	}

	user, err := dg.User("@me")
	if err != nil {
		log.Fatalln("Ошибка при проверке токена:", err)
	}

	log.Printf("Токен рабочий! Бот: %s#%s\n", user.Username, user.Discriminator)
}
