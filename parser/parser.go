package parser

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Job struct {
	Title string
	Link  string
}

func ParseJobs(html string) []Job {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	var jobs []Job

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h2").Text()
		link, exists := s.Attr("data-href")

		if title == "" || !exists {
			return
		}

		jobs = append(jobs, Job{
			Title: strings.TrimSpace(title),
			Link:  "https://remoteok.com" + link,
		})
	})
	for _, job := range jobs {
		log.Println("DEBUG:", job.Title, job.Link)
	}

	log.Println("Jobs found:", len(jobs))

	return jobs
}

func MatchKeyword(job Job, keyword string) bool {
	return strings.Contains(strings.ToLower(job.Title), strings.ToLower(keyword))
}
