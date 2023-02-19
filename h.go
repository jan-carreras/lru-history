package main

import (
	"flag"
	"github.com/jan-carreras/lru-history/internal"
	runner "github.com/jan-carreras/lru-history/internal/run"
	"github.com/jan-carreras/lru-history/internal/storage"
	"github.com/jan-carreras/lru-history/internal/view"
	"log"
	"os"
	"path"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	var add bool

	flag.BoolVar(&add, "add", false, "Add a command executed to history")
	flag.Parse()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dir := path.Join(homeDir, ".h_history")
	s := storage.NewStorage(dir)
	app := internal.NewApp(s, view.NewRenderer(), runner.NewRunner())

	if add {
		return app.AddToHistory(os.Stdin)
	}

	if err := app.History(os.Getenv("PWD")); err != nil {
		return err
	}

	return nil
}
