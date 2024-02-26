package models

type Job struct {
	ID     string `json:"_id"`
	Type   string `json:"type"`
	Prompt string `json:"prompt"`
	Seed   int64  `json:"seed,omitempty"`
	Model  string `json:"model"`
}

type GalleryJob struct {
	ImageData string `json:"imageData"`
	ID        string `json:"_id"`
	Type      string `json:"type"`
	Prompt    string `json:"prompt"`
	Seed      string `json:"seed"`
	Status    string `json:"status"`
}
