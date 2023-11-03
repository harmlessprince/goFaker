package tests

import (
	"fmt"
	"github.com/harmlessprince/goFaker/providers"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var baseUserAgentInstanceTest = providers.NewBaseUserAgent()

//func TestName(t *testing.T) {
//	t.Run("description", func(t *testing.T) {})
//	t.Run("description", func(t *testing.T) {})
//}

func TestMacProcessor(t *testing.T) {
	t.Run("test MacProcessor a valid mac processor", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.MacProcessor()
		assert.Contains(t, baseUserAgentInstanceTest.GetMacProcessors(), expected)
	})
}

func TestLinuxProcessor(t *testing.T) {
	t.Run("test LinuxProcessor a valid linux processor", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.LinuxProcessor()
		assert.Contains(t, baseUserAgentInstanceTest.GetLinuxProcessors(), expected)
	})
}

func TestUserAgent(t *testing.T) {
	t.Run("test UserAgent a valid user agent", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.UserAgent()
		assert.NotEmpty(t, expected)
	})
}

func TestChrome(t *testing.T) {
	t.Run("test Chrome returns a valid chrome browser", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.Chrome()
		assert.NotEmpty(t, expected)
	})
	t.Run("test Chrome user agent returns string that contains '(KHTML, like Gecko) Chrome/'", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.Chrome()
		assert.True(t, strings.Contains(expected, "(KHTML, like Gecko) Chrome/"))
	})
}
func TestMicrosoftEdge(t *testing.T) {
	t.Run("test MicrosoftEdge returns a valid MicrosoftEdge browser", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.MicrosoftEdge()
		assert.NotEmpty(t, expected)
	})
	t.Run("test MicrosoftEdge user agent returns string that contains 'Edg'", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.MicrosoftEdge()
		assert.True(t, strings.Contains(expected, "Edg"))
	})
}

func TestFireFox(t *testing.T) {
	t.Run("test FireFox returns a valid FireFox user agent", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.FireFox()
		assert.NotEmpty(t, expected)
	})
	t.Run("test FireFox returns string that contains 'Firefox/'", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.FireFox()
		fmt.Println(expected)
		assert.True(t, strings.Contains(expected, "Firefox/"))
	})
}

func TestSafari(t *testing.T) {
	t.Run("test Safari returns a valid Safari browser", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.Safari()
		assert.NotEmpty(t, expected)
	})
	t.Run("test Safari user agent returns string that contains 'Safari/'", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.Safari()
		assert.True(t, strings.Contains(expected, "Safari/"))
	})
}

func TestOpera(t *testing.T) {
	t.Run("test Opera returns a valid Opera browser", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.Opera()
		assert.NotEmpty(t, expected)
	})
	t.Run("test Opera user agent starts with 'Opera/'", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.Opera()
		assert.True(t, strings.HasPrefix(expected, "Opera/"))
	})
}

func TestInternetExplorer(t *testing.T) {
	t.Run("test InternetExplorer returns a valid InternetExplorer browser", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.InternetExplorer()
		assert.NotEmpty(t, expected)
	})

	t.Run("test InternetExplorer user agent starts with Mozilla/5.0 (compatible; MSIE", func(t *testing.T) {
		expected, _ := baseUserAgentInstanceTest.InternetExplorer()
		assert.True(t, strings.HasPrefix(expected, "Mozilla/5.0 (compatible; MSIE"))
	})
}
