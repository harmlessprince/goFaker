package providers

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type UserAgentInterface interface {
	MacProcessor() (string, error)
	LinuxProcessor() (string, error)
	UserAgent() (string, error)
	Chrome() (string, error)
	MicrosoftEdge() (string, error)
	FireFox() (string, error)
	Safari() (string, error)
	Opera() (string, error)
	InternetExplorer() (string, error)
	WindowsPlatformToken() (string, error)
	MacPlatformToken() (string, error)
	IosMobileToken() (string, error)
	LinuxPlatformToken() (string, error)
}

type BaseUserAgent struct {
	BaseProvider
	userAgents            []string
	windowsPlatformTokens []string
	linuxProcessors       []string
	macProcessors         []string
	lang                  []string
}

func (ua *BaseUserAgent) GetUserAgents() []string {
	return ua.userAgents
}

func (ua *BaseUserAgent) SetUserAgents(userAgents []string) {
	ua.userAgents = userAgents
}

func (ua *BaseUserAgent) GetWindowsPlatformTokens() []string {
	return ua.windowsPlatformTokens
}

func (ua *BaseUserAgent) SetWindowsPlatformTokens(tokens []string) {
	ua.windowsPlatformTokens = tokens
}

func (ua *BaseUserAgent) GetLinuxProcessors() []string {
	return ua.linuxProcessors
}

func (ua *BaseUserAgent) SetLinuxProcessors(processors []string) {
	ua.linuxProcessors = processors
}

func (ua *BaseUserAgent) GetMacProcessors() []string {
	return ua.macProcessors
}

func (ua *BaseUserAgent) SetMacProcessors(processors []string) {
	ua.macProcessors = processors
}
func defaultUserAgents() []string {
	return []string{
		"firefox", "chrome", "internetExplorer", "opera", "safari", "msedge",
	}
}

func defaultWindowsPlatformTokens() []string {
	return []string{
		"Windows NT 6.2", "Windows NT 6.1", "Windows NT 6.0", "Windows NT 5.2", "Windows NT 5.1",
		"Windows NT 5.01", "Windows NT 5.0", "Windows NT 4.0", "Windows 98; Win 9x 4.90", "Windows 98",
		"Windows 95", "Windows CE",
	}

}

func defaultLinuxProcessors() []string {
	return []string{
		"i686", "x86_64",
	}
}

func defaultMacProcessors() []string {
	return []string{
		"Intel", "PPC", "U; Intel", "U; PPC",
	}
}

func defaultLang() []string {
	return []string{
		"en-US", "sl-SI", "nl-NL",
	}
}

func NewBaseUserAgent() *BaseUserAgent {
	ug := &BaseUserAgent{
		userAgents:            defaultUserAgents(),
		windowsPlatformTokens: defaultWindowsPlatformTokens(),
		linuxProcessors:       defaultLinuxProcessors(),
		macProcessors:         defaultMacProcessors(),
		lang:                  defaultLang(),
	}
	return ug
}

func (ua *BaseUserAgent) MacProcessor() (string, error) {
	processor := helperInstance.RandomElementString(ua.macProcessors)
	return processor, nil
}

func (ua *BaseUserAgent) LinuxProcessor() (string, error) {
	processor := helperInstance.RandomElementString(ua.linuxProcessors)
	return processor, nil
}

func (ua *BaseUserAgent) UserAgent() (string, error) {
	agent := helperInstance.RandomElementString(ua.userAgents)
	userAgent, err := ua.BaseProvider.Parse(strings.ToTitle(agent), ua)
	if err != nil {
		return "", err
	}
	return userAgent, nil
}

func (ua *BaseUserAgent) Chrome() (string, error) {
	firstNumberBtw, err := ua.NumberBetween(531, 536)
	if err != nil {
		return "", err
	}
	secondNumberBtw, err := ua.NumberBetween(0, 2)
	if err != nil {
		return "", err
	}
	saf := strconv.Itoa(firstNumberBtw) + strconv.Itoa(secondNumberBtw)
	windowsPlatformToken, err := ua.WindowsPlatformToken()
	if err != nil {
		return "", err
	}
	macPlatformToken, err := ua.MacPlatformToken()
	if err != nil {
		return "", err
	}
	linuxPlatformToken, err := ua.LinuxPlatformToken()
	if err != nil {
		return "", err
	}
	platFormTokens := []string{
		linuxPlatformToken,
		windowsPlatformToken,
		macPlatformToken,
	}

	var platforms []string
	for i := 0; i < len(platFormTokens); i++ {
		numberBtw36to40, err := ua.NumberBetween(36, 40)
		if err != nil {
			log.Fatal(err)
		}
		numberBtw800to899, err := ua.NumberBetween(800, 899)
		if err != nil {
			log.Fatal(err)
		}
		var platform strings.Builder
		platform.WriteString("(")
		platform.WriteString(platFormTokens[i])
		platform.WriteString(") AppleWebKit/" + saf + "(KHTML, like Gecko) Chrome/")
		platform.WriteString(strconv.Itoa(numberBtw36to40) + ".0.")
		platform.WriteString(strconv.Itoa(numberBtw800to899) + ".0 Mobile Safari/" + saf)
		platforms = append(platforms, platform.String())
	}
	platform, err := ua.BaseProvider.RandomElementFromStringSlice(platforms)
	if err != nil {
		log.Fatal(err)
	}
	return "Mozilla/5.0 " + platform, nil
}

