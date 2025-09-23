package releaser

import (
	"github.com/goodylabs/releaser/utils"
)

type ReleaseCfg struct {
	ReleaseName string `json:"releaseName"`
	LastCheck   string `json:"lastCheck"`
}

func (b *ReleaseCfg) DontNeedCheck(path string) bool {
	cfg, err := utils.ReadJSONFromFile[ReleaseCfg](path)
	if err != nil {
		return true
	}
	return cfg.LastCheck != utils.GetCurrentDate()
}

func (b *ReleaseCfg) WriteReleaseCfg(path string, cfg *ReleaseCfg) error {
	return utils.WriteJSONToFile(path, cfg)
}
