package releaser

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/releaser/ports"
	"github.com/goodylabs/releaser/utils"
)

type ReleaserInstance struct {
	provider ports.Provider
	appDir   string
	prompter ports.Prompter
}

func NewReleaserInstance(appDir string, provider ports.Provider, prompter ports.Prompter) *ReleaserInstance {
	return &ReleaserInstance{
		appDir:   appDir,
		provider: provider,
		prompter: prompter,
	}
}

func (e *ReleaserInstance) getConfigPath() string {
	return filepath.Join(e.appDir, "config.json")
}

func (e *ReleaserInstance) Run() (bool, error) {
	var releaseCfg ReleaseCfg

	configPath := e.getConfigPath()

	if releaseCfg.DontNeedCheck(configPath) {
		return false, nil
	}

	fmt.Println("Checking for updates...")

	newestRelease, err := e.provider.GetNewestReleaseName()
	if err != nil {
		return false, err
	}

	confirmMsg := fmt.Sprintf("New version %s is available. Do you want to update?", newestRelease)
	confirm, err := e.prompter.Confirm(confirmMsg)
	if err != nil {
		return false, err
	}

	releaseCfg.ReleaseName = newestRelease
	releaseCfg.LastCheck = utils.GetCurrentDate()
	if err := releaseCfg.WriteReleaseCfg(configPath, &releaseCfg); err != nil {
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
	var releaseCfg ReleaseCfg

	configPath := e.getConfigPath()

	fmt.Println("Checking for updates...")

	newestRelease, err := e.provider.GetNewestReleaseName()
	if err != nil {
		return err
	}

	releaseCfg.ReleaseName = newestRelease
	releaseCfg.LastCheck = utils.GetCurrentDate()
	if err := releaseCfg.WriteReleaseCfg(configPath, &releaseCfg); err != nil {
		return err
	}

	return e.provider.PerformUpdate(e.appDir)
}
