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

	if err := releaseCfg.LoadFromFile(configPath); err != nil {
		fmt.Print("Could not read release config - trying to update anyway, error: ", err)
	}

	if releaseCfg.CheckDontNeedCheck() {
		return false, nil
	}

	fmt.Println("Checking for updates...")

	newestRelease, err := e.provider.GetNewestReleaseName()
	if err != nil {
		return false, err
	}

	if releaseCfg.CheckIsReleaseTheNewest(newestRelease) {
		return false, nil
	}

	confirmMsg := fmt.Sprintf("New version %s is available. Do you want to update?", newestRelease)
	confirm, err := e.prompter.Confirm(confirmMsg)
	if err != nil {
		return false, err
	}

	releaseCfg.LastCheck = utils.GetCurrentDate()
	if err := releaseCfg.SaveToFile(configPath); err != nil {
		return false, err
	}

	if !confirm {
		return false, nil
	}

	if err := e.provider.PerformUpdate(e.appDir); err != nil {
		return false, err
	}

	releaseCfg.ReleaseName = newestRelease
	if err := releaseCfg.SaveToFile(configPath); err != nil {
		return false, err
	}

	return true, nil
}

func (e *ReleaserInstance) ForceUpdate() error {
	var releaseCfg ReleaseCfg

	configPath := e.getConfigPath()

	releaseCfg.LoadFromFile(configPath)

	fmt.Println("Checking for updates...")

	newestRelease, err := e.provider.GetNewestReleaseName()
	if err != nil {
		return err
	}

	releaseCfg.ReleaseName = newestRelease
	releaseCfg.LastCheck = utils.GetCurrentDate()
	if err := releaseCfg.SaveToFile(configPath); err != nil {
		return err
	}

	return e.provider.PerformUpdate(e.appDir)
}
