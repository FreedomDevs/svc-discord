package autorole

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Register(session *discordgo.Session) {
	session.AddHandler(onGuildMemberAdd)
	log.Println("OnGuildMemberAdd ивент зарегистрирован")
}
