package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"ufc_stats_api/internal/models"

	"github.com/gocolly/colly/v2"
)

var errSkipRow = errors.New("skip row")

func FighterCrawler(c *colly.Collector) {

	c.OnHTML("tr.b-statistics__table-row", func(e *colly.HTMLElement) {

		fighter, err := parseFighterRow(e)
		if err != nil {
			if err == errSkipRow {
				return
			}
			log.Println(err)
			return
		}
		b, err := json.MarshalIndent(fighter, "", "  ")
		if err != nil {
			log.Println(err)
			return
		}

		b = b
	})

	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		link := fmt.Sprintf("http://www.ufcstats.com/statistics/fighters?char=%c&page=all", letter)
		c.Visit(link)
		link = fmt.Sprintf("Visited %s", link)
		fmt.Println(link)
	}

}

func parseFighterRow(e *colly.HTMLElement) (models.Fighter, error) {
	var fighter models.Fighter
	var err error
	firstName := strings.TrimSpace(e.ChildText("td:nth-child(1)"))
	lastName := strings.TrimSpace(e.ChildText("td:nth-child(2)"))

	if firstName == "" || lastName == "" {
		return models.Fighter{}, errSkipRow
	}
	fighter.Name = fmt.Sprintf("%s %s", firstName, lastName)

	n := e.ChildText("td:nth-child(3)")
	if n != "" {
		fighter.Nickname = &n
	}

	h := e.ChildText("td:nth-child(4)")
	if h != "--" {
		h = strings.TrimSpace(h)
		h = strings.TrimSuffix(h, "\"")
		fighter.Height = &h
	}

	w := e.ChildText("td:nth-child(5)")
	if w != "--" {
		fighter.WeightClass = &w
	}
	r := e.ChildText("td:nth-child(6)")
	r = strings.TrimSpace(r)
	r = strings.TrimSuffix(r, "\"")
	if r != "--" {
		ReachIn, err := strconv.ParseFloat(r, 16)
		if err != nil {
			return models.Fighter{}, err
		}
		fighter.ReachIn = int(ReachIn)

	}
	Wins := e.ChildText("td:nth-child(8)")
	fighter.Wins, err = strconv.Atoi(Wins)
	if err != nil {
		return models.Fighter{}, err
	}
	Losses := e.ChildText("td:nth-child(9)")
	fighter.Losses, err = strconv.Atoi(Losses)
	if err != nil {
		return models.Fighter{}, err
	}
	Draws := e.ChildText("td:nth-child(10)")
	fighter.Draws, err = strconv.Atoi(Draws)
	if err != nil {
		return models.Fighter{}, err
	}

	return fighter, nil
}