func (ua *BaseUserAgent) MicrosoftEdge() (string, error) {
	firstNumberBtw, _ := ua.BaseProvider.NumberBetween(531, 537)
	secondNumberBtw, _ := ua.BaseProvider.NumberBetween(0, 2)
	saf := strconv.Itoa(firstNumberBtw) + strconv.Itoa(secondNumberBtw)
	chrv, _ := ua.BaseProvider.NumberBetween(79, 99)
	chrvToString := strconv.Itoa(chrv) + ".0"
	platformTokenNumberBtw4000To4844, _ := ua.BaseProvider.NumberBetween(4000, 4844)
	platformTokenNumberBtw10To99, _ := ua.BaseProvider.NumberBetween(10, 99)
	platformTokenNumberBtw1000To1146, _ := ua.BaseProvider.NumberBetween(1000, 1146)
	platformTokenNumberBtw0To99, _ := ua.BaseProvider.NumberBetween(0, 99)
	windowsPlatformToken, err := ua.WindowsPlatformToken()
	if err != nil {
		return "", err
	}
	macPlatformToken, err := ua.MacPlatformToken()
	if err != nil {
		return "", err
	}
	linuxPlatformToken, err := ua.LinuxPlatformToken()
	if err != nil {
		return "", err
	}
	iosPlatformToken, err := ua.IosMobileToken()
	if err != nil {
		return "", err
	}
	platforms := []string{
		"(" + windowsPlatformToken + ") AppleWebKit/" + saf + "(KHTML, like Gecko) Chrome/" + chrvToString + "." + strconv.Itoa(platformTokenNumberBtw4000To4844) + "." + strconv.Itoa(platformTokenNumberBtw10To99) + " Safari/" + saf + " Edg/" + chrvToString + strconv.Itoa(platformTokenNumberBtw1000To1146) + "." + strconv.Itoa(platformTokenNumberBtw0To99),
		"(" + macPlatformToken + ") AppleWebKit/" + saf + "(KHTML, like Gecko) Chrome/" + chrvToString + "." + strconv.Itoa(platformTokenNumberBtw4000To4844) + "." + strconv.Itoa(platformTokenNumberBtw10To99) + " Safari/" + saf + " Edg/" + chrvToString + strconv.Itoa(platformTokenNumberBtw1000To1146) + "." + strconv.Itoa(platformTokenNumberBtw0To99),
		"(" + linuxPlatformToken + ") AppleWebKit/" + saf + "(KHTML, like Gecko) Chrome/" + chrvToString + "." + strconv.Itoa(platformTokenNumberBtw4000To4844) + "." + strconv.Itoa(platformTokenNumberBtw10To99) + " Safari/" + saf + " Edg/A" + chrvToString + strconv.Itoa(platformTokenNumberBtw1000To1146) + "." + strconv.Itoa(platformTokenNumberBtw0To99),
		"(" + iosPlatformToken + ") AppleWebKit/" + saf + "(KHTML, like Gecko) Version/15.0 EdgiOS/" + chrvToString + "." + strconv.Itoa(platformTokenNumberBtw1000To1146) + "." + strconv.Itoa(platformTokenNumberBtw10To99) + " Mobile/15E148 Safari/" + saf,
	}
	platform, err := ua.BaseProvider.RandomElementFromStringSlice(platforms)
	if err != nil {
		log.Fatal(err)
	}
	return "Mozilla/5.0 " + platform, nil
}

