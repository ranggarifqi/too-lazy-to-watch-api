package auth

import (
	"bytes"
	"encoding/binary"
	"too-lazy-to-watch-api/src/user"

	"github.com/google/uuid"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

type supabaseAuthRepository struct {
	client *supabase.Client
}

func NewSupabaseAuthRepository() IAuthRepository {
	return &supabaseAuthRepository{}
}

func (r *supabaseAuthRepository) SignUpByEmail(payload ISignupPayload) (*user.User, error) {
	res, err := r.client.Auth.Signup(types.SignupRequest{
		Email:    payload.Email,
		Password: payload.Password,
	})
	if err != nil {
		return nil, err
	}

	userPayload := struct {
		id    uuid.UUID
		email string
		name  string
	}{res.ID, payload.Email, payload.Name}

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
