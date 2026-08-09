package crawler

import (
	"log"
	"strings"
	"time"
	"ufc_stats_api/internal/models"
	"ufc_stats_api/internal/storage"

	"github.com/gocolly/colly/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func EventCrawler(c *colly.Collector, pool *pgxpool.Pool) {
	c.OnHTML("tr.b-statistics__table-row a.b-link", func(e *colly.HTMLElement) {
		var event models.Event
		event.EventName = strings.TrimSpace(e.Text)
		eventdate := strings.TrimSpace(e.DOM.Parent().Find("span.b-statistics__date").Text())
		date, err := time.Parse("January 02, 2006", eventdate)
		if err != nil {
			log.Println("Error parsing date from", e.Text)
			return
		}
		event.Date = date
		location := strings.TrimSpace(e.DOM.Closest("tr").Find("td").Eq(1).Text())
		event.Location = location
		event.URL = e.Attr("href")
		if err := storage.InsertEvent(pool, &event); err != nil {
			log.Println(err)
			return
		}
	})
	c.Visit("http://ufcstats.com/statistics/events/completed?page=all")
}
