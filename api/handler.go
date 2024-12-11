package main

import (
	"encoding/json"
	"net/http"
	"time"

	d "github.com/heisenburgger/pdty-app/domain"
)

func (a *application) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: validate the incoming data
	var req struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		a.logger.Error("decoding body", "err", err)
		w.WriteHeader(500)
		return
	}

	err = a.repo.CreateProject(d.Project{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		a.logger.Error("creating project", "err", err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(`{"message": "ok"}`))
}

func (a *application) getProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := a.repo.GetProjects()
	if err != nil {
		a.logger.Error("get projects", "err", err)
		w.WriteHeader(500)
		return
	}

	response := struct {
		Projects []*d.Project `json:"projects"`
	}{
		Projects: projects,
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (a *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: validate the incoming data
	var req struct {
		Name      string     `json:"name"`
		Status    string     `json:"status"`
		Notes     *string    `json:"notes"`
		ProjectID int        `json:"projectID"`
		Deadline  *time.Time `json:"deadline"`
		Scheduled *time.Time `json:"scheduled"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		a.logger.Error("decoding body", "err", err)
		w.WriteHeader(500)
		return
	}

	err = a.repo.CreateTask(d.Task{
		Name:      req.Name,
		Status:    req.Status,
		Scheduled: req.Scheduled,
		Deadline:  req.Deadline,
		Notes:     req.Notes,
		ProjectID: req.ProjectID,
	})
	if err != nil {
		a.logger.Error("creating task", "err", err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(`{"message": "ok"}`))
}

func (a *application) getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := a.repo.GetTasks()
	if err != nil {
		a.logger.Error("get tasks", "err", err)
		w.WriteHeader(500)
		return
	}

	response := struct {
		Tasks []*d.Task `json:"tasks"`
	}{
		Tasks: tasks,
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
