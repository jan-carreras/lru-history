package internal

import (
	"github.com/jan-carreras/lru-history/internal/models"
	"github.com/jan-carreras/lru-history/internal/storage"
	"github.com/jan-carreras/lru-history/internal/view"
	"os"
	"path"
	"strings"
	"testing"
)

func TestApp_AddToHistory(t *testing.T) {
	storageDir := path.Join(t.TempDir(), "history_file")
	app := NewApp(storage.NewStorage(storageDir), makeTestRenderer(), makeTestRunner())

	input := []string{
		`{"created_at":1676675709,"pwd":"/Users/jan/dev/h","command":"make"}` + "\n",
		`{"created_at":1676675709,"pwd":"/Users/jan/dev/h","command":"make install"}` + "\n",
	}

	for _, i := range input {
		if err := app.AddToHistory(strings.NewReader(i)); err != nil {
			t.Errorf("no error expected: %q", err)
		}
	}
	raw, err := os.ReadFile(storageDir)
	if err != nil {
		t.Errorf("no error expected: %q", err)
	}

	want := strings.Join(input, "")
	if want != string(raw) {
		t.Errorf("unexpected content of history file:\n%s\nwant:\n%s", string(raw), want)
	}
}

func addToHistory(t *testing.T, commands []string, app *App) {
	t.Run("add to history file", func(t *testing.T) {
		for _, i := range commands {
			if err := app.AddToHistory(strings.NewReader(i)); err != nil {
				t.Errorf("no error expected: %q", err)
			}
		}
	})
}

func TestReadHistory(t *testing.T) {
	storageDir := path.Join(t.TempDir(), "history_file")
	app := NewApp(storage.NewStorage(storageDir), makeTestRenderer(), makeTestRunner())

	input := []string{
		`{"created_at":1676675709,"pwd":"/Users/jan/dev/h","command":"make"}`,
		`{"created_at":1676675709,"pwd":"/Users/jan/dev/h","command":"make install"}`,
	}

	addToHistory(t, input, app)

	if err := app.History(storageDir); err != nil {
		t.Errorf("no error expected: %q", err)
	}
}

type testRenderer struct{}

func (tr *testRenderer) Render(cmds []models.Counter) (*models.HistoryLine, error) {
	return nil, view.ErrNoCommandSelected
}

func makeTestRenderer() *testRenderer {
	return &testRenderer{}
}

type testRunner struct{}

func (r *testRunner) Run(command models.HistoryLine) error {
	return nil
}

func makeTestRunner() *testRunner {
	return &testRunner{}
}
