package main

import (
	"log"
	"os"
	"path/filepath"

	releaser "github.com/goodylabs/releaser"
	"github.com/goodylabs/releaser/providers/github"
	"github.com/joho/godotenv"
)

func main() {
	devPath := filepath.Join(".development")

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}

	githubUser := os.Getenv("GITHUB_USER")
	githubRepo := os.Getenv("GITHUB_REPO")

	app := releaser.ConfigureGithubApp(
		devPath,
		&github.GithubOpts{
			User: githubUser,
			Repo: githubRepo,
		})

	if _, err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
