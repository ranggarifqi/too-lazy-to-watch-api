package helper

import (
	"log"
	"os"

	"github.com/supabase-community/supabase-go"
)

func GetSupabaseClient() *supabase.Client {
	client, err := supabase.NewClient(os.Getenv("SUPABASE_PROJECT_URL"), os.Getenv("SUPABASE_API_KEY"), nil)
	if err != nil {
		log.Fatalf("Failed to create Supabase client: %s", err)
	}
	return client
}
