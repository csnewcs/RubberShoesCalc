package Bot

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func InitBot(token string, slashCommands map[string]func(*discordgo.Session, *discordgo.InteractionCreate)) (*discordgo.Session, error) {
	session, _ := discordgo.New("Bot " + token)
	session.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if interaction.Type == discordgo.InteractionApplicationCommand {
			if function, ok := slashCommands[interaction.ApplicationCommandData().Name]; ok {
				function(session, interaction)
			}
		}
	})
	session.AddHandler(func(session *discordgo.Session, ready *discordgo.Ready) {
		slog.Info("Bot is ready")
	})
	err := session.Open()
	return session, err
}

func KillBot(session *discordgo.Session) {
	//something...
	session.Close()
}
