package main

import (
	"log"
	"os"
	"os/signal"
	"svc-discord/config"
	"svc-discord/utils"
	"svc-discord/warns"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Println("Запуск svc-discord..")
	config.LoadEnvVars()

	session, err := discordgo.New(config.GetDiscordBotToken())
	if err != nil {
		log.Fatalln("Неверный токен:", err)
	}

	user, err := session.User("@me")
	if err != nil {
		log.Fatalln("Ошибка при проверке токена:", err)
	}

	log.Printf("Токен рабочий! Бот: %s#%s\n", user.Username, user.Discriminator)

	err = session.Open()
	if err != nil {
		panic(err)
	}
	log.Println("Бот запущен... Нажмите CTRL+C чтобы завершить.")

	utils.ClearAllCommands(session)

	go warns.Register(session)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	session.Close()
}
