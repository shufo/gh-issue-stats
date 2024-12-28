package types

import (
	"fmt"
	"time"
)

// Issue represents a GitHub issue with relevant fields
type Issue struct {
	Number      int        `json:"number"`
	Title       string     `json:"title"`
	State       string     `json:"state"`
	Labels      []Label    `json:"labels"`
	PullRequest *struct{}  `json:"pull_request,omitempty"`
	CreatedAt   *time.Time `json:"created_at"`
	ClosedAt    *time.Time `json:"closed_at"`
}

// Label represents a GitHub issue label
type Label struct {
	Name string `json:"name"`
}

type DayDuration time.Duration

func (d DayDuration) String() string {
	days := time.Duration(d).Round(24 * time.Hour)
	return fmt.Sprintf("%d days", int(days.Hours()/24))
}

const UnlabeledLabel = "*unlabeled*"

// LabelStat stores statistics for a specific label
type LabelStat struct {
	Name            string  `json:"name"`
	Open            int     `json:"open"`
	Closed          int     `json:"closed"`
	Total           int     `json:"total"`
	OpenPercentage  float64 `json:"openPercentage"`
	AvgCloseTime    float64 `json:"avgCloseTime"`
	MedianCloseTime float64 `json:"medianCloseTime"`
}

// OverallStats stores the overall issue statistics
type OverallStats struct {
	Total           int     `json:"total"`
	Open            int     `json:"open"`
	Closed          int     `json:"closed"`
	OpenPercentage  float64 `json:"openPercentage"`
	AvgCloseTime    float64 `json:"avgCloseTime"`
	MedianCloseTime float64 `json:"medianCloseTime"`
}

// Statistics combines both label and overall statistics
type Statistics struct {
	LabelStats   []LabelStat  `json:"labelStats"`
	OverallStats OverallStats `json:"overallStats"`
}
