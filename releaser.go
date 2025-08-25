package releaser

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/releaser/adapters/prompter"
	"github.com/goodylabs/releaser/ports"
	"github.com/goodylabs/releaser/providers/github"
	"github.com/goodylabs/releaser/release"
	"github.com/goodylabs/releaser/utils"
)

type ReleaserInstance struct {
	release  *release.ReleaseCfg
	provider ports.Provider
	appDir   string
	prompter *prompter.Prompter
}

func (e *ReleaserInstance) Run() (bool, error) {
	configPath := filepath.Join(e.appDir, "config.json")
	if !e.release.CheckNeedsCheck(configPath) {
		return false, nil
	}

	fmt.Println("Checking for updates...")

	newestRelease, err := e.provider.GetNewestReleaseName()
	if err != nil {
		return false, err
	}

	confirmMsg := fmt.Sprintf("New version %s is available. Do you want to update? ([y]/n)", newestRelease)
	confirm, err := e.prompter.Confirm(confirmMsg)
	if err != nil {
		return false, err
	}

	e.release.ReleaseName = newestRelease
	e.release.LastCheck = utils.GetCurrentDate()
	if err := e.release.WriteReleaseCfg(configPath, e.release); err != nil {
		return false, err
	}

	if !confirm {
		return false, nil
	}

	if err := e.provider.PerformUpdate(e.appDir); err != nil {
		return false, err
	}
	return true, nil
}

func (e *ReleaserInstance) ForceUpdate() error {
	configPath := filepath.Join(e.appDir, "config.json")

	fmt.Println("Checking for updates...")

	newestRelease, err := e.provider.GetNewestReleaseName()
	if err != nil {
		return err
	}

	e.release.ReleaseName = newestRelease
	e.release.LastCheck = utils.GetCurrentDate()
	if err := e.release.WriteReleaseCfg(configPath, e.release); err != nil {
		return err
	}

	return e.provider.PerformUpdate(e.appDir)
}

func ConfigureGithubApp(appDir string, opts *github.GithubOpts) *ReleaserInstance {
	return &ReleaserInstance{
		appDir:   appDir,
		release:  release.NewReleaseCfg(),
		provider: github.NewGithubApp(opts),
		prompter: prompter.NewPrompter(),
	}
}
