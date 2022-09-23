package util

import (
	"log"
	"net"
	"net/url"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//TODO improve accuracy
func IsValidHttpsUrl(urlString string) (bool, error) {
	urlObj, err := url.ParseRequestURI(urlString)
	if err != nil {
		log.Println("error while url parsing")
		log.Println(err)
		return false, &UrlParsingError{}
	}

	// Check it's an acceptable scheme
	switch urlObj.Scheme {
	case "https":
	case "ftp":
	default:
		log.Println(urlObj.Scheme + " is not accepted")
		return false, &UrlProtocolNotAcceptedError{}
	}

	// Check it's a valid domain name
	_, err = net.LookupHost(urlObj.Host)
	if err != nil {
		log.Println("error while url host lookup")
		log.Println(err)
		return false, &UrlHostInvalidError{}
	}

	return true, nil
}
