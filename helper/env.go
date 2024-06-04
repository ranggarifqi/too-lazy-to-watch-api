package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnv(pathToEnvFile string) {
	err := godotenv.Load(pathToEnvFile)
	if err != nil {
		panic(err)
	}

	if os.Getenv("SUPABASE_PROJECT_URL") == "" {
		panic("SUPABASE_PROJECT_URL env must be set")
	}

	if os.Getenv("SUPABASE_API_KEY") == "" {
		panic("SUPABASE_API_KEY env must be set")
	}
}