func (ua *BaseUserAgent) FireFox() (string, error) {
	versionNumberBtw35To37, _ := ua.BaseProvider.NumberBetween(35, 37)
	startTime, err := time.Parse(time.DateOnly, "2010-01-01")
	if err != nil {
		log.Fatal(err)
	}
	endTime := time.Now()
	versionNumberBtwDateTimeNumber, _ := ua.BaseProvider.NumberBetween(int(startTime.Unix()), int(endTime.Unix()))
	versionYear := time.Unix(int64(versionNumberBtwDateTimeNumber), 0).Year()
	versionMonth := time.Unix(int64(versionNumberBtwDateTimeNumber), 0).Month()
	versionDay := time.Unix(int64(versionNumberBtwDateTimeNumber), 0).Day()
	version := "Gecko/" + strconv.Itoa(versionYear) + strconv.Itoa(int(versionMonth)) + strconv.Itoa(versionDay) + " Firefox/" + strconv.Itoa(versionNumberBtw35To37) + ".0"
	numberBtw2To6, _ := ua.BaseProvider.NumberBetween(2, 6)
	numberBtw0To2, _ := ua.BaseProvider.NumberBetween(0, 2)
	numberBtw5To7, _ := ua.BaseProvider.NumberBetween(5, 7)
	radomLang, _ := ua.BaseProvider.RandomElementFromStringSlice(ua.lang)
	windowsPlatformToken, err := ua.WindowsPlatformToken()
	if err != nil {
		return "", err
	}
	macPlatformToken, err := ua.MacPlatformToken()
	if err != nil {
		return "", err
	}
	linuxPlatformToken, err := ua.LinuxPlatformToken()
	if err != nil {
		return "", err
	}
	platforms := []string{
		"(" + windowsPlatformToken + "; " + radomLang + "; rv:1.9." + strconv.Itoa(numberBtw0To2) + ".20)" + version,
		"(" + linuxPlatformToken + "; rv:" + strconv.Itoa(numberBtw5To7) + ".0)" + version,
		"(" + macPlatformToken + " rv:" + strconv.Itoa(numberBtw2To6) + ".0)" + version,
	}
	platform, err := ua.BaseProvider.RandomElementFromStringSlice(platforms)
	if err != nil {
		log.Fatal(err)
	}
	return "Mozilla/5.0 " + platform, nil
}

func (ua *BaseUserAgent) Safari() (string, error) {
	numberBtw531To535, _ := ua.BaseProvider.NumberBetween(531, 535)
	numberBtw1To50, _ := ua.BaseProvider.NumberBetween(1, 50)
	numberBtw1To7, _ := ua.BaseProvider.NumberBetween(1, 7)
	numberBtw4To5, _ := ua.BaseProvider.NumberBetween(4, 5)
	numberBtw1T5, _ := ua.BaseProvider.NumberBetween(1, 5)
	numberBtw2T6, _ := ua.BaseProvider.NumberBetween(2, 6)
	numberBtw7T8, _ := ua.BaseProvider.NumberBetween(7, 8)
	numberBtwOT2, _ := ua.BaseProvider.NumberBetween(0, 2)
	numberBtw1T2, _ := ua.BaseProvider.NumberBetween(1, 2)
	numberBtw3To4, _ := ua.BaseProvider.NumberBetween(3, 4)
	numberBtw111To119, _ := ua.BaseProvider.NumberBetween(111, 119)
	radomLang, _ := ua.BaseProvider.RandomElementFromStringSlice(ua.lang)
	saf := strconv.Itoa(numberBtw531To535) + "." + strconv.Itoa(numberBtw1To50) + "." + strconv.Itoa(numberBtw1To7)
	bm := &BaseMiscellaneous{}
	version := strconv.Itoa(numberBtw4To5) + ".0." + strconv.Itoa(numberBtw1T5)
	boolean, _ := bm.Boolean()
	if boolean {
		numberBtw0To1, _ := ua.BaseProvider.NumberBetween(0, 1)
		version = strconv.Itoa(numberBtw4To5) + ".0." + strconv.Itoa(numberBtw0To1)
	}
	mobileDevices := []string{
		"iPhone; CPU iPhone OS",
		"iPad; CPU OS",
	}
	mobileDevice := helperInstance.RandomElementString(mobileDevices)
	windowsPlatformToken, err := ua.WindowsPlatformToken()
	if err != nil {
		return "", err
	}
	macPlatformToken, err := ua.MacPlatformToken()
	if err != nil {
		return "", err
	}
	platforms := []string{
		"(Windows; U; " + windowsPlatformToken + ") AppleWebKit/$saf (KHTML, like Gecko) Version/" + version + " Safari/" + saf,
		"(" + macPlatformToken + " rv:" + strconv.Itoa(numberBtw2T6) + ".0; " + radomLang + ") AppleWebKit/" + saf + "(KHTML, like Gecko) Version/" + version + " Safari/" + saf,
		"(" + mobileDevice + " " + strconv.Itoa(numberBtw7T8) + "_" + strconv.Itoa(numberBtwOT2) + "_" + strconv.Itoa(numberBtw1T2) + " like Mac OS X; " + radomLang + ") AppleWebKit/$saf (KHTML, like Gecko) Version/" + strconv.Itoa(numberBtw3To4) + ".0.5 Mobile/8B" + strconv.Itoa(numberBtw111To119) + " Safari/6$saf",
	}
	platform, err := ua.BaseProvider.RandomElementFromStringSlice(platforms)
	if err != nil {
		return "", err
	}
	return "Mozilla/5.0 " + platform, nil

}

