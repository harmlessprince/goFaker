package tests

import (
	"fmt"
	"github.com/harmlessprince/goFaker/extensions"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/stretchr/testify/assert"
	"image"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var baseImageInstance = &providers.BaseImage{}

func TestImageUrl(t *testing.T) {
	t.Run("validate image url is generated without supplying any options", func(t *testing.T) {
		expected := baseImageInstance.ImageUrl()
		assert.True(t, len(expected) > 0)
	})
	t.Run("validate image url is valid", func(t *testing.T) {
		expected := baseImageInstance.ImageUrl()
		_, err := url.ParseRequestURI(expected)
		assert.Nil(t, err)
	})
	t.Run("validate ImageUrl response is reachable", func(t *testing.T) {
		expected := baseImageInstance.ImageUrl()
		assert.True(t, isURLReachable(expected))
	})
	t.Run("validate ImageUrl is uses 640x480 as default size", func(t *testing.T) {
		expected := baseImageInstance.ImageUrl()
		pattern := `^https://via.placeholder.com/640x480.png/.*$`
		assert.Regexp(t, pattern, expected)
	})
	t.Run("validate ImageUrl is uses custom with and height", func(t *testing.T) {
		expected := baseImageInstance.ImageUrl(extensions.ImageUrlParams{
			Width:  800,
			Height: 600,
		})
		pattern := `^https://via.placeholder.com/800x600.png/.*$`
		assert.Regexp(t, pattern, expected)
	})
	t.Run("validate ImageUrl is uses custom category", func(t *testing.T) {
		expected := baseImageInstance.ImageUrl(extensions.ImageUrlParams{
			Width:    800,
			Height:   600,
			Category: "nature",
		})
		pattern := `^https://via.placeholder.com/800x600.png/[a-zA-Z0-9_]{6}\?text=nature\+.*$`

		assert.Regexp(t, pattern, expected)
	})

	t.Run("validate ImageUrl is uses custom text", func(t *testing.T) {
		expected := baseImageInstance.ImageUrl(extensions.ImageUrlParams{
			Width:    800,
			Height:   600,
			Category: "nature",
			Word:     "Faker",
		})
		pattern := `^https://via.placeholder.com/800x600.png/[a-zA-Z0-9_]{6}\?text=nature\+Faker.*$`

		assert.Regexp(t, pattern, expected)
	})
}

func TestImage(t *testing.T) {
	t.Run("validate download with defaults", func(t *testing.T) {
		err := checkURLConnection(providers.BaseImageBaseUrl)
		if err != nil {
			return
		}
		file := baseImageInstance.Image()
		fmt.Println(file)
		assert.FileExists(t, file)
		checkImageProperties(t, file, 640, 480, "png")
	})
	t.Run("validate download with different image format", func(t *testing.T) {
		err := checkURLConnection(providers.BaseImageBaseUrl)
		if err != nil {
			t.Error(err)
			return
		}
		file := baseImageInstance.Image()
		fmt.Println(file)
		assert.FileExists(t, file)
		checkImageProperties(t, "dd", 640, 480, "png")
	})
}

func isURLReachable(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusOK {
		return true
	}

	return false
}

func checkURLConnection(url string) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	if response.StatusCode < 200 || response.StatusCode > 300 {
		return fmt.Errorf("%s is offline, skipping test", url)
	}

	return nil
}

func checkImageProperties(t *testing.T, file string, width int, height int, format string) {
	imgFile, err := os.Open(file)
	if err != nil {
		t.Errorf("Error opening the image file: %v", err)
		return
	}
	defer imgFile.Close()
	imgConfig, _, err := image.DecodeConfig(imgFile)
	if err != nil {
		t.Errorf("Error reading image properties: %v", err)
		return
	}
	imageInfo := pathInfo(file)
	imageFormats := baseImageInstance.GetImageFormatConstants()
	assert.Equal(t, width, imgConfig.Width)
	assert.Equal(t, height, imgConfig.Height)
	assert.Equal(t, width, imgConfig.Width)
	assert.Equal(t, imageFormats[format], imageInfo["extension"])
}
func pathInfo(path string) map[string]string {
	result := make(map[string]string)

	filename := filepath.Base(path)
	ext := filepath.Ext(filename)

	dirname, basename := filepath.Split(path)
	basename = basename[:len(basename)-len(ext)]
	result["dirname"] = dirname
	result["basename"] = basename
	result["extension"] = ext
	result["filename"] = filename

	return result
}
