package model

type Resume struct {
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	PhoneNumber   int64      `json:"phone_number"`
	LinkedinURL   string     `json:"linkedin_url"`
	PortofolioURL string     `json:"portofolio_url"`
	Ocupations    Occupation `json:"ocupations"`
	Educations    Education  `json:"educations"`
	Achievements  []string   `json:"achievements"`
}
