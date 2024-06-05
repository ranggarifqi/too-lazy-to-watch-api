package auth

import (
	"encoding/json"
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
		Id    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}{res.ID.String(), payload.Email, payload.Name}

	data, _, err := r.client.From(user.TABLE_NAME).Insert(userPayload, true, "id", "", "").Execute()
	if err != nil {
		r.adminDeleteUser(res.ID)

		return nil, err
	}

	users := []user.User{}

	err = json.Unmarshal(data, &users)
	if err != nil {
		r.adminDeleteUser(res.ID)
		r.deleteUserById(res.ID)

		return nil, err
	}

	return &users[0], nil
}

func (r *supabaseAuthRepository) LoginByPassword(email string, password string) (string, error) {
	res, err := r.client.SignInWithEmailPassword(email, password)
	if err != nil {
		return "", err
	}

	return res.AccessToken, nil
}

func (r *supabaseAuthRepository) adminDeleteUser(userId uuid.UUID) {
	err := r.authClient.AdminDeleteUser(types.AdminDeleteUserRequest{
		UserID: userId,
	})
	if err != nil {
		fmt.Printf("Failed to delete user auth: %v\n", err)
	}
}

func (r *supabaseAuthRepository) deleteUserById(userId uuid.UUID) {
	_, _, err := r.client.From(user.TABLE_NAME).Delete("", "").Eq("id", userId.String()).Execute()
	if err != nil {
		fmt.Printf("Failed to delete a user: %v\n", err)
	}
}
