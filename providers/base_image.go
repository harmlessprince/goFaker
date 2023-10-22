package providers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/harmlessprince/goFaker/extensions"
	"golang.org/x/exp/slices"
	"image"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

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

func (i *BaseImage) ImageUrl(params ...extensions.ImageUrlParams) string {
	var imageUrlParams extensions.ImageUrlParams
	if len(params) > 0 {
		imageUrlParams = extensions.NewImageUrlParams(params[0])
	} else {
		imageUrlParams = extensions.DefaultImageUrlParams()
	}
	imageFormats := i.GetImageFormats()
	if slices.Contains(imageFormats, imageUrlParams.Format) == false {
		log.Fatal(fmt.Sprintf("invalid image format %s. Allowable formats are: %s", imageUrlParams.Format, imageFormats))
		return ""
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
		imageParts = append(imageParts, baseLorem.Word())
	}
	var backgroundColor string
	if imageUrlParams.Gray == true {
		backgroundColor = "CCCCCC"
	} else {
		color := BaseColor{}
		safeHexColor := color.SafeHexColor()
		backgroundColor = strings.Replace(safeHexColor, "#", "", -1)
	}
	countImageParts := len(imageParts)
	var text string
	if countImageParts > 0 {
		text = "?text=" + url.QueryEscape(strings.Join(imageParts, " "))
	}
	return fmt.Sprintf("%s/%s/%s%s", BaseImageBaseUrl, size, backgroundColor, text)
}

func (i *BaseImage) Image(params ...extensions.ImageParams) string {
	imageParam := extensions.DefaultImageParams()
	if len(params) > 0 {
		imageParam = extensions.NewImageParams(params[0])
	}
	if imageParam.Dir == "" {
		imageParam.Dir = os.TempDir()
	}
	if _, err := helperInstance.IsDirectory(imageParam.Dir); err != nil {
		log.Fatal(err)
	}
	if _, err := helperInstance.IsWritable(imageParam.Dir); err != nil {
		log.Fatal(err)
	}
	hostName, _ := os.Hostname()
	uniqueID := time.Now().UnixNano()
	data := hostName + strconv.FormatInt(uniqueID, 10)
	hash := md5.Sum([]byte(data))
	hashedString := hex.EncodeToString(hash[:])

	fileName := fmt.Sprintf("%s.%s", hashedString, imageParam.Format)
	filePath := imageParam.Dir + fileName
	imageUrl := i.ImageUrl(extensions.ImageUrlParams{
		Width:     imageParam.Width,
		Height:    imageParam.Height,
		Category:  imageParam.Category,
		Randomize: imageParam.Randomize,
		Word:      imageParam.Word,
		Gray:      imageParam.Gray,
		Format:    imageParam.Format,
	})
	imgFile, err := os.Open(imageUrl)
	if err != nil {

	}
	imgConfig, _, err := image.DecodeConfig(imgFile)
	_, err := i.downloadFileFromUrl(imageUrl, filePath)
	if err != nil {
		log.Fatal(err)
	}
	if imageParam.FullPath {
		return filePath
	}
	return fileName
}

func (i *BaseImage) downloadFileFromUrl(url, filePath string) (bool, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return false, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
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
