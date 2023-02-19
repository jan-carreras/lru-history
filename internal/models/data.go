package models

import "time"

type HistoryLine struct {
	ExecutedAt time.Time
	Directory  string
	Command    string
}

type Counter struct {
	Count   int
	Command HistoryLine
}
