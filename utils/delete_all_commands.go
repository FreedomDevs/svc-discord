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
		log.Println("Failed to fetch global commands:", err)
	} else {
		for _, cmd := range globalCommands {
			err := session.ApplicationCommandDelete(userID, "", cmd.ID)
			if err != nil {
				log.Println("Failed to delete global command:", cmd.Name, err)
			} else {
				log.Println("Deleted global command:", cmd.Name)
			}
		}
	}

	// 2. Удаляем команды на всех guild
	for _, g := range session.State.Guilds {
		log.Println("Clearing commands on guild:", g.ID)

		guildCommands, err := session.ApplicationCommands(userID, g.ID)
		if err != nil {
			log.Println("Cannot fetch commands for guild", g.ID, err)
			continue
		}

		for _, cmd := range guildCommands {
			err := session.ApplicationCommandDelete(userID, g.ID, cmd.ID)
			if err != nil {
				log.Println("Failed to delete command:", cmd.Name, err)
			} else {
				log.Println("Deleted command:", cmd.Name)
			}
		}
	}

	log.Println("All commands cleared!")
}
