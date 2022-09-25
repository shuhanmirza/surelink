package util

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/url"
	gedis "surelink-go/redisStore"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IsValidHttpsUrl TODO improve accuracy and latency
func IsValidHttpsUrl(ctx *gin.Context, redisStore *gedis.RedisStore, urlString string) (bool, error) {
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

	redisKey := REDIS_URL_HOST_VALIDITY_PREFIX + urlObj.Host
	redisValue, err := redisStore.Client.Get(ctx, redisKey).Bool()

	//redis entry does not exist
	if err != nil {
		// Check it's a valid domain name
		_, err = net.LookupHost(urlObj.Host)
		if err != nil {
			log.Println("error while url host lookup")
			log.Println(err)

			go redisStore.Client.Set(ctx, redisKey, false, REDIS_URL_HOST_VALIDITY_TTL)

			return false, &UrlHostInvalidError{}
		}

		go redisStore.Client.Set(ctx, redisKey, true, REDIS_URL_HOST_VALIDITY_TTL)

		return true, nil
	}

	if false == redisValue {
		log.Println("error while url host lookup")
		log.Println(err)
		return false, &UrlHostInvalidError{}
	}

	return true, nil

}
