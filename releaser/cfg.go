package releaser

import (
	"github.com/goodylabs/releaser/utils"
)

type ReleaseCfg struct {
	ReleaseName string `json:"release"`
	LastCheck   string `json:"lastCheck"`
}

func (r *ReleaseCfg) LoadFromFile(path string) error {
	cfg, err := utils.ReadJSONFromFile[ReleaseCfg](path)
	if err != nil {
		return err
	}
	*r = cfg
	return nil
}

func (r *ReleaseCfg) CheckWasTodayVersionChecked() bool {
	return r.LastCheck == utils.GetCurrentDate()
}

func (r *ReleaseCfg) SaveToFile(path string) error {
	return utils.WriteJSONToFile(path, &r)
}

func (r *ReleaseCfg) CheckIsReleaseTheNewest(newest string) bool {
	return r.ReleaseName == newest
}

func (r *ReleaseCfg) UpdateLastCheckDate() {
	r.LastCheck = utils.GetCurrentDate()
}
