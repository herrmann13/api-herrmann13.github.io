import (
	"net/http"
	"encoding/json"
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
			{Name: "Personal Portfolio", Description: "A personal website showcasing my projects and skills.", URL: "https://myportfolio.com"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(projects)
	}