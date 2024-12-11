package db

import d "github.com/heisenburgger/pdty-app/domain"

func (r *Repo) CreateProject(project d.Project) error {
	query := `
	  INSERT INTO projects(name, description) VALUES(?, ?)
	`
	_, err := r.db.Exec(query, project.Name, project.Description)
	return err
}

func (r *Repo) GetProjects() ([]*d.Project, error) {
	query := `
		SELECT id, name, description FROM projects
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]*d.Project, 0)
	for rows.Next() {
		var project d.Project
		err := rows.Scan(&project.ID, &project.Name, &project.Description)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
