// Package models defines basic data structures shared on the application
package models

import "time"

// HistoryLine represents a single command that has been executed
type HistoryLine struct {
	ExecutedAt time.Time
	Directory  string
	Command    string
}

// Counter represents the number of times a command has been executed
type Counter struct {
	Count   int
	Command HistoryLine
}
