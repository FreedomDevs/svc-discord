package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ClearAllCommands(session *discordgo.Session) {
	userID := session.State.User.ID

	// 1. Удаляем глобальные команды
	globalCommands, err := session.ApplicationCommands(userID, "")
	if err != nil {
		log.Println("Не удалось получить глобальные команды:", err)
	} else {
		for _, cmd := range globalCommands {
			err := session.ApplicationCommandDelete(userID, "", cmd.ID)
			if err != nil {
				log.Println("Не удалось удалить глобальную команду:", cmd.Name, err)
			} else {
				log.Println("Удалена глобальная команда:", cmd.Name)
			}
		}
	}

	// 2. Удаляем команды на всех guild
	for _, g := range session.State.Guilds {
		log.Println("Удаление команд на дискорд сервере:", g.ID)

		guildCommands, err := session.ApplicationCommands(userID, g.ID)
		if err != nil {
			log.Println("Не удалось получить список команд на дискорд сервере:", g.ID, err)
			continue
		}

		for _, cmd := range guildCommands {
			err := session.ApplicationCommandDelete(userID, g.ID, cmd.ID)
			if err != nil {
				log.Println("Не удалось удалить команду:", cmd.Name, err)
			} else {
				log.Println("Удалена команда:", cmd.Name)
			}
		}
	}

	log.Println("Все старые команды удалены!")
}
