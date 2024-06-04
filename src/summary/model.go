package summary

import (
	"time"

	"github.com/google/uuid"
)

type Summary struct {
	ID         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"user_id"`
	Summary    string    `json:"summary"`
	TotalJob   string    `json:"total_job"`
	CurrentJob string    `json:"current_job"`
	Status     string    `json:"status"`
	VideoUrl   string    `json:"video_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
