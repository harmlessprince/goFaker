package extensions

type ImageUrlParams struct {
	Width     int
	Height    int
	Category  string
	Randomize bool
	Word      string
	Gray      bool
	Format    string
}
type ImageParams struct {
	Width     int
	Height    int
	Category  string
	Randomize bool
	Word      string
	Gray      bool
	Format    string
	Dir       string
	FullPath  bool
}
type ImageExtension interface {
	ImageUrl(params ...ImageUrlParams) string

	Image(params ...ImageParams) string
}

// DefaultImageUrlParams returns default options for ImageUrl
func DefaultImageUrlParams() ImageUrlParams {
	return ImageUrlParams{
		Width:     640,
		Height:    480,
		Category:  "",
		Randomize: true,
		Word:      "",
		Gray:      false,
		Format:    "png",
	}
}

func DefaultImageParams() ImageParams {
	return ImageParams{
		Width:     640,
		Height:    480,
		Category:  "",
		Randomize: true,
		Word:      "",
		Gray:      false,
		Format:    "png",
		Dir:       "",
		FullPath:  true,
	}
}

func NewImageParams(userParams ImageParams) ImageParams {
	defaults := DefaultImageParams()

	if userParams.Width == 0 {
		userParams.Width = defaults.Width
	}
	if userParams.Height == 0 {
		userParams.Height = defaults.Height
	}
	if userParams.Category == "" {
		userParams.Category = defaults.Category
	}
	if userParams.Randomize == false {
		userParams.Randomize = defaults.Randomize
	}
	if userParams.Word == "" {
		userParams.Word = defaults.Word
	}
	if userParams.Gray == false {
		userParams.Gray = defaults.Gray
	}
	if userParams.Format == "" {
		userParams.Format = defaults.Format
	}
	if userParams.Dir == "" {
		userParams.Dir = defaults.Dir
	}
	if userParams.FullPath == false {
		userParams.FullPath = defaults.FullPath
	}

	return userParams
}

func NewImageUrlParams(userParams ImageUrlParams) ImageUrlParams {
	defaults := DefaultImageUrlParams()

	if userParams.Width == 0 {
		userParams.Width = defaults.Width
	}
	if userParams.Height == 0 {
		userParams.Height = defaults.Height
	}
	if userParams.Category == "" {
		userParams.Category = defaults.Category
	}
	if userParams.Randomize == false {
		userParams.Randomize = defaults.Randomize
	}
	if userParams.Word == "" {
		userParams.Word = defaults.Word
	}
	if userParams.Gray == false {
		userParams.Gray = defaults.Gray
	}
	if userParams.Format == "" {
		userParams.Format = defaults.Format
	}

	return userParams
}
