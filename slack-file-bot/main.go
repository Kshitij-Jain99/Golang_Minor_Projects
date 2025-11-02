package main

import (
	"context"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-9820101296102-9822954071430-Zxn6VIgJoueWxjP8y3Ci4h85")
	os.Setenv("CHANNEL_ID", "C09PQLC8615")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channel := os.Getenv("CHANNEL_ID")

	files := []string{"pplx-at-work.pdf"}

	for _, f := range files {
		// Check if file exists and can be accessed
		fileInfo, err := os.Stat(f)
		if err != nil {
			fmt.Printf("Error accessing file %s: %v\n", f, err)
			continue
		}

		fmt.Printf("File %s found, size: %d bytes\n", f, fileInfo.Size())

		// Try using Reader instead of File path
		fileReader, err := os.Open(f)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", f, err)
			continue
		}
		defer fileReader.Close()

		_, err = api.UploadFileV2Context(context.Background(), slack.UploadFileV2Parameters{
			Filename: f,
			FileSize: int(fileInfo.Size()),
			Channel:  channel,
			Reader:   fileReader, // Use Reader instead of File
		})
		if err != nil {
			fmt.Printf("Error uploading %s: %v\n", f, err)
			continue
		}

		fmt.Printf("✅ Uploaded: %s\n", f)
	}
}
