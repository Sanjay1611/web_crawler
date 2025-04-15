package result

type DownloadResult struct {
	Source  string
	Content []byte
	Err     error
}
