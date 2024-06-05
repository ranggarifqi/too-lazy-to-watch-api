package helper

import (
	"log"
	"os"

	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/supabase-go"
)

func GetSupabaseClient() (*supabase.Client, gotrue.Client) {
	client, err := supabase.NewClient(os.Getenv("SUPABASE_PROJECT_URL"), os.Getenv("SUPABASE_API_KEY"), nil)
	if err != nil {
		log.Fatalf("Failed to create Supabase client: %s", err)
	}
	adminClient := client.Auth.WithToken(os.Getenv("SUPABASE_API_KEY"))
	return client, adminClient
}
