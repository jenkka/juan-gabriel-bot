package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jenkka/juan-gabriel-bot/config"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore your own messages
	if m.Author.ID == BotId {
		return
	}

	if fmt.Contains(m.Content, "trece") == "<@"+BotId+"> ping" || m.Content == "<@"+BotId+"> ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}

	//if m.Content == config.BotPrefix+"ping" {
	//	_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	//}
}
