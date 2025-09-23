package ports

type Provider interface {
	GetNewestReleaseName() (string, error)
	PerformUpdate(appDir string) error
}
