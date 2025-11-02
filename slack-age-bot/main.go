package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("--------")
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "token") //xoxb-9820101296102-9825746440756-0rp6fc9otsAGFbqXnzilDhIg
	os.Setenv("SLACK_APP_TOKEN", "token") //xapp-1-A09R0DWPY0G-9818716564149-c14316aaaf2ed32b5106ef0ac71ef9a9bd327a910ce8f33aecbf2ae38f2e9c66

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Examples: "my yob is 1990",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Please enter a valid year (e.g., 1990).")
				return
			}
			currentYear := time.Now().Year()
			age := currentYear - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
