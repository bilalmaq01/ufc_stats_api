package main

import (
	"context"
	"log"
	"ufc_stats_api/internal/config"
	"ufc_stats_api/internal/crawler"

	"github.com/gocolly/colly/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"))
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Cookie", cfg.Cookie)
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})
	crawler.FighterCrawler(c, pool)
}
