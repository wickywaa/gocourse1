package handlers

import (
	"net/http"

	"github.com/wickywaa/gocourse1/pkg/config"
	"github.com/wickywaa/gocourse1/pkg/models"
	"github.com/wickywaa/gocourse1/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

// is the repository type

type Repository struct {
	App *config.AppConfig
}

// NewRepo  creates a new repositiory
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers  sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{

		StringMap: stringMap,
	})
}
