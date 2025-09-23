package api

import (
	"github.com/goodylabs/releaser/adapters/prompter"
	"github.com/goodylabs/releaser/providers/github"
	"github.com/goodylabs/releaser/releaser"
)

type GithubAppOpts struct {
	User string
	Repo string
}

func ConfigureGithubApp(appDir string, opts *GithubAppOpts) *releaser.ReleaserInstance {
	return releaser.NewReleaserInstance(
		appDir,
		github.NewGithubApp(&github.GithubOpts{
			User: opts.User,
			Repo: opts.Repo,
		}),
		prompter.NewPrompter(),
	)
}
