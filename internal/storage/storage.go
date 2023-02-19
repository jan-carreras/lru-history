// Package storage reads/writes the history file
package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gofrs/flock"
	"github.com/jan-carreras/lru-history/internal/models"
)

const (
	writeFileOptions = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	readFileOptions  = os.O_RDONLY
)

// Storage abstract the access to the History file safely
type Storage struct {
	historyPath string
	mux         *flock.Flock
}

// NewStorage returns an Storage
func NewStorage(historyPath string) *Storage {
	return &Storage{
		historyPath: historyPath,
		mux:         flock.New(fmt.Sprintf("%s.lock", historyPath)),
	}
}

// AddHistoryLine add an entry to the History file
func (s *Storage) AddHistoryLine(input io.Reader) error {
	if err := s.mux.Lock(); err != nil {
		return err
	}
	defer func() { _ = s.mux.Unlock() }()

	f, err := os.OpenFile(s.historyPath, writeFileOptions, 0600)
	if err != nil {
		return fmt.Errorf("unable to open history file: %w", err)
	}
	defer func() { _ = f.Close() }()

	buf := bufio.NewWriter(f)
	if _, err := io.Copy(buf, input); err != nil {
		return err
	}

	return buf.Flush()
}

// ReadHistory read the History file and return all the entries
func (s *Storage) ReadHistory() ([]models.HistoryLine, error) {
	if err := s.mux.Lock(); err != nil {
		return nil, err
	}
	defer func() { _ = s.mux.Unlock() }()

	f, err := os.OpenFile(s.historyPath, readFileOptions, 0600)
	if err != nil {
		return nil, fmt.Errorf("unable to open history file: %w", err)
	}
	defer func() { _ = f.Close() }()

	historyLines := make([]models.HistoryLine, 0)

	var hl historyLine
	d := json.NewDecoder(f)
	for d.More() {
		if err := d.Decode(&hl); err != nil {
			return nil, fmt.Errorf("unable to decode at offset: %d: %w", d.InputOffset(), err)
		}
		historyLines = append(historyLines, models.HistoryLine{
			Directory:  hl.Directory,
			Command:    hl.Command,
			ExecutedAt: time.Unix(int64(hl.CreatedAt), 0),
		})
	}

	return historyLines, nil
}
