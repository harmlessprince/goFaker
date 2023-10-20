package extensions

type FileExtension interface {
	MimeType() string
	Extension() string
	FilePath() string
}
