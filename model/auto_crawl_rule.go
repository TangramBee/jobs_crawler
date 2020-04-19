package model

import "time"

const (
	AutoCrawlOn = iota
	AutoCrawOff
)

type AutoCrawlRule struct {
	ID             int
	Domain         string
	AllURL         string
	IncrURL        string
	Keywords       string
	ListSelector   string
	ResultSelector string
	PageField      string
	MaxPage        int
	Ext            string
	State          int
	CreatedAt      time.Time
}
