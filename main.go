package main

import (
	"fmt"

	"github.com/jenkka/juan-gabriel-bot/bot"
	"github.com/jenkka/juan-gabriel-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Start()

	<-make(chan struct{})

	return
}