func (ua *BaseUserAgent) Opera() (string, error) {
	numberBtw8To12, _ := ua.BaseProvider.NumberBetween(8, 12)
	numberBtw160To355, _ := ua.BaseProvider.NumberBetween(160, 355)
	numberBtw10To12, _ := ua.BaseProvider.NumberBetween(10, 12)
	numberBtw8To9, _ := ua.BaseProvider.NumberBetween(8, 9)
	numberBtw10To99, _ := ua.BaseProvider.NumberBetween(10, 99)
	radomLang, _ := ua.BaseProvider.RandomElementFromStringSlice(ua.lang)
	linuxPlatformToken, err := ua.LinuxPlatformToken()
	if err != nil {
		return "", err
	}
	windowsPlatformToken, err := ua.WindowsPlatformToken()
	if err != nil {
		return "", err
	}
	platforms := []string{
		"(" + linuxPlatformToken + "; " + radomLang + ") Presto/2." + strconv.Itoa(numberBtw8To12) + "." + strconv.Itoa(numberBtw160To355) + " Version/" + strconv.Itoa(numberBtw10To12) + ".00",
		"(" + windowsPlatformToken + "; " + radomLang + ") Presto/2." + strconv.Itoa(numberBtw8To12) + " Version/" + strconv.Itoa(numberBtw10To12) + ".00",
	}
	platform, err := ua.BaseProvider.RandomElementFromStringSlice(platforms)
	if err != nil {
		log.Fatal(err)
	}
	return "Opera/" + strconv.Itoa(numberBtw8To9) + "." + strconv.Itoa(numberBtw10To99) + "." + platform, nil
}

func (ua *BaseUserAgent) InternetExplorer() (string, error) {
	numberBtw0To1, _ := ua.BaseProvider.NumberBetween(0, 1)
	numberBtw5To11, _ := ua.BaseProvider.NumberBetween(5, 11)
	numberBtw3To5, _ := ua.BaseProvider.NumberBetween(3, 5)
	windowPlatformToken, err := ua.WindowsPlatformToken()
	if err != nil {
		return "", err
	}
	return "Mozilla/5.0 (compatible; MSIE " + strconv.Itoa(numberBtw5To11) + ".0; " + windowPlatformToken + "; Trident/" + strconv.Itoa(numberBtw3To5) + "." + strconv.Itoa(numberBtw0To1) + ")", nil
}

func (ua *BaseUserAgent) WindowsPlatformToken() (string, error) {
	token, err := ua.BaseProvider.RandomElementFromStringSlice(ua.windowsPlatformTokens)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ua *BaseUserAgent) IosMobileToken() (string, error) {
	firstNumberBtw, err := ua.NumberBetween(13, 16)
	if err != nil {
		log.Fatal(err)
	}
	secondNumberBtw, err := ua.NumberBetween(0, 2)
	if err != nil {
		log.Fatal(err)
	}
	iosVersion := strconv.Itoa(firstNumberBtw) + "_" + strconv.Itoa(secondNumberBtw)
	return "iPhone; CPU iPhone OS " + iosVersion + " like Mac OS X", nil
}

func (ua *BaseUserAgent) LinuxPlatformToken() (string, error) {
	processor, err := ua.BaseProvider.RandomElementFromStringSlice(ua.linuxProcessors)
	if err != nil {
		log.Fatal(err)
	}
	return "X11; Linux " + processor, nil
}

func (ua *BaseUserAgent) MacPlatformToken() (string, error) {
	processor, err := ua.BaseProvider.RandomElementFromStringSlice(ua.macProcessors)
	if err != nil {
		log.Fatal(err)
	}
	numberBtw5to8, err := ua.BaseProvider.NumberBetween(5, 8)
	if err != nil {
		log.Fatal(err)
	}
	numberBtw0to9, err := ua.BaseProvider.NumberBetween(5, 8)
	if err != nil {
		log.Fatal(err)
	}
	var token strings.Builder
	token.WriteString("Macintosh; ")
	token.WriteString(processor)
	token.WriteString(" Mac OS X 10_")
	token.WriteString(strconv.Itoa(numberBtw5to8) + "_")
	token.WriteString(strconv.Itoa(numberBtw0to9))
	return token.String(), nil
}
