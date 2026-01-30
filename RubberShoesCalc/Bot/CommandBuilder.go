package Bot

import (
	"github.com/bwmarrin/discordgo"
)

type BotCommandBuilder struct {
	Name        string
	Description string
	Function    func(*discordgo.Session, *discordgo.InteractionCreate)
	Args        []*discordgo.ApplicationCommandOption
}

func NewBotCommandBuilder(name string) BotCommandBuilder {
	return BotCommandBuilder{Name: name}
}
func (botCommandBuilder BotCommandBuilder) WithFunction(function func(*discordgo.Session, *discordgo.InteractionCreate)) BotCommandBuilder {
	botCommandBuilder.Function = function
	return botCommandBuilder
}
func (botCommandBuilder BotCommandBuilder) WithDescription(description string) BotCommandBuilder {
	botCommandBuilder.Description = description
	return botCommandBuilder
}
func (botCommandBuilder BotCommandBuilder) AddArg(arg *discordgo.ApplicationCommandOption) BotCommandBuilder {
	botCommandBuilder.Args = append(botCommandBuilder.Args, arg)
	return botCommandBuilder
}
