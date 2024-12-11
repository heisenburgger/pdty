package db

import d "github.com/heisenburgger/pdty-app/domain"

func (r *Repo) CreateTask(task d.Task) error {
	query := `
	  INSERT INTO tasks(name, status, notes, scheduled, deadline, project_id) 
	  VALUES(?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, task.Name, task.Status, task.Notes, task.Scheduled, task.Deadline, task.ProjectID)
	return err
}

func (r *Repo) GetTasks() ([]*d.Task, error) {
	query := `
		SELECT id, name, status, notes, scheduled, deadline FROM tasks
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*d.Task, 0)
	for rows.Next() {
		var task d.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.Notes, &task.Scheduled, &task.Deadline)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
