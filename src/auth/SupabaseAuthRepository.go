package auth

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"too-lazy-to-watch-api/src/user"

	"github.com/google/uuid"
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

type supabaseAuthRepository struct {
	client     *supabase.Client
	authClient gotrue.Client
}

func NewSupabaseAuthRepository(client *supabase.Client, authClient gotrue.Client) IAuthRepository {
	return &supabaseAuthRepository{
		client:     client,
		authClient: authClient,
	}
}

// TODO: Change supabase client to
func (r *supabaseAuthRepository) SignUpByEmail(payload ISignupPayload) (*user.User, error) {
	fmt.Println("Creating auth user")
	res, err := r.authClient.AdminCreateUser(types.AdminCreateUserRequest{
		Email:        payload.Email,
		Password:     &payload.Password,
		EmailConfirm: true,
	})
	if err != nil {
		return nil, err
	}

	userPayload := struct {
		id    uuid.UUID
		email string
		name  string
	}{res.ID, payload.Email, payload.Name}

	fmt.Println("Creating User data")
	data, _, err := r.client.From(user.TABLE_NAME).Insert(userPayload, true, "id", "", "").Execute()
	if err != nil {
		return nil, err
	}

	var structData user.User

	err = binary.Read(bytes.NewReader(data), binary.LittleEndian, &structData)
	if err != nil {
		return nil, err
	}

	return &structData, nil
}
