package model

import (
	"strconv"
	"time"
)

type JobInfo struct {
	ID        int
	Name      string
	Company   string
	City      string
	Salary    string
	Education string
	WorkExp   string
	JD        string
	Welfare   string
	URL       string
	Phone     string
	CreatedAt time.Time
}

func (j *JobInfo) String() string {
	return "id:" + strconv.Itoa(j.ID) + ";name:" + j.Name
}
