package domain_parser

import (
	"errors"
	"net/url"
	"strings"

	"github.com/feiyizhou/domain-parser/top_level_domain"
)

// GetTopLevelDomain get the top level domain
func GetTopLevelDomain(rawURL string) (string, error) {
	domainArr, err := getDomainArr(rawURL)
	if err != nil {
		return "", err
	}
	return domainArr[len(domainArr)-1], err
}

// GetDomainByLevel get the domain by level of domain
//
// The level of domain start at 0, 0 means the top level domain,
// 1 means first level domain, 2 means secondary domain and so on
func GetDomainByLevel(rawURL string, level int) (string, error) {
	domainArr, err := getDomainArr(rawURL)
	if err != nil {
		return "", err
	}
	if level > len(domainArr)-1 || level < 0 {
		return "", errors.New("The level of domain is invalid ")
	}
	return domainArr[len(domainArr)-1-level], err
}

// getHost get the host of url
func getHost(rawURL string) (string, error) {
	fullURL, err := url.Parse(rawURL)
	if err == nil && fullURL != nil {
		return fullURL.Host, err
	}
	return "", err
}

// splitHost split host to string array without empty element
func splitHost(rawURL string) ([]string, error) {
	rawHost, err := getHost(rawURL)
	if err != nil {
		return nil, err
	}
	if len(rawHost) == 0 {
		return nil, errors.New("Parser url failed ")
	}
	rawHost = strings.TrimPrefix(rawHost, "www.")
	hostArr := strings.Split(rawHost, ".")
	for index, host := range hostArr {
		if len(host) == 0 {
			hostArr = append(hostArr[:index], hostArr[index+1:]...)
		}
	}
	return hostArr, err
}

// getDomainArr get the domain string arr of rawURL
func getDomainArr(rawURL string) ([]string, error) {
	hostArr, err := splitHost(rawURL)
	if err != nil {
		return nil, err
	}
	var (
		rootDomainArr []string
		domainMap     = top_level_domain.TopLevelDomainTree
	)
	for i := len(hostArr) - 1; i >= 0; i-- {
		if domain, ok := domainMap[hostArr[i]]; ok {
			rootDomainArr = append([]string{hostArr[i]}, rootDomainArr...)
			hostArr = hostArr[:i]
			if domain == nil {
				break
			} else {
				domainMap = *domain
			}
		}
	}
	if len(rootDomainArr) == 0 {
		return nil, errors.New("Unsupported top level domain. Check your rawURL or add new top level domain to top_level_domain.json ")
	}
	hostArr = append(hostArr, strings.Join(rootDomainArr, "."))
	return hostArr, err
}
