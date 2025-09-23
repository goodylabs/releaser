package releaser_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/releaser/releaser"
	"github.com/goodylabs/releaser/tests/testutils"
	"github.com/stretchr/testify/assert"
)

func TestCfgLoadFromFile(t *testing.T) {
	testRscDir := testutils.GetTestRscDir()
	cfgPath := filepath.Join(testRscDir, "ok.json")

	var cfg releaser.ReleaseCfg
	err := cfg.LoadFromFile(cfgPath)
	assert.NoError(t, err)
	assert.Equal(t, "v1.0.0", cfg.ReleaseName)
	assert.Equal(t, "2023-10-01", cfg.LastCheck)
}
