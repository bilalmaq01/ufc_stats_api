# UFC Stats API

A Go-based scraper and REST API for historical UFC fight, fighter, and event data. A Colly-based crawler pulls raw data and normalizes it into PostgreSQL, and a separate HTTP service exposes it for querying.

## How it works

The project is split into two binaries that share a common data layer:

- **`cmd/crawler`** — crawls UFC event and fighter pages with [Colly](https://github.com/gocolly/colly), parses the HTML, and writes normalized records into PostgreSQL.
- **`cmd/api`** — a lightweight HTTP server (stdlib `net/http`) that reads from the same database and serves fighter data as JSON.

Both connect to Postgres through [pgx](https://github.com/jackc/pgx), using connection pooling via `pgxpool`.

## Data model

Data is normalized across three core entities:

- **Fighters** — name, physical stats (height, reach, stance, DOB), and career striking/grappling averages (SLpM, strike accuracy/defense, takedown average/accuracy/defense, submission average)
- **Events** — event name, date, location
- **Fights** — links two fighters to an event, with result, method, round, time, and title-fight flag
- **Fight Stats** — per-round breakdown of strikes and takedowns by target (head/body/leg) and position (distance/clinch/ground)

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/fighters` | Returns all fighters in the database |
| `GET` | `/fighters/search?name=<name>` | Looks up a single fighter by name |

Responses are JSON. A search with no match returns `404`; a missing `name` parameter returns `400`.

## Getting started

### Prerequisites

- Go 1.26+
- PostgreSQL

### Setup

1. Clone the repo and copy the env template:
   ```bash
   git clone https://github.com/bilalmaq01/ufc_stats_api.git
   cd ufc_stats_api
   cp .env.example .env
   ```
2. Set `DATABASE_URL` in `.env` to your Postgres connection string. `PORT` defaults to `8080` if unset.
3. Run the crawler to populate the database:
   ```bash
   go run ./cmd/crawler
   ```
4. Start the API server:
   ```bash
   go run ./cmd/api
   ```
5. Query it:
   ```bash
   curl http://localhost:8080/fighters
   curl "http://localhost:8080/fighters/search?name=Jon+Jones"
   ```

## Tech stack

Go · PostgreSQL · [pgx](https://github.com/jackc/pgx) · [Colly](https://github.com/gocolly/colly) · [godotenv](https://github.com/joho/godotenv)

## Status

Actively in development. Current focus is on filling out fight and fight-stats endpoints alongside the existing fighter endpoints.
