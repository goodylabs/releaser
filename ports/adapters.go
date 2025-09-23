package ports

type HttpClient interface {
	DoGet(url string, out any) error
}

type OsHelper interface {
	GetOSType() string
	GetArch() (string, error)
	DownloadBinary(url, filepath string) error
	MakeDirIfNotExist(dir string) error
}

type Prompter interface {
	Confirm(message string) (bool, error)
}
