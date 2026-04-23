package projects

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	Name    string   `json:"name"`
	Summary string   `json:"summary"`
	Stack   []string `json:"stack"`
	RepoURL string   `json:"repo_url"`
}

type ProjectsResponse struct {
	Data []Project `json:"data"`
}

func ProjectsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects := []Project{
			{Name: "Personal Portfolio", Summary: "A personal website showcasing my projects and skills.", RepoURL: "https://github.com/herrmann13/myportfolio"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(projects)
	}
}
