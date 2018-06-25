package internet

import (
	"fmt"
	"github.com/DarthPestilane/gomock/helper"
	"github.com/DarthPestilane/gomock/random/basic"
	"github.com/DarthPestilane/gomock/random/name"
	"strings"
)

var (
	EmailDomains = []string{"gmail.com", "yahoo.com", "hotmail.com"}
	TLDs         = []string{"com", "info", "net", "org", "io", "tv", "biz", "cc", "top"}
)

func Email(domain ...string) string {
	fullname := strings.Split(name.Name(), " ")
	idx := helper.Random(2)
	user := fullname[idx]

	var emailDomain string
	if len(domain) > 0 && domain[0] != "" {
		emailDomain = domain[0]
	} else {
		helper.RandomSlice(EmailDomains, &emailDomain)
	}
	return fmt.Sprintf("%s@%s", user, emailDomain)
}

// DomainWord
// example "clark"
func DomainWord() string {
	var word string
	helper.RandomSlice(name.LastNames, &word)
	return strings.ToLower(word)
}

// TLD
// example "com"
func TLD() string {
	var tld string
	helper.RandomSlice(TLDs, &tld)
	return tld
}

// DomainName
// example "clark.com"
func DomainName() string {
	return fmt.Sprintf("%s.%s", DomainWord(), TLD())
}

// Url
func Url() string {
	hasWWW := basic.Bool()
	shouldHTTPS := basic.Bool()
	hasSuffix := basic.Bool()
	hasPath := basic.Bool()
	var url string
	if shouldHTTPS {
		url = "https://"
	} else {
		url = "http://"
	}
	if hasWWW {
		url += "www."
	}
	url += DomainName() + "/"
	if hasPath {
		url += basic.String(basic.OptionString{MaxLen: 5, MinLen: 2, FromChars: strings.Join(basic.Lower, "")})
		if hasSuffix {
			url += ".html"
		}
	}
	return url
}
