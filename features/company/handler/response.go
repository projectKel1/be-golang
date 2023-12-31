package handler

type CompanyResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Type        string `json:"type"`
	Visi        string `json:"visi"`
	Misi        string `json:"misi"`
	StartedHour string `json:"started_hour"`
	EndedHour   string `json:"ended_hour"`
}
