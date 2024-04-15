# domain-parser

[![Go Report Card](https://pkg.go.dev/badge/github.com/feiyizhou/domain-parser)](https://pkg.go.dev/github.com/feiyizhou/domain-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/feiyizhou/domain-parser)](https://goreportcard.com/report/github.com/feiyizhou/domain-parser)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

Domain Parser can parse any URL. You can obtain the root domain directly or obtain the domain based on the domain level.

> required go 1.21+

## Install

```shell
go install github.com/feiyizhou/domain-parser@latest 
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/feiyizhou/domain-parser"
)

var rawURL = "https://github.com/feiyizhou"

func main() {
	topLevelDomain, err := domain_parser.GetTopLevelDomain(rawURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(topLevelDomain)

	domain, err := domain_parser.GetDomainByLevel(rawURL, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(domain)
	
	domainArr, err := domain_parser.GetDomainArr(rawURL)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(domainArr)
}

```

## License

Domain Parser is available as open source under the terms of the [MIT License](https://github.com/feiyizhou/domain-parser/blob/main/LICENSE)