package providers

import (
	"fmt"
	"github.com/harmlessprince/goFaker/extensions"
	"golang.org/x/exp/slices"
	"log"
	"net/url"
	"strings"
)

const BASE_URL = "https://via.placeholder.com"
const FORMAT_JPG = "jpg"
const FORMAT_PNG = "png"
const FORMAT_JPEG = "jpeg"

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
		FORMAT_PNG, FORMAT_JPG, FORMAT_JPEG,
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
		backgroundColor = strings.Replace(safeHexColor, "", "#", -1)
	}
	countImageParts := len(imageParts)
	var text string
	if countImageParts > 0 {
		text = "?text=" + url.QueryEscape(strings.Join(imageParts, " "))
	}
	return fmt.Sprintf("%s/%s/%s%s", BASE_URL, size, backgroundColor, text)
}

func (i *BaseImage) Image(params ...extensions.ImageParams) string {
	//TODO implement me
	panic("implement me")
}
