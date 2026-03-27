package data

import "github.com/dsluss/resume-website/backend/models"

// SeedResume returns sample resume data. Replace with your own details.
func SeedResume() models.Resume {
	return models.Resume{
		Bio: models.Bio{
			Name:     "Your Name",
			Title:    "Software Engineer",
			Email:    "you@example.com",
			Phone:    "+1 (555) 000-0000",
			Location: "City, State",
			Summary:  "Passionate software engineer with experience building scalable systems and learning new technologies.",
			Links: []models.Link{
				{Label: "GitHub", URL: "https://github.com/yourusername"},
				{Label: "LinkedIn", URL: "https://linkedin.com/in/yourusername"},
			},
		},
		Experience: []models.Experience{
			{
				Company:   "Acme Corp",
				Role:      "Senior Software Engineer",
				StartDate: "2022-01",
				EndDate:   "Present",
				Location:  "Remote",
				Highlights: []string{
					"Led migration of monolith to microservices, reducing deploy time by 40%",
					"Mentored 3 junior engineers",
					"Built internal CLI tooling in Go",
				},
			},
			{
				Company:   "Widgets Inc",
				Role:      "Software Engineer",
				StartDate: "2019-06",
				EndDate:   "2021-12",
				Location:  "New York, NY",
				Highlights: []string{
					"Developed REST APIs serving 1M+ requests/day",
					"Introduced automated testing, raising coverage from 20% to 80%",
				},
			},
		},
		Education: []models.Education{
			{
				Institution: "State University",
				Degree:      "Bachelor of Science",
				Field:       "Computer Science",
				StartDate:   "2015-09",
				EndDate:     "2019-05",
			},
		},
		Skills: []models.SkillGroup{
			{Category: "Languages", Skills: []string{"Go", "JavaScript", "TypeScript", "Python", "SQL"}},
			{Category: "Frontend", Skills: []string{"React", "HTML", "CSS", "Vite"}},
			{Category: "Backend", Skills: []string{"REST APIs", "gRPC", "PostgreSQL", "Redis"}},
			{Category: "DevOps", Skills: []string{"Docker", "Kubernetes", "GitHub Actions", "Linux"}},
		},
	}
}
