package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/goodylabs/releaser/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	githubUser := os.Getenv("GITHUB_USER")
	githubRepo := os.Getenv("GITHUB_REPO")

	devPath := filepath.Join(".development")

	app := api.ConfigureGithubApp(
		devPath,
		&api.GithubAppOpts{
			User: githubUser,
			Repo: githubRepo,
		})

	if _, err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
