package util

import "time"

const REDIS_REDIRECTION_KEY_PREFIX = "CACHE_REDIRECTION_"
const REDIS_CAPTCHA_KEY_PREFIX = "CACHE_CAPTCHA_"
const REDIS_VALID_HOST_URL_PREFIX = "CACHE_VALID_HOST_URL_"
const FONT_COMIC_PATH = "assets/comic.ttf"
const CAPTCHA_TEXT_LENGTH = 6
const SHORT_URL_UID_LENGTH = 6

var REDIS_CAPTCHA_TTL, _ = time.ParseDuration("5m")
var REDIS_URL_HOST_VALIDITY_TTL, _ = time.ParseDuration("120h")
var REDIS_URL_MAP_TTL, _ = time.ParseDuration("5m")
