package domain_parser

import (
	"fmt"
	"log"
	"testing"
)

var rawURL = "https://github.com/feiyizhou"

func Test_GetTopLevelDomain(t *testing.T) {
	topLevelDomain, err := GetTopLevelDomain(rawURL)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(topLevelDomain)
}

func Test_GetDomainByLevel(t *testing.T) {
	domain, err := GetDomainByLevel(rawURL, 0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(domain)
}
