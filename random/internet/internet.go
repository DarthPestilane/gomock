package internet

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/DarthPestilane/gofake/helper"
	"github.com/DarthPestilane/gofake/random/basic"
	"github.com/DarthPestilane/gofake/random/name"
	"net"
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
// example "(http|https)://[www.]clark.(com|top|cc...)/[path][.html]"
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
	return split(meta, 2, ":")
}

// IPv4
// example "12.23.45.76"
func IPv4() string {
	long := basic.Int(basic.OptionNumber{Min: int(ip2long("1.0.0.1")), Max: int(ip2long("255.255.255.254"))})
	return long2ip(uint32(long))
}

// IPv6
// example "05d2:1d4c:0802:ec87:12c7:ca94:e78e:ca2a"
func IPv6() string {
	meta := fmt.Sprintf("%x", md5.Sum([]byte(basic.String())))
	return split(meta, 4, ":")
}

// LocalIPv4
// example "10.0.1.12" or "192.168.1.10"
func LocalIPv4() string {
	var long int
	if basic.Bool() {
		long = basic.Int(basic.OptionNumber{Min: int(ip2long("10.0.0.1")), Max: int(ip2long("10.255.255.254"))})
	} else {
		long = basic.Int(basic.OptionNumber{Min: int(ip2long("192.168.0.1")), Max: int(ip2long("192.168.255.254"))})
	}
	return long2ip(uint32(long))
}

func long2ip(long uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, long)
	ip := net.IP(ipByte)
	return ip.String()
}

func ip2long(ipAddr string) uint32 {
	ip := net.ParseIP(ipAddr)
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func split(from string, step int, concat string) string {
	length := len(from)
	arr := []string{}
	for i := 0; i < length; i += step {
		arr = append(arr, from[i:i+step])
	}
	return strings.Join(arr, concat)
}
