package model

import "time"

type CrawlRule struct {
	ID           int
	Domain       string
	Name         string
	JobName      string
	JobCompany   string
	JobCity      string
	JobWorkExp   string
	JobSalary    string
	JobEducation string
	JobJD        string
	JobWelfare   string
	CreatedAt    time.Time
}