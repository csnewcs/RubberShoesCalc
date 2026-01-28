package main

import (
	b "github.com/csnewcs/RubberShoesCalc/Bot"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	conf, err := LoadConfig("./config.json")
	if err != nil {
		slog.Error("Error occured when load config file", err)
	}
	bot, err := b.InitBot(conf.Token, nil)
	if err != nil {
		slog.Error("Error occured when turn on the bot", err)
	}
	slog.Info("Bot is started...", "ID", bot.State.User.ID)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	slog.Info("Process is interrupted... exit")
	b.KillBot(bot)
}
