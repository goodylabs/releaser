package github

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/releaser/adapters/httpconnector"
	"github.com/goodylabs/releaser/adapters/oshelper"
	"github.com/goodylabs/releaser/ports"
)

type GithubOpts struct {
	User string
	Repo string
}

type githubApp struct {
	opts           GithubOpts
	newReleaseName string
	newReleaseUrl  string
	httpconnector  *httpconnector.HttpClient
	oshelper       *oshelper.OsHelper
}

func NewGithubApp(opts *GithubOpts) ports.Provider {
	return &githubApp{
		opts:          *opts,
		httpconnector: httpconnector.NewHttpClient(),
	}
}

func (g *githubApp) GetNewestReleaseName() (string, error) {
	lastestReleaseUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", g.opts.User, g.opts.Repo)
	if err := g.httpconnector.DoGet(lastestReleaseUrl, &releaseRes); err != nil {
		return "", err
	}

	fmt.Println("Latest Release URL:", lastestReleaseUrl)

	g.newReleaseName = releaseRes.TagName

	osType := g.oshelper.GetOSType()
	osArch, err := g.oshelper.GetArch()
	if err != nil {
		return "", err
	}

	assetName := fmt.Sprintf("tug-%s-%s", osType, osArch)
	for _, asset := range releaseRes.Assets {
		if asset.Name == assetName {
			g.newReleaseUrl = asset.BrowserDownloadURL
			break
		}
	}

	return g.newReleaseName, nil
}

func (g *githubApp) PerformUpdate(appDir string) error {
	osType := g.oshelper.GetOSType()
	osArch, err := g.oshelper.GetArch()
	if err != nil {
		return err
	}

	if g.newReleaseUrl == "" {
		return fmt.Errorf("no compatible binary found for %s-%s", osType, osArch)
	}

	fmt.Println("Downloading binary from:", g.newReleaseUrl)

	binnaryDir := filepath.Join(appDir, "bin")
	if err := g.oshelper.MakeDirIfNotExist(binnaryDir); err != nil {
		return err
	}

	binnaryPath := filepath.Join(binnaryDir, g.opts.Repo)
	if err := g.oshelper.DownloadBinary(g.newReleaseUrl, binnaryPath); err != nil {
		return err
	}

	fmt.Println("Updated binary at:", binnaryPath)

	return nil
}
