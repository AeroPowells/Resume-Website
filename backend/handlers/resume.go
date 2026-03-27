package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dsluss/resume-website/backend/data"
	"github.com/dsluss/resume-website/backend/models"
)

var resume models.Resume

func init() {
	resume = data.SeedResume()
}

// writeJSON is a helper that sets the Content-Type header and encodes v as JSON.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// Health handles GET /api/health
func Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// GetResume handles GET /api/resume
func GetResume(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, resume)
}

// GetBio handles GET /api/resume/bio
func GetBio(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, resume.Bio)
}

// GetExperience handles GET /api/resume/experience
func GetExperience(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, resume.Experience)
}

// GetEducation handles GET /api/resume/education
func GetEducation(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, resume.Education)
}

// GetSkills handles GET /api/resume/skills
func GetSkills(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, resume.Skills)
}
