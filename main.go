package main

import (
	"log"
	"net/http"
	"ufc_stats_api/crawler"
	"ufc_stats_api/handlers"

	"github.com/gocolly/colly/v2"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/fighters", handlers.GetAllFighters)
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"))
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Cookie", "_fmc=1783440808.984b83b12b4ac36dde01d6f661552ec215ca401052a4c91d2405a16f18a799cd")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})
	crawler.FighterCrawler(c)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
