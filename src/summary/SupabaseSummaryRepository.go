package summary

import (
	"encoding/json"

	"github.com/supabase-community/supabase-go"
)

type supabaseSummaryRepository struct {
	client *supabase.Client
}

func (s *supabaseSummaryRepository) DeleteById(id string) error {
	_, _, err := s.client.From(TABLE_NAME).Delete("", "").Eq("id", id).Execute()
	if err != nil {
		return err
	}

	return nil
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

func NewSupabaseSummaryRepository(client *supabase.Client) ISummaryRepository {
	return &supabaseSummaryRepository{
		client: client,
	}
}
