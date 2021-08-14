package model

import "time"

type InMemoryRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DocumentFilters struct {
	StartDateStr string `json:"startDate"`
	EndDateStr   string `json:"endDate"`
	MinCount     int    `json:"minCount"`
	MaxCount     int    `json:"maxCount"`
	StartDate    time.Time
	EndDate      time.Time
}

type Record struct {
	Key        string    `json:"key"`
	TotalCount int       `json:"totalCount"`
	CreatedAt  time.Time `json:"createdAt"`
}
