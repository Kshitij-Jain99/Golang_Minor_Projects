package main

import (
	"context"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-9820101296102-9822954071430-5gI5PpRJESs3IKa2fKVjICfW")
	os.Setenv("CHANNEL_ID", "C09PQLC8615")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channel := os.Getenv("CHANNEL_ID")

	files := []string{"pplx-at-work.pdf"}

	for _, f := range files {
		_, err := api.UploadFileV2Context(context.Background(), slack.UploadFileV2Parameters{
			File:    f,
			Channel: channel,
		})
		if err != nil {
			fmt.Printf("Error uploading %s: %v\n", f, err)
			continue
		}

		fmt.Printf("✅ Uploaded: %s\n", f)
	}
}
