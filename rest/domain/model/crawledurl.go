package model

import "time"

type CrawledUrl struct {
	Id        int       `json: "id"`
	Url       string    `json: "url"`
	CreatedAt time.Time `json: "created_at"`
}
