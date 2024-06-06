package summary

import (
	"encoding/json"

	"github.com/supabase-community/supabase-go"
)

type supabaseSummaryRepository struct {
	client *supabase.Client
}

func NewSupabaseSummaryRepository(client *supabase.Client) ISummaryRepository {
	return &supabaseSummaryRepository{
		client: client,
	}
}

func (s *supabaseSummaryRepository) Create(payload CreateSummaryPayload) (*Summary, error) {
	res, _, err := s.client.From(TABLE_NAME).Insert(payload, false, "", "", "").Execute()
	if err != nil {
		return nil, err
	}

	summary := []Summary{}
	err = json.Unmarshal(res, &summary)
	if err != nil {
		return nil, err
	}

	return &summary[0], nil
}
