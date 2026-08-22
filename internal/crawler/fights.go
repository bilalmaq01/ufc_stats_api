package crawler

import (
	"fmt"
	"log"
	"strconv"
	"ufc_stats_api/internal/models"
	"ufc_stats_api/internal/storage"

	"github.com/gocolly/colly/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func FightCrawler(c *colly.Collector, pool *pgxpool.Pool) {
	c.OnHTML("tr.js-fight-details-click", func(f *colly.HTMLElement) {
		var fight models.Fight
		var err error
		fight.URL = f.Attr("data-link")
		fighterURL := f.ChildAttrs("td:nth-child(2) a", "href")
		if len(fighterURL) != 2 {
			return
		}
		fight.Fighter1ID, err = storage.GetFighterIDByURL(pool, fighterURL[0])
		if err != nil {
			return
		}
		fight.Fighter2ID, err = storage.GetFighterIDByURL(pool, fighterURL[1])
		if err != nil {
			return
		}
		fight.Method = f.ChildText("td:nth-child(8)")
		fight.Time = f.ChildText("td:nth-child(10)")
		fight.WeightClass = f.ChildText("td:nth-child(7)")
		fight.Round, err = strconv.Atoi(f.ChildText("td:nth-child(9)"))
		if err != nil {
			return
		}
		eventURL := f.Request.URL.String()
		fight.EventID, err = storage.GetEventIDByURL(pool, eventURL)
		if err != nil {
			return
		}
		if f.ChildText("td:nth-child(1) i.b-flag__text") == "win" {
			fight.WinnerID = &fight.Fighter1ID
		}
		fight.IsTitle = f.DOM.Find("td:nth-child(7) img[src='http://1e49bc5171d173577ecd-1323f4090557a33db01577564f60846c.r80.cf1.rackcdn.com/belt.png']").Length() > 0
		if err := storage.InsertFight(pool, &fight); err != nil {
			log.Println(err)
			return
		}
	})
	events, err := storage.GetAllEvents(pool)
	if err != nil {
		return
	}

	for _, event := range events {
		c.Visit(event.URL)
		fmt.Println("visited", event.URL)
	}
}
