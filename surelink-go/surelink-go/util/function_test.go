package util

import (
	"log"
	"testing"
)

//TODO: write proper test
func TestUrlCheck(t *testing.T) {
	validUrlList := []string{
		"https://google.com",
		"https://google.com?q=qweqwe",
		"https://foo_bar.example.com/",
		"https://a.b-c.de",
		"https://-.~_!$&'()*+,;=:%40:80%2f::::::@example.com",
		"https://उदाहरण.परीक्षा",
		"https://例子.测试",
		"https://1337.net",
		"https://مثال.إختبار",
		"https://foo.bar/?q=Test%20URL-encoded%20stuff",
		"https://j.mp",
		"https://code.google.com/events/#&product=browser",
		"https://☺.damowmow.com/",
		"https://foo.com/(something)?after=parens",
		"https://foo.com/unicode_(✪)_in_parens",
		"https://foo.com/blah_(wikipedia)_blah#cite-1",
		"https://foo.com/blah_(wikipedia)#cite-1",
		"https://⌘.ws/",
		"https://⌘.ws",
		"https://➡.ws/䨹",
		"https://142.42.1.1:8080/",
		"https://142.42.1.1/",
		"https://userid:password@example.com/",
		"https://userid:password@example.com",
		"https://userid@example.com:8080/",
		"https://userid@example.com:8080",
		"https://userid@example.com/",
		"https://userid@example.com",
		"https://userid:password@example.com:8080/",
		"https://userid:password@example.com:8080",
		"https://✪df.ws/123",
		"https://www.example.com/foo/?bar=baz&inga=42&quux",
		"https://www.example.com/wpstyle/?p=364",
		"https://foo.com/blah_blah_(wikipedia)_(again)",
		"https://foo.com/blah_blah_(wikipedia)",
		"https://foo.com/blah_blah/",
		"https://foo.com/blah_blah",
	}

	inValidUrlList := []string{
		"http://223.255.255.254",
		"http://a.b-c.de",
		"ftp://foo.bar/baz",
		"http://j.mp",
		"http://foo.com/unicode_(✪)_in_parens",
		"http://⌘.ws/",
		"http://⌘.ws",
		"http://➡.ws/䨹",
		"http://142.42.1.1:8080/",
		"http://10.1.1.254",
		"http://10.1.1.1",
		"https://10.1.1.1",
		"https://.www.foo.bar./",
		"https://.www.foo.bar/",
		"https://3628126748",
		"https://123.123.123",
		"https://1.1.1.1.1",
		"https://224.1.1.1",
		"https://10.1.1.255",
		"https://10.1.1.0",
		"https://0.0.0.0",
		"https://a.b-.co",
		"https://-a.b.co",
		"https://a.b--c.de/",
		"https://-error-.invalid/",
		"ftps://foo.bar/",
		"https://foo.bar/foo(bar)baz quux",
		":// should fail",
		"https:// shouldfail.com",
		"h://test",
		"rdar://1234",
		"foo.com",
		"https:///a",
		"///",
		"///a",
		"//a",
		"//",
		"https://foo.bar?q=Spaces should be encoded",
		"https://##/",
		"https://##",
		"https://#",
		"https://??/",
		"https://??",
		"https://?",
		"https://../",
		"https://..",
		"https://.",
		"https://",
	}

	log.Println("~~~~~~ VALID URLs")
	for i := 0; i < len(validUrlList); i++ {
		valid, _ := IsValidHttpsUrl(ctx, redisStore, validUrlList[i])
		log.Println(validUrlList[i], valid)
	}

	log.Println("~~~~~~ INVALID URLs")
	for i := 0; i < len(inValidUrlList); i++ {
		valid, _ := IsValidHttpsUrl(ctx, redisStore, inValidUrlList[i])
		log.Println(inValidUrlList[i], valid)
	}
}
