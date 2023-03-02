package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sagarhande/roomifyr/pkg/config"
	"github.com/sagarhande/roomifyr/pkg/models"
	"github.com/sagarhande/roomifyr/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Rooms is the handler for the different types of rooms
func (m *Repository) Rooms(w http.ResponseWriter, r *http.Request) {

	roomType := chi.URLParam(r, "roomType")

	switch roomType {
	case "generals-quarters":
		fmt.Println("rendering.. generals-quarters.page.html")
		render.RenderTemplate(w, "generals-quarters.page.html", &models.TemplateData{})
	case "majors-suite":
		fmt.Println("rendering.. majors-suite.page.html")
		render.RenderTemplate(w, "majors-suite.page.html", &models.TemplateData{})
	default:
		fmt.Println("rendering.. rooms.page.html")
		render.RenderTemplate(w, "rooms.page.html", &models.TemplateData{})
	}
}

// Reservation is the handler for the reservation
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "reservation.page.html", &models.TemplateData{})
}

// Contact is the handler for the contact
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.html", &models.TemplateData{})
}
