package providers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

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
type ImageInterface interface {
	ImageUrl(params ...ImageUrlParams) (string, error)

	Image(params ...ImageParams) (string, error)
}

// DefaultImageUrlParams returns default options for ImageUrl
func defaultImageUrlParams() ImageUrlParams {
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

func defaultImageParams() ImageParams {
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

func newImageParams(userParams ImageParams) ImageParams {
	defaults := defaultImageParams()

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

func newImageUrlParams(userParams ImageUrlParams) ImageUrlParams {
	defaults := defaultImageUrlParams()

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

const BaseImageBaseUrl = "https://via.placeholder.com"
const FormatJpg = "jpg"
const FormatPng = "png"
const FormatJpeg = "jpeg"

type BaseImage struct {
	categories []string
}

func (i *BaseImage) GetCategories() []string {
	return []string{
		"abstract", "animals", "business", "cats", "city", "food", "nightlife",
		"fashion", "people", "nature", "sports", "technics", "transport",
	}
}

func (i *BaseImage) GetImageFormats() []string {
	return []string{
		FormatPng, FormatJpg, FormatJpeg,
	}
}
func (i *BaseImage) GetImageFormatConstants() map[string]string {
	return map[string]string{
		FormatPng:  "png",
		FormatJpg:  "jpg",
		FormatJpeg: "jpeg",
	}
}

func (i *BaseImage) ImageUrl(params ...ImageUrlParams) (string, error) {
	var imageUrlParams ImageUrlParams
	if len(params) > 0 {
		imageUrlParams = newImageUrlParams(params[0])
	} else {
		imageUrlParams = defaultImageUrlParams()
	}
	imageFormats := i.GetImageFormats()
	if slices.Contains(imageFormats, imageUrlParams.Format) == false {
		return "", errors.New(fmt.Sprintf("invalid image format %s. Allowable formats are: %s", imageUrlParams.Format, imageFormats))
	}

	size := fmt.Sprintf("%dx%d.%s", imageUrlParams.Width, imageUrlParams.Height, imageUrlParams.Format)

	var imageParts []string
	if imageUrlParams.Category != "" {
		imageParts = append(imageParts, imageUrlParams.Category)
	}

	if imageUrlParams.Word != "" {
		imageParts = append(imageParts, imageUrlParams.Word)
	}

	if imageUrlParams.Randomize == true {
		baseLorem := &BaseLorem{}
		word, err := baseLorem.Word()
		if err != nil {
			return "", err
		}
		imageParts = append(imageParts, word)
	}
	var backgroundColor string
	if imageUrlParams.Gray == true {
		backgroundColor = "CCCCCC"
	} else {
		color := BaseColor{}
		safeHexColor, err := color.SafeHexColor()
		if err != nil {
			return "", err
		}
		backgroundColor = strings.Replace(safeHexColor, "#", "", -1)
	}
	countImageParts := len(imageParts)
	var text string
	if countImageParts > 0 {
		text = "?text=" + url.QueryEscape(strings.Join(imageParts, " "))
	}
	return fmt.Sprintf("%s/%s/%s%s", BaseImageBaseUrl, size, backgroundColor, text), nil
}

func (i *BaseImage) Image(params ...ImageParams) (string, error) {
	imageParam := defaultImageParams()
	if len(params) > 0 {
		imageParam = newImageParams(params[0])
	}
	if imageParam.Dir == "" {
		imageParam.Dir = os.TempDir()
	}
	if _, err := helperInstance.IsDirectory(imageParam.Dir); err != nil {
		return "", err
	}
	if _, err := helperInstance.IsWritable(imageParam.Dir); err != nil {
		return "", err
	}
	hostName, _ := os.Hostname()
	uniqueID := time.Now().UnixNano()
	data := hostName + strconv.FormatInt(uniqueID, 10)
	hash := md5.Sum([]byte(data))
	hashedString := hex.EncodeToString(hash[:])

	fileName := fmt.Sprintf("%s.%s", hashedString, imageParam.Format)
	filePath := imageParam.Dir + fileName
	imageUrl, err := i.ImageUrl(ImageUrlParams{
		Width:     imageParam.Width,
		Height:    imageParam.Height,
		Category:  imageParam.Category,
		Randomize: imageParam.Randomize,
		Word:      imageParam.Word,
		Gray:      imageParam.Gray,
		Format:    imageParam.Format,
	})
	if err != nil {
		return "", err
	}
	_, err = i.downloadFileFromUrl(imageUrl, filePath)
	if err != nil {
		return "", err
	}
	if imageParam.FullPath {
		return filePath, nil
	}
	return fileName, nil
}

func (i *BaseImage) downloadFileFromUrl(url, filePath string) (bool, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return false, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}
	// Check if the HTTP status code is 200 (OK)
	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("HTTP request failed with status code %d", response.StatusCode)
	}
	// Copy the response body to the local file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *BaseImage) fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// Handle other errors, if necessary
		return false
	}
	return true
}
