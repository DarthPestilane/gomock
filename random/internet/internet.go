package internet

import (
	"fmt"
	"github.com/DarthPestilane/gofake/helper"
	"github.com/DarthPestilane/gofake/random/basic"
	"github.com/DarthPestilane/gofake/random/name"
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

// MacAddr
// example "32:F1:39:2F:D6:18"
func MacAddr() string {
	from := append(basic.Upper, basic.NumsStr...)
	meta := basic.String(basic.OptionString{MaxLen: 12, MinLen: 12, FromChars: strings.Join(from, "")})
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s", meta[0:2], meta[2:4], meta[4:6], meta[6:8], meta[8:10], meta[10:12])
}
