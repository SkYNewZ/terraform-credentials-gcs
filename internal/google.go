package internal

import (
	"context"

	"golang.org/x/oauth2/google"
)

// GetGoogleAccessToken return a default Google access token
func GetGoogleAccessToken(ctx context.Context) (string, error) {
	credentials, err := google.FindDefaultCredentials(ctx, "")
	if err != nil {
		return "", err
	}

	token, err := credentials.TokenSource.Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
