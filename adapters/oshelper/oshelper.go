package oshelper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

type OsHelper struct{}

func NewOsHelper() *OsHelper {
	return new(OsHelper)
}

func (o *OsHelper) GetOSType() string {
	return strings.ToLower(runtime.GOOS)
}

func (o *OsHelper) GetArch() (string, error) {
	arch := runtime.GOARCH
	switch arch {
	case "amd64", "x86_64":
		return "amd64", nil
	case "arm64", "aarch64":
		return "arm64", nil
	default:
		return "", fmt.Errorf("unsupported architecture: %s", arch)
	}
}

func (o *OsHelper) DownloadBinary(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	err = os.Chmod(filepath, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (o *OsHelper) MakeDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
