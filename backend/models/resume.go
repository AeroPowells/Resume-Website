package models

// Resume is the top-level data model for the resume API.
type Resume struct {
	Bio        Bio          `json:"bio"`
	Experience []Experience `json:"experience"`
	Education  []Education  `json:"education"`
	Skills     []SkillGroup `json:"skills"`
}

type Bio struct {
	Name     string   `json:"name"`
	Title    string   `json:"title"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Location string   `json:"location"`
	Summary  string   `json:"summary"`
	Links    []Link   `json:"links"`
}

type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Experience struct {
	Company     string   `json:"company"`
	Role        string   `json:"role"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Location    string   `json:"location"`
	Highlights  []string `json:"highlights"`
}

type Education struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Field       string `json:"field"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type SkillGroup struct {
	Category string   `json:"category"`
	Skills   []string `json:"skills"`
}
