package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/heisenburgger/pdty-app/db"
)

type application struct {
	repo   *db.Repo
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	repo, err := db.New()
	if err != nil {
		log.Fatalf("failed to connect to db: %+v", err)
	}
	logger.Info("init", "message", "connected to db 💾")

	app := application{
		repo:   repo,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/projects", app.createProjectHandler)
	mux.HandleFunc("GET /api/v1/projects", app.getProjectsHandler)
	mux.HandleFunc("POST /api/v1/tasks", app.createTaskHandler)
	mux.HandleFunc("GET /api/v1/tasks", app.getTasksHandler)

	logger.Info("init", "message", "api ready to go 🚀")
	if err := http.ListenAndServe(":6969", mux); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
