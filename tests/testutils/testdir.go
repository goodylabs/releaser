package testutils

import (
	"os"
	"path/filepath"
)

func GetTestRscDir() string {
	return filepath.Join(GetRootDir(), ".testresources")
}

func GetRootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic("could not get working directory")
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			panic("could not find project root")
		}
		dir = parent
	}
}
