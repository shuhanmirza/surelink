package util

import "time"

const RedisRedirectionKeyPrefix = "CACHE_REDIRECTION_"
const RedisCaptchaKeyPrefix = "CACHE_CAPTCHA_"
const RedisCaptchaQueueKey = "CACHE_QUEUE_CAPTCHA"
const RedisValidHostUrlPrefix = "CACHE_VALID_HOST_URL_"
const FontComicPath = "assets/comic.ttf"
const CaptchaTextLength = 6
const ShortUrlUidLength = 6
const CaptchaQueueMaxSize = 60

var RedisCaptchaTtl, _ = time.ParseDuration("5m")
var RedisUrlHostValidityTtl, _ = time.ParseDuration("120h")
var RedisUrlMapTtl, _ = time.ParseDuration("5m")

const CronSpecEvery10Min = "@every 1m"
