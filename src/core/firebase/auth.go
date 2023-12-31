package firebase

import (
	"FM/src/core/exception"
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type FirebaseAuth struct {
	auth.Client
}

func NewFirebaseAuth(firebase *firebase.App) FirebaseAuth {
	client, err := firebase.Auth(context.Background())
	if err != nil {
		exception.PanicLogging(err)
	}
	return FirebaseAuth{Client: *client}
}

type UserPayload struct {
	Email   string `json:"email"`
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func (auth FirebaseAuth) VerifyIDToken(ctx context.Context, idToken string) (UserPayload, error) {
	token, err := auth.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return UserPayload{}, err
	}

	claims := token.Claims

	userID := claims["user_id"].(string)
	name := claims["name"].(string)
	email := claims["email"].(string)
	picture := claims["picture"].(string)

	return UserPayload{
		UserID:  userID,
		Name:    name,
		Email:   email,
		Picture: picture,
	}, nil
}

