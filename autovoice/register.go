package autovoice

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Register(session *discordgo.Session) {
	session.AddHandler(onVoiceStateUpdate)
	log.Println("OnVoice ивент зарегистрирован")
}
