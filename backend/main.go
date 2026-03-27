package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dsluss/resume-website/backend/handlers"
)

func main() {
	mux := http.NewServeMux()

	// CORS middleware wraps the mux so all routes allow the React dev server.
	handler := corsMiddleware(mux)

	// Routes
	mux.HandleFunc("GET /api/health", handlers.Health)
	mux.HandleFunc("GET /api/resume", handlers.GetResume)
	mux.HandleFunc("GET /api/resume/bio", handlers.GetBio)
	mux.HandleFunc("GET /api/resume/experience", handlers.GetExperience)
	mux.HandleFunc("GET /api/resume/education", handlers.GetEducation)
	mux.HandleFunc("GET /api/resume/skills", handlers.GetSkills)

	addr := ":8080"
	fmt.Printf("Resume API listening on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}

// corsMiddleware adds permissive CORS headers for local development.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
