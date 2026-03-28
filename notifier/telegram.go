package notifier

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"golang-job-scraper/parser"
)

func SendTelegram(job parser.Job) {
	botToken := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")

	msg := url.QueryEscape(fmt.Sprintf(
		"🔥 Job Found!\n\nTitle: %s\nLink: %s",
		job.Title,
		job.Link,
	))

	url := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken,
		chatID,
		msg,
	)

	http.Get(url)
}
