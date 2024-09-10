package config

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/joho/godotenv"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type FirebaseServiceAccount struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

func InitializeFirebase() *firebase.App {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	firebaseConfigJSON := os.Getenv("FIREBASE_JSON")
	if firebaseConfigJSON == "" {
		log.Fatal("FIREBASE_CONFIG not found in .env")
	}
	// Parse JSON into struct
	var config FirebaseServiceAccount
	err = json.Unmarshal([]byte(firebaseConfigJSON), &config)
	if err != nil {
		log.Fatal("Error parsing FIREBASE_CONFIG JSON:", err)
	}

	// Convert the struct back to JSON
	credJSON, err := json.Marshal(config)
	if err != nil {
		log.Fatal("Error converting config to JSON:", err)
	}

	ctx := context.Background()

	opt := option.WithCredentialsJSON(credJSON)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	return app
}

func FirebaseAuthInitialize(app *firebase.App) *auth.Client {
	ctx := context.Background()
	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}
	return firebaseAuth
}
