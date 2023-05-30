package main

import (
	"fmt"

	"golang.org/x/oauth2"
)

func main() {
	// Create an OAuth2 client.
	client := oauth2.NewClient(oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	})

	// Get the authorization URL.
	authURL := client.AuthCodeURL("YOUR_REDIRECT_URI")

	// Redirect the user to the authorization URL.
	fmt.Println("Visit the following URL in your browser:", authURL)

	// Wait for the user to grant authorization.
	fmt.Println("Waiting for the user to grant authorization...")
	var code string
	fmt.Scanf("%s", &code)

	// Exchange the authorization code for an access token.
	token, err := client.Exchange(oauth2.NoContext, code)
	if err != nil {
		panic(err)
	}

	// Use the access token to make requests to the Google API.
	fmt.Println("Access token:", token.AccessToken)
}
