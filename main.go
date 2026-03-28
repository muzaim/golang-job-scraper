package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"golang-job-scraper/notifier"
	"golang-job-scraper/parser"
	"golang-job-scraper/scraper"
	"golang-job-scraper/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	url := os.Getenv("URL")

	cache := storage.NewCache()
	ticker := time.NewTicker(10 * time.Minute)

	for {
		log.Println("Running scraper...")

		html, err := scraper.Fetch(url)
		if err != nil {
			log.Println("error fetch:", err)
			continue
		}
		log.Println("HTML length:", len(html))

		jobs := parser.ParseJobs(html)

		for _, job := range jobs {
			if !cache.Exists(job.Link) {
				notifier.SendTelegram(job)
				cache.Add(job.Link)
			}
		}

		log.Println("HTML length:", len(html))
		log.Println(html[:300])
		<-ticker.C
	}
}
