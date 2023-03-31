package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net"
	"net/url"
	"strings"
	"surelink-go/infrastructure"
	"surelink-go/util"
)

const alphabet = "abcdefghijklmnopqrstwxyzABCDEFGHIJKLMNOPQRSTUVXYZ0123456789="

type UtilityService struct {
	cache  *infrastructure.Cache
	random *rand.Rand
}

func NewUtilityService(cache *infrastructure.Cache, random *rand.Rand) UtilityService {
	return UtilityService{
		cache:  cache,
		random: random,
	}
}

func (s UtilityService) IsValidHttpsUrl(ctx *gin.Context, urlString string) (bool, error) {
	urlObj, err := url.ParseRequestURI(urlString)
	if err != nil {
		log.Println("error while url parsing")
		log.Println(err)
		return false, &util.UrlParsingError{}
	}

	// Check it's an acceptable scheme
	switch urlObj.Scheme {
	case "https":
	case "ftp":
	default:
		log.Println(urlObj.Scheme + " is not accepted")
		return false, &util.UrlProtocolNotAcceptedError{}
	}

	redisKey := util.REDIS_VALID_HOST_URL_PREFIX + urlObj.Host
	redisValue, err := s.cache.Client.Get(ctx, redisKey).Bool() //returns err when redis entry does not exist

	if err != nil {
		// Check it's a valid domain name
		_, err = net.LookupHost(urlObj.Host)
		if err != nil {
			log.Println("error while url host lookup")
			log.Println(err)

			go s.cache.Client.Set(ctx, redisKey, false, util.REDIS_URL_HOST_VALIDITY_TTL)

			return false, &util.UrlHostInvalidError{}
		}

		go s.cache.Client.Set(ctx, redisKey, true, util.REDIS_URL_HOST_VALIDITY_TTL)

		return true, nil
	}

	if false == redisValue {
		log.Println("error while url host lookup")
		log.Println(err)
		return false, &util.UrlHostInvalidError{}
	}

	return true, nil
}

func (s UtilityService) RandomInt(min, max int64) int64 {
	return min + s.random.Int63n(max-min+1)
}

func (s UtilityService) RandomBool() bool {
	if s.random.Intn(2) == 0 {
		return false
	}
	return true
}

func (s UtilityService) RandomStringAlphabet(n int) string {
	var stringBuilder strings.Builder

	lenAlpha := len(alphabet)
	for i := 0; i < n; i++ {
		char := alphabet[s.random.Intn(lenAlpha)]
		stringBuilder.WriteByte(char)
	}

	return stringBuilder.String()
}
